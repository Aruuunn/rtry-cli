package cmd

import (
	"flag"
	"fmt"
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

func Execute(version string) {
	flag.Usage = func() {
		fmt.Println("Usage:\n\trtry [OPTIONS] COMMAND\nExample:\n\trtry \"ping google.com\"\nOptions:")
		flag.PrintDefaults()
	}

	timeout := flag.Int("timeout", 1000, "timeout after to retry in milliseconds")
	exitCode := flag.Int("code", 0, "expected exit code to stop")
	flag.Parse()

	commandLineArgs := flag.CommandLine.Args()

	if len(commandLineArgs) == 0 {
		fmt.Println("rtry ", version)
		fmt.Println("Description: (Re-)Try executing a command till it succeeds. ")
		fmt.Println("Author: Arun Murugan")
		flag.Usage()
		os.Exit(0)
	}

	Run(Config{
		Timeout:       *timeout,
		CommandString: join(commandLineArgs),
		ExitCode:      *exitCode,
	})
}

func join(args []string) string {
	r := ""

	for _, arg := range args {
		r = r + " " + arg
	}

	return r
}
