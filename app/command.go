package main

import (
	"os"
	"os/exec"
)

func RunPathCommand(command string, args []string, out *os.File, errOut *os.File) {
	cmd := exec.Command(command, args...)

	if out != nil {
		cmd.Stdout = out
	}

	if errOut != nil {
		cmd.Stderr = errOut
	}

	cmd.Run()
}