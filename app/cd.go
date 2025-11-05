package main

import (
	"fmt"
	"os"
)

func cdCommand(arg string) {
	err := os.Chdir(arg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", arg)
	}
}