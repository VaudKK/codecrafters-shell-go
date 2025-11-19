package main

import (
	"os"
	"os/exec"
	"strings"
)

func runPathCommand(command string, args []string, out *os.File, errOut *os.File) {
	cmd := exec.Command(command, splitWithQuotes(strings.Join(args," "))...)

	if out != nil {
		cmd.Stdout = out
	}

	if errOut != nil {
		cmd.Stderr = errOut
	}

	cmd.Run()
}