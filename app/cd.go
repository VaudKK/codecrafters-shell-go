package main

import (
	"fmt"
	"os"
)

func cdCommand(arg string) {

	if arg == "~" {
		homeDir,err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", arg)
			return
		}
		arg = homeDir
	}

	err := os.Chdir(arg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", arg)
	}
}