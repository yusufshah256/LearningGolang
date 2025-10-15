package main

import (
	"fmt"
)

func main() {

	A := 1
	B := 4

	fmt.Println(A < B)  // true
	fmt.Println(A <= B) // true
	fmt.Println(A > B)  // false
	fmt.Println(A >= B) // false
	fmt.Println(A == B) // false
	fmt.Println(A != B) // true
}
