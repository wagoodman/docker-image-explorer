module github.com/wagoodman/dive

go 1.13

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/awesome-gocui/gocui v0.6.0
	github.com/awesome-gocui/keybinding v1.0.1-0.20190805183143-864552bd36b7
	github.com/cespare/xxhash v1.1.0
	github.com/docker/cli v0.0.0-20190906153656-016a3232168d
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.7.3-0.20190309235953-33c3200e0d16
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.7.0
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.2 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/logrusorgru/aurora v0.0.0-20190803045625-94edacc10f9b
	github.com/lunixbochs/vtclean v1.0.0
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/phayes/permbits v0.0.0-20190612203442-39d7c581d2ee
	github.com/sergi/go-diff v1.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297
	golang.org/x/sys v0.0.0-20200212091648-12a6c2dcc1e4 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20190620144150-6af8c5fc6601 // indirect
	google.golang.org/grpc v1.21.1 // indirect
	gotest.tools v2.2.0+incompatible // indirect
)

// related to an invalid pseudo version in github.com/docker/distribution@v0.0.0-20181126153310-93e082742a009850ac46962150b2f652a822c5ff
replace github.com/docker/distribution => github.com/docker/distribution v2.7.0-rc.0.0.20181024170156-93e082742a00+incompatible

// related to an invalid pseudo version in github.com/docker/distribution@v0.0.0-20181126153310-93e082742a009850ac46962150b2f652a822c5ff
replace github.com/docker/docker => github.com/docker/engine v0.0.0-20190822205725-ed20165a37b4

// relates to https://github.com/golangci/golangci-lint/issues/581
replace github.com/go-critic/go-critic => github.com/go-critic/go-critic v0.3.5-0.20190526074819-1df300866540

replace github.com/golangci/errcheck => github.com/golangci/errcheck v0.0.0-20181223084120-ef45e06d44b6

replace github.com/golangci/go-tools => github.com/golangci/go-tools v0.0.0-20190318060251-af6baa5dc196

replace github.com/golangci/gofmt => github.com/golangci/gofmt v0.0.0-20181222123516-0b8337e80d98

replace github.com/golangci/gosec => github.com/golangci/gosec v0.0.0-20190211064107-66fb7fc33547

replace github.com/golangci/ineffassign => github.com/golangci/ineffassign v0.0.0-20190609212857-42439a7714cc

replace github.com/golangci/lint-1 => github.com/golangci/lint-1 v0.0.0-20190420132249-ee948d087217

replace mvdan.cc/unparam => mvdan.cc/unparam v0.0.0-20190209190245-fbb59629db34
