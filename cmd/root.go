package cmd

import (
	"os"
	"os/exec"
	"time"
)

func GetDefaultShell() string {
	shell := os.Getenv("SHELL")

	if len(shell) == 0 {
		return "sh"
	}

	return shell
}

func ExecuteCommand(commandString string) int {
	cmd := exec.Command(GetDefaultShell(), "-c", commandString)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	return cmd.ProcessState.ExitCode()
}

func Run(config Config) {

	for {
		code := ExecuteCommand(config.CommandString)

		if code == config.ExitCode {
			break
		}

		time.Sleep(time.Millisecond * time.Duration(config.Timeout))
	}
}
