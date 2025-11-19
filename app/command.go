package main

import (
	"os"
	"os/exec"
	"strings"
)

func RunPathCommand(command string, args []string, out *os.File, errOut *os.File) {
	rejoined := strings.Join(args," ");
	cmd := exec.Command(command, rejoined)

	if out != nil {
		cmd.Stdout = out
	}

	if errOut != nil {
		cmd.Stderr = errOut
	}

	cmd.Run()
}