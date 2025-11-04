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
				if IsBuiltIn(commandTokens[1]){
					fmt.Fprintf(os.Stdout,"%s is a shell builtin\n",commandTokens[1])
				}else if exists, filePath := IsCommandInPath(commandTokens[1]); exists {
					fmt.Fprintf(os.Stdout,"%s is %s\n",commandTokens[1],filePath)
				}else{
					fmt.Fprintf(os.Stdout,"%s: not found\n",commandTokens[1])
				}
			}
		default:
			if exists,_ := IsCommandInPath(commandTokens[0]); exists {
				RunPathCommand(commandTokens[0],commandTokens[1:],os.Stdout,os.Stderr)
			}else{
				fmt.Fprintf(os.Stdout,"%s: not found\n",commandTokens[0])
			}
	}
}
