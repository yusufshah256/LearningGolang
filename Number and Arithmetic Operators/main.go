package main

import (
	"fmt"
)

func main() {

	// "+ operations"
	A := 10
	A += 5 // A = A + 5
	result := 0.5 * float64(A)
	fmt.Printf("The Type of the variable is: %T\nand the result of the AO is: %v\n", result, result)

	// "- operations"
	R := 7.5
	R -= 1.5 // R = R - 1.5
	result2 := 2.5 * R
	fmt.Printf("The Type of the variable is: %T\nand the result of the AO is: %v\n", result2, result2)

	// "/ operations"
	C := 20
	C /= 2
	result3 := 3 * C
	fmt.Printf("The Type of the variable is: %T\nand the result of the AO is: %v\n", result3, result3)

	// "% operations"
	B := 66
	B %= 7
	result4 := 4 * B
	fmt.Printf("The Type of the variable is: %T\nand the result of the AO is: %v\n", result4, result4)

	// "* operations"
	D := 77
	D *= 2
	result5 := 5 * D
	fmt.Printf("The Type of the variable is: %T\nand the result of the AO is: %v\n", result5, result5)

}
