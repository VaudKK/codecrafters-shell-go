package main

import (
	"fmt"
	"os"
)

var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	fmt.Fprint(os.Stdout, "$ ")
}
