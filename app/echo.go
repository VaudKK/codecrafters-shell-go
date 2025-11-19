package main

import (
	"fmt"
	"os"
	"strings"
)

func echoCommand(args []string) {
	fmt.Fprintf(os.Stdout,"%s\n", cleanInput(args))
}

func cleanInput(args []string) string {
	hasQuote := false

	rejoined := strings.Join(args, " ")

	if strings.ContainsAny(rejoined, `'`) {
		hasQuote = true
	}

	if hasQuote {
		return removeQuotes(rejoined, `'`)
	} else {
		cleaned := strings.Fields(rejoined)
		return strings.Join(cleaned, " ")
	}
}

func removeQuotes(s string, quoteChar string) string {
	cleaned := strings.ReplaceAll(s, quoteChar, "")
	return cleaned
}
