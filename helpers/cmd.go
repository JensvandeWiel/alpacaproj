package helpers

import (
	"os"
	"os/exec"
)

func RunCommand(wd string, toStdOut bool, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = wd
	if toStdOut {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
