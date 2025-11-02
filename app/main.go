package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var _ = fmt.Fprint
var _ = os.Stdout

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprint(os.Stderr,"Error reading input")
			return
		}

		runCommand(command)
	}
	
}


func runCommand(command string){

	commandTokens := strings.Split(command," ")

	switch commandTokens[0] {
		case "exit":
			if len(commandTokens) > 1 {
				exitCode := commandTokens[1]
				if code, err := strconv.Atoi(exitCode); err == nil {
					os.Exit(code)
				}
			}else{
				os.Exit(0)
			}
		default:
			fmt.Fprint(os.Stdout,command[:len(command)-1] + ": command not found\n")
	}
}
