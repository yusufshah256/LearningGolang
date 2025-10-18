package main

import (
	"fmt"
)

func main() {

	//Escape Sequence Characters

	// \b backspace
	fmt.Println("Hello, \b\bWorld!")

	// \f form feed
	fmt.Println("Hello, \fWorld!")

	// \n new line
	fmt.Println("Hello, \nWorld!")

	// \r carriage return
	fmt.Println("Hello, \rWorld")

	// \t tab
	fmt.Println("Hello, \tWorld!")

	// \v vertical tab
	fmt.Println("Hello, \vWorld!")

	// \\ backslash
	// 1.
	fmt.Println("Hello, \\World!")
	// 2.
	fmt.Println(`Hello, \World!`)

	// ' single quote
	fmt.Println("Hello, 'World!")

	// \" double quote
	fmt.Println("Hello, \"World!")

}
