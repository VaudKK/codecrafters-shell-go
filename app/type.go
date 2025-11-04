package main

import (
	"os"
	"os/exec"
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

		if info, err := os.Stat(fullPath); err == nil {
			if info.Mode().IsRegular() && info.Mode().Perm()&0111 != 0 {
				return true, fullPath
			}
		}
	}

	return false,""
}


func RunPathCommand(command string,args []string,out *os.File,errOut *os.File){
	cmd := exec.Command(command,args...)

	if out != nil {
		cmd.Stdout = out
	}

	if errOut != nil {
		cmd.Stderr = errOut
	}

	cmd.Run()
}
