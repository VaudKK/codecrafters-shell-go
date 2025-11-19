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

func splitWithQuotes(s string) []string {
	var result []string
	var current string
	inQuote := false

	for i := 0; i < len(s); i++ {
		if s[i] == '\'' {
			inQuote = !inQuote
			// current += string(s[i])
		} else if s[i] == ' ' && !inQuote {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(s[i])
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}