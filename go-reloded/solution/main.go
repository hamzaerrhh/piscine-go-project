package main

import (
	"fmt"
	"go_reloaded/helper"
	"os"
)


func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	
 //protect file 
    helper.ProtectLayer(inputFile,outputFile)

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	inputText := string(data)

	// Tokenize the text
	tokens := helper.Tokenize(inputText)
	
	// Process the tokens
	processedTokens := helper.ProcessTokens(tokens)
	
	// Convert back to string
	result := helper.TokensToString(processedTokens)
	
	// Final cleanup
	result = helper.CleanText(result)
	
	// Write to output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Successfully processed %s -> %s\n", inputFile, outputFile)
}

