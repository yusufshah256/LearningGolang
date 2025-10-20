package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	// Type Conversion in Go

	// int8 int16 int32 int64
	var var1 int8 = 127
	fmt.Printf("The value is : %v\nThe Type is : %T\n", var1, var1)

	var var2 int16 = int16(var1)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", var2, var2)

	var var3 int32 = int32(var2)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", var3, var3)

	var var4 int64 = int64(var3)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", var4, var4)

	c := int(var4)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", c, c)

	// int --> float32 --> float64

	var intVar int = 654323456

	var floatVar32 float32 = float32(intVar)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", floatVar32, floatVar32)
	var floatVar64 float64 = float64(intVar)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", floatVar64, floatVar64)

	// float64 --> float32 --> int

	var floatVarA float64 = 12345.6789

	var float64Intoint int = int(floatVarA)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", float64Intoint, float64Intoint)

	var float32Var float32 = float32(floatVarA)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", float32Var, float32Var)

	var float32ToInt int = int(float32Var)
	fmt.Printf("The value is : %v\nThe Type is : %T\n", float32ToInt, float32ToInt)

	intVar1 := 5 / 2
	// int divided by int results in int
	fmt.Printf("The value is : %v\nThe Type is of result : %T\n", intVar1, intVar1)

	intVar2 := 5.0 / 2

	// float divided by int results in float
	fmt.Printf("The value is : %v\nThe Type is of result : %T\n", intVar2, intVar2)

	//  int => string
	intvar3 := 65

	fmt.Printf("The value is: %v\nThe Type is: %T\n", strconv.Itoa(intvar3), strconv.Itoa(intvar3))

	// string => int

	stringVar1 := "1234"

	intVar4, error := strconv.Atoi(stringVar1)
	if error != nil {

		log.Fatal(error)

	}

	fmt.Printf("The value is: %v\nThe Type is: %T\n", intVar4, intVar4)

	stringVar2 := "This is my string"

	fmt.Printf("The value is: %v\nThe Type is: %T\n", stringVar2, stringVar2)

	// string => byte slice

	byteSlice := []byte(stringVar2)
	fmt.Printf("The value is in byte slices: %v\nThe Type is: %T\n", byteSlice, byteSlice)

	// byte slice => string
	stringVar3 := string(byteSlice)
	fmt.Printf("The value is in string from byte slice is : %v\nThe Type is: %T\n", stringVar3, stringVar3)

}
