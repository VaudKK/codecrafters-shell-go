package main

import (
	"fmt"
	"os"
	"strings"
)

func echoCommand(args []string) {
	out := ""
	hasQuote := false

	rejoined := strings.Join(args, " ")

	if strings.ContainsAny(rejoined, `'"`) {
		hasQuote = true
	}

	if hasQuote {
		out = removeQuotes(rejoined, `'`)
		fmt.Fprintf(os.Stdout, "%s\n", out)
	} else {
		cleaned := strings.Fields(rejoined)
		out = strings.Join(cleaned, " ")
		fmt.Fprintf(os.Stdout, "%s\n", out)
	}
}

func removeQuotes(s string, quoteChar string) string {
	cleaned := strings.ReplaceAll(s, quoteChar, "")
	return cleaned
}
