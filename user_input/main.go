package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Prompt ask the user for their name
	fmt.Println("Enter Your name: ")

	// creating a new reader for the standard input

	inputReader := bufio.NewReader(os.Stdin)

	//Read the input until the newline character
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error Reading input: ", err)
		return
	}

	//Trim White spaces and the newline character from the input
	name := strings.TrimSpace(input)

	//Print The sanitized Input
	fmt.Printf("Hello, %v\n", name)
}
