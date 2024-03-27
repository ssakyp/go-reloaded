package main

import (
	"fmt"
	"go-reloaded/converters"
	"go-reloaded/handlers"
	"go-reloaded/words"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Invalid input")
		return
	}
	if os.Args[1] != "sample.txt" || !strings.Contains(os.Args[2], ".txt") {
		fmt.Println("Invalid input. Please enter go run . sample.txt result.txt")
		return
	}
	sampleFile := os.Args[1]

	sampleByteSlice, _ := os.ReadFile(sampleFile)

	sampleStr := converters.ByteSliceToString(sampleByteSlice)

	sampleRuneSlice := words.Split(sampleStr)

	sampleStr = handlers.Editor(sampleRuneSlice) // HERE

	fileResultName := os.Args[2]
	// Write the string to the file
	err := os.WriteFile(fileResultName, []byte(sampleStr), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}

//apos 
