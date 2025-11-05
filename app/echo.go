package main

import (
	"fmt"
	"os"
)

func echoCommand(args []string) {
	out := ""

	for _, arg := range args {
		// get the string within quotes
		out += parseString(arg, '\'')
	}

	fmt.Fprintf(os.Stdout, "%s\n", out)
}

func parseString(value string, quoteChar byte) string {
	result := ""

	if len(value) <= 2 {
		return result
	}

	// if the first and last char are the quote char remove the quotes
	if value[0] == quoteChar && value[len(value)-1] == quoteChar {
		result = value[1 : len(value)-1]
	}

	return result
}