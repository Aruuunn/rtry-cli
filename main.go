package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arunmurugan78/gtry/cmd"
)

func main() {

	flag.Usage = func() {
		fmt.Println("Usage:\n\tgtry [OPTIONS] COMMAND\nExample:\n\tgtry \"echo Hello there!\"\nOptions:")
		flag.PrintDefaults()
	}

	timeout := flag.Int("timeout", 1000, "timeout after to retry in milliseconds")

	flag.Parse()

	commandLineArgs := flag.CommandLine.Args()

	if len(commandLineArgs) != 1 {
		fmt.Println("Error: COMMAND expected")
		flag.Usage()
		os.Exit(1)
	}

	cmd.Run(cmd.Config{
		Timeout:       *timeout,
		CommandString: commandLineArgs[0],
	})
}
