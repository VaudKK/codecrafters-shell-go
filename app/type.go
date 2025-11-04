package main

import "strings"

var commands = []string{
	"exit",
	"echo",
}

func IsValidCommand(cmd string) bool {
	for _, command := range commands {
		if strings.ToLower(cmd) == command {
			return true
		}
	}

	return false
}
