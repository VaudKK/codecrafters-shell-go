package main

import (
	"fmt"
	"os"
	"strconv"
)

func echoCommand(args []string) {
	out := ""

	for _, arg := range args {
		unQuoted, err := strconv.Unquote(arg)

		if err != nil {
			out += arg
		}else{
			out += unQuoted
		}
	}

	fmt.Fprintf(os.Stdout, "%s\n", out)
}