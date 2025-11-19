package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runPathCommand(command string, args []string, out *os.File, errOut *os.File) {

	fmt.Println(args)

	cmd := exec.Command(command, args...)

	if out != nil {
		cmd.Stdout = out
	}

	if errOut != nil {
		cmd.Stderr = errOut
	}

	cmd.Run()
}