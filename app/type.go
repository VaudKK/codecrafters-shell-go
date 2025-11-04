package main

import (
	"os"
	"runtime"
	"strings"
)

var commands = []string{
	"exit",
	"echo",
	"type",
}

func IsBuiltIn(cmd string) bool {
	for _, command := range commands {
		if strings.ToLower(cmd) == command {
			return true
		}
	}

	return false
}

func IsCommandInPath(cmd string) (bool,string){
	systemPath := strings.Split(os.Getenv("PATH"),string(os.PathListSeparator))

	for _, pathDir := range systemPath {
		fullPath := pathDir + string(os.PathSeparator) + cmd

		if runtime.GOOS == "windows" {
			fullPath += ".exe"
		}

		if _, err := os.Stat(fullPath); err == nil {
			return true,fullPath
		}
	}

	return false,""
}
