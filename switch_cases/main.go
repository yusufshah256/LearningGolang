package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Seed the random number generator to get different results each time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 1 and 10 and store it in randNumber
	randNumber := r.Intn(10) + 1 // Random number between 1 and 10

	fmt.Println("Random number is : ", randNumber)

	switch randNumber {

	case 1:
		fmt.Println("The number is One")
	case 2:
		fmt.Println("The number is Two")
	case 3:
		fmt.Println("The number is Three")
	case 4:
		fmt.Println("The number is Four")
	case 5:
		fmt.Println("The number is Five")
	case 6:
		fmt.Println("The number is Six")
	case 7:
		fmt.Println("The number is Seven")
	case 8:
		fmt.Println("The number is Eight")
	case 9:
		fmt.Println("The number is Nine")
	case 10:
		fmt.Println("The number is Ten")
	default:
		fmt.Println("Number is out of range")
	}

}
