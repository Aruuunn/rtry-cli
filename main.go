package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arunmurugan78/rtry/cmd"
)

const (
	version = "v1.1.0"
)

func main() {

	flag.Usage = func() {
		fmt.Println("Usage:\n\trtry [OPTIONS] COMMAND\nExample:\n\trtry \"ping google.com\"\nOptions:")
		flag.PrintDefaults()
	}

	timeout := flag.Int("timeout", 1000, "timeout after to retry in milliseconds")

	flag.Parse()

	commandLineArgs := flag.CommandLine.Args()

	if len(commandLineArgs) == 0 {
		fmt.Println("rtry ", version)
		fmt.Println("Description: (Re-)Try executing a command till it succeeds. ")
		fmt.Println("Author: Arun Murugan")
		flag.Usage()
		os.Exit(0)
	}

	cmd.Run(cmd.Config{
		Timeout:       *timeout,
		CommandString: Join(commandLineArgs),
	})

}

func Join(args []string) string {
	r := ""

	for _, arg := range args {
		r = r + " " + arg
	}

	return r
}
