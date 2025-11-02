package main

import (
	"bufio"
	"fmt"
	"os"
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

		fmt.Fprint(os.Stdout,command[:len(command)-1] + ": command not found\n")
	}
	
}
