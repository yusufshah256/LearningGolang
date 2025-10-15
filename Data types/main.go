package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {

	//Boolean
	//True or False Values
	//Ones or Zeros

	//Boolean
	myBool := false
	fmt.Printf("Value is : %v\nType is :%T\n\n", myBool, myBool)

	//Integer = int
	randomNumber := rand.IntN(100)
	myInt := 43 * randomNumber
	fmt.Printf("Value is : %v\nType is: %T\n\n", myInt, myInt)

	//strings = str
	myStr := "Here iam learning a new language" + "?.and this is an added string to the original string"
	fmt.Printf("Value is : %v\nType is: %T\n\n", myStr, myStr)

	//floating = float
	//calculating the area of the circle using the formula area = pi * r^2
	radius := rand.Float64() * (10.0)
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("The Area of the Circle of radius %.2f :\nis : %.2f\nType is: %T\n\n", radius, area, area)

	//Calculate the Circumference of the circle using the formula circumference = 2 * pi * r
	circumference := 2 * math.Pi * radius
	fmt.Printf("The Circumference of the Circle with Radius %.2f :\nis : %.2f\nType is: %T\n\n", radius, circumference, circumference)

	//[]string = Array or Slice
	var myString = []string{"Here", "Iam", "Learning", "A", "New", "Language"}
	fmt.Printf("Values is : %v\nTypes is: %T\n\n", myString, myString)

}
