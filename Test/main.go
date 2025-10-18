package main

import (
	"fmt"
	"strings"
)

func main() {
	// Both declarations create identical string variables
	var str1 string = "Hello World"
	str2 := "Hello World"

	// Both work the same with strings.Index
	fmt.Printf("str1 type: %T, index: %d\n", str1, strings.Index(str1, "World"))
	fmt.Printf("str2 type: %T, index: %d\n", str2, strings.Index(str2, "World"))
}
