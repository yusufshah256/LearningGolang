package main

import (
	"flag"
	"fmt"
)

// Flag.DataType
// Flag.String("Name" , "DefaultValue", "Description")
// Flag.Integer ("Name" , 0, "Description")
// Flag.Bool ("Name" , true/false, "Description")

func main() {

	name := flag.String("name", "Admin", "Full Name of the User")
	age := flag.Int("age", 2, "Age of the User")
	isAdmin := flag.Bool("isAdmin", true, "is the User an Admin")
	flag.Parse()

	fmt.Printf("Hello, %v\n", *name)
	fmt.Printf("You are %v years old\n", *age)

	switch *isAdmin {
	case true:
		fmt.Println("You have Admin Access")

	case false:
		fmt.Println("You are a Regular User")
	}
}
