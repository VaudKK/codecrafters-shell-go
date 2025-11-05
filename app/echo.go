package main

import (
	"fmt"
	"os"
	"strings"
)

func echoCommand(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, ""))
}