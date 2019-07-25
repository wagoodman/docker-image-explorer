package ci

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"

	"github.com/dustin/go-humanize"
	"github.com/logrusorgru/aurora"
	"github.com/wagoodman/dive/image"
)

const (
	RuleUnknown = iota
	RulePassed
	RuleFailed
	RuleWarning
	RuleDisabled
)

type Rule interface {
	Key() string
	Configuration() string
	Evaluate(*image.AnalysisResult) (RuleStatus, string)
}

type GenericCiRule struct {
	key         string
	configValue string
	evaluator   func(*image.AnalysisResult, string) (RuleStatus, string)
}

type RuleStatus int

type RuleResult struct {
	status  RuleStatus
	message string
}

func newGenericCiRule(key string, configValue string, evaluator func(*image.AnalysisResult, string) (RuleStatus, string)) *GenericCiRule {
	return &GenericCiRule{
		key:         key,
		configValue: configValue,
		evaluator:   evaluator,
	}
}

func (rule *GenericCiRule) Key() string {
	return rule.key
}

func (rule *GenericCiRule) Configuration() string {
	return rule.configValue
}

func (rule *GenericCiRule) Evaluate(result *image.AnalysisResult) (RuleStatus, string) {
	return rule.evaluator(result, rule.configValue)
}

func (status RuleStatus) String() string {
	switch status {
	case RulePassed:
		return "PASS"
	case RuleFailed:
		return aurora.Bold(aurora.Inverse(aurora.Red("FAIL"))).String()
	case RuleWarning:
		return aurora.Blue("WARN").String()
	case RuleDisabled:
		return aurora.Blue("SKIP").String()
	default:
		return aurora.Inverse("Unknown").String()
	}
}

func loadCiRules(config *viper.Viper) []Rule {
	var rules = make([]Rule, 0)
	var ruleKey = "lowestEfficiency"
	rules = append(rules, newGenericCiRule(
		ruleKey,
		config.GetString(fmt.Sprintf("rules.%s", ruleKey)),
		func(analysis *image.AnalysisResult, value string) (RuleStatus, string) {
			lowestEfficiency, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return RuleFailed, fmt.Sprintf("invalid config value ('%v'): %v", value, err)
			}
			if lowestEfficiency > analysis.Efficiency {
				return RuleFailed, fmt.Sprintf("image efficiency is too low (efficiency=%v < threshold=%v)", analysis.Efficiency, lowestEfficiency)
			}
			return RulePassed, ""
		},
	))

	ruleKey = "highestWastedBytes"
	rules = append(rules, newGenericCiRule(
		ruleKey,
		config.GetString(fmt.Sprintf("rules.%s", ruleKey)),
		func(analysis *image.AnalysisResult, value string) (RuleStatus, string) {
			highestWastedBytes, err := humanize.ParseBytes(value)
			if err != nil {
				return RuleFailed, fmt.Sprintf("invalid config value ('%v'): %v", value, err)
			}
			if analysis.WastedBytes > highestWastedBytes {
				return RuleFailed, fmt.Sprintf("too many bytes wasted (wasted-bytes=%v > threshold=%v)", analysis.WastedBytes, highestWastedBytes)
			}
			return RulePassed, ""
		},
	))

	ruleKey = "highestUserWastedPercent"
	rules = append(rules, newGenericCiRule(
		ruleKey,
		config.GetString(fmt.Sprintf("rules.%s", ruleKey)),
		func(analysis *image.AnalysisResult, value string) (RuleStatus, string) {
			highestUserWastedPercent, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return RuleFailed, fmt.Sprintf("invalid config value ('%v'): %v", value, err)
			}
			if highestUserWastedPercent < analysis.WastedUserPercent {
				return RuleFailed, fmt.Sprintf("too many bytes wasted, relative to the user bytes added (%%-user-wasted-bytes=%v > threshold=%v)", analysis.WastedUserPercent, highestUserWastedPercent)
			}

			return RulePassed, ""
		},
	))

	return rules
}
