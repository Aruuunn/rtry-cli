package cmd_test

import (
	"log"
	"testing"

	"github.com/arunmurugan78/rtry/cmd"
)

// very simple test
func TestExecCommand(t *testing.T) {
	code := cmd.ExecuteCommand("echo hello world")

	if code != 0 {
		log.Fatalln("expected code 0 on exit")
	}

}
