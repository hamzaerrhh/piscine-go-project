package main

import (
	"fmt"
	"os"

	"go_reloaded/helper"
)

func main() {
	// get the file result
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		os.Exit(0)
	}
	tokens := helper.Tokenize(string(data))

	fmt.Print(tokens)
}
