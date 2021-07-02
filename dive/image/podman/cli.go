// +build linux

package podman

import (
	"fmt"
	"github.com/wagoodman/dive/utils"
	"io"
	"os"
	"os/exec"
)

// runPodmanCmd runs a given Podman command in the current tty
func runPodmanCmd(cmdStr string, args ...string) error {
	if !isPodmanClientBinaryAvailable() {
		return fmt.Errorf("cannot find podman client executable")
	}

	allArgs := utils.CleanArgs(append([]string{cmdStr}, args...))

	cmd := exec.Command("podman", allArgs...)
	cmd.Env = os.Environ()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func streamPodmanCmd(args ...string) (error, io.Reader) {
	if !isPodmanClientBinaryAvailable() {
		return fmt.Errorf("cannot find podman client executable"), nil
	}

	cmd := exec.Command("podman", utils.CleanArgs(args)...)
	cmd.Env = os.Environ()

	reader, writer, err := os.Pipe()
	if err != nil {
		return err, nil
	}
	defer writer.Close()

	cmd.Stdout = writer
	cmd.Stderr = os.Stderr

	return cmd.Start(), reader
}

func isPodmanClientBinaryAvailable() bool {
	_, err := exec.LookPath("podman")
	return err == nil
}
