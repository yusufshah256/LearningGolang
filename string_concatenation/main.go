package main

import (
	"fmt"
)

func main() {

	A := "John"
	B := " Doe"

	A += B // A = A + B
	fmt.Printf("The merged string is: %v\n", A)

	C := "Hello,"
	D := "World!\n"

	fmt.Printf("The Concatenating String is: %v %v", C, D)
}
