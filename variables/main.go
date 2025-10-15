package main

import (
	"fmt"
)

func main() {

	var myString string = "This is my String Variable"
	fmt.Println(myString)

	var myString2 string
	myString2 = "This is the Second String Variable"
	fmt.Println(myString2)

	myString3 := "This is the Third String Variable"
	fmt.Println(myString3)

	myString4, myString5 := "This is the Fourth String Variable", "This is the Fifth String Variable"
	fmt.Printf("This is the Fourth var %v and this is the Fifth var %v", myString4, myString5)

	var myInt1, myInt2 = 43, 56
	fmt.Printf("\nThis is the First Int var %v and this is the Second Int var %v\n", myInt1, myInt2)

	var str1, str2 = "Hello", "World"
	fmt.Printf("%v %v\n", str1, str2)

	var int1, int2 int
	int1, int2 = 23, 45
	fmt.Printf("This is the First Int var %v and this is the Second Int var %v\n", int1, int2)

	Mynum := 44
	fmt.Println(Mynum)

	var myFloat1 float32 = 3.14
	var myFloat2 float64 = 3.14159265359
	fmt.Printf("This is the First Float var %v and this is the Second Float var %v\n", myFloat1, myFloat2)

	var mybool1, mybool2 = true, false
	fmt.Printf("This is the First Bool var %v and this is the Second Bool var %v\n", mybool1, mybool2)

	myBool1 := true
	fmt.Printf("This is the Bool var %v\n", myBool1)

	mySlice := []string{"Here", "Iam", "Learning", "A", "New", "Language"}
	fmt.Printf("Values is : %v\nTypes is: %T\n\n", mySlice, mySlice)

	var mySlice3 = []string{"Here", "Iam", "Learning", "A", "New", "Language"}
	fmt.Printf("Values is : %v\nTypes is: %T\n\n", mySlice3, mySlice3)

	mySlice2 := []int{23, 45, 67, 89, 90}
	fmt.Printf("Values is : %v\nTypes is: %T\n\n", mySlice2, mySlice2)

}
