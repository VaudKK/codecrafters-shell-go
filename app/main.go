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

	commandTokens := strings.Split(command[:len(command)-1]," ")

	switch commandTokens[0] {
		case "exit":
			if len(commandTokens) > 1 {
				exitCode := commandTokens[1]
				if code, err := strconv.Atoi(exitCode); err == nil {
					os.Exit(code)
				}else{
					os.Exit(0)
				}
			}else{
				os.Exit(0)
			}
		case "echo":
			if len(commandTokens) > 1{
				fmt.Fprintf(os.Stdout,"%s\n",strings.Join(commandTokens[1:]," "))
			}
		case "type":
			if len(commandTokens) > 1{
				if IsValidCommand(commandTokens[1]){
					fmt.Fprintf(os.Stdout,"%s is a shell builtin\n",commandTokens[1])
				}else{
					fmt.Fprintf(os.Stdout,"%s: not found\n",commandTokens[1])
				}
			}
		default:
			fmt.Fprintf(os.Stdout,"%s: not found\n",commandTokens[0])
	}
}
