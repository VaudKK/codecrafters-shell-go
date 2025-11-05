package main

import "os"

func pwdCommand() string {
	currentDir, err := os.Getwd()

	if err != nil {
		return ""
	}

	return currentDir
}