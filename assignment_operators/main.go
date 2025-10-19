package main

import "fmt"

func main() {

	// Assignment Operators
	// = , += , -= , *=, /= , %= , &=, ^=, <<= , >>=, := , |= , ==
	var var1 int = 10
	fmt.Printf("The Original value of var is: %v\n", var1)
	var1 += 5
	fmt.Printf("The value after += 5 is: %v\n", var1)
	var1 -= 3
	fmt.Printf("The value after -= 3 is: %v\n", var1)
	var1 *= 2
	fmt.Printf("The value after *= 2 is: %v\n", var1)
	var1 /= 4
	fmt.Printf("The value after /= 4 is: %v\n", var1)
	var1 %= 3
	fmt.Printf("The value after %%= 3 is: %v\n", var1)

	var2 := 20
	fmt.Printf("The Original value of var2 is: %v\n", var2)
	var2 &= 15
	fmt.Printf("The value after &= 15 is: %v\n", var2)
	var2 ^= 5
	fmt.Printf("The value after ^= 5 is: %v\n", var2)
	var2 |= 10
	fmt.Printf("The value after |= 10 is: %v\n", var2)
	var2 <<= 2
	fmt.Printf("The value after <<= 2 is: %v\n", var2)
	var2 >>= 1
	fmt.Printf("The value after >>= 1 is: %v\n", var2)

}
