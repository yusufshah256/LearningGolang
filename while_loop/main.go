package main

import (
	"fmt"
	"time"
)

func main() {

	// Go only has a for loop, which can be used as a while loop
	// Example of a while loop that counts from 1 to 5
	count := 1
	for count <= 5 {
		println("Count is:", count)
		count++
	}

	// Infinite loop example like while in other languages
	myInt := 1
	for {
		fmt.Println(myInt)
		myInt++
		time.Sleep(2 * time.Second)
		if myInt > 5 {
			break
		}
	}

	// Example to check odd or even numbers from 0 to 10 using the for loop
	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}

	// Another way to represent this loop like this

	myInt = 0
	i := 0

	for {

		if i <= 10 {
			if myInt%2 == 0 {
				fmt.Printf("%d is even\n", myInt)
			} else {
				fmt.Printf("%d is odd\n", myInt)
			}
			myInt++
			i++
		} else {
			break
		}
	}

}
