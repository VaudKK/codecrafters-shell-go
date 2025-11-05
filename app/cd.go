package main

import "os"

func cdCommand(arg string) {
	err := os.Chdir(arg)

	if err != nil {
		return
	}
}