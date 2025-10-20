package main

import "fmt"

func main() {

	// if-else statements

	/*
		if true/false {
			This block will execute
		} else {
			This block will not execute
		}

	*/

	var myInt int

	fmt.Printf("Please Enter a number to check if it a Odd or Even: ")

	fmt.Scanln(&myInt)

	if myInt == 0 {

		fmt.Printf("The number is zero\n")

	} else if myInt%2 == 0 {

		fmt.Printf("The number is even\n")

	} else {

		fmt.Printf("The number is odd\n")
	}

}
