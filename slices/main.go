package main

import "fmt"

func main() {

	mySlice1 := []string{"One", "Two", "Three"}
	var mySlice2 = []string{"Four", "Three", "Two", "One"}
	var mySlice3 []string = []string{"Third ", "way", "To", "Create", "Slice"}
	fmt.Println(mySlice1)
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)

	mySlice1 = append(mySlice1, "Four")
	fmt.Printf("The Slice after appending %s\n", mySlice1)
}
