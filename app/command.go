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
	inSingleQuote := false
	inDoubleQuote := false

	for i := 0; i < len(s); i++ {

		if s[i] == '"' {
			inDoubleQuote = !inDoubleQuote
		}

		if s[i] == '\'' {
			if inDoubleQuote {
				result = append(result, string(s[i]))
			}else{
				inSingleQuote = !inSingleQuote
			}
		} else if s[i] == ' ' && !inSingleQuote {
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