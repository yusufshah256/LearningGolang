package main

import (
	"fmt"
)

func main() {

	myMap := map[string]int{
		"cat":   40,
		"Dog":   50,
		"Camel": 60,
	}

	var myMap2 = map[string]int{

		"cat":   40,
		"Dog":   50,
		"Camel": 60,
	}

	var myMap1 map[string]int = map[string]int{

		"car":     66,
		"swimmer": 55,
		"man":     54,
	}

	var myMap3 = make(map[string]string)

	myMap3["Gold"] = "100$"
	myMap3["Silver"] = "10$"
	myMap3["Egg"] = "1$"

	fmt.Println(myMap)
	fmt.Println(myMap1)
	fmt.Println(myMap2)
	fmt.Println(myMap3)

	delete(myMap, "cat")
	fmt.Println(myMap)
	_, avlaibleQuery := myMap2["cat"]

	fmt.Printf("Is this element value Availble %v?: %t\n", myMap2["cat"], avlaibleQuery)

	// Iterate through  the maps elements

	for key, value := range myMap2 {
		fmt.Printf("The Key is '%v' and the value is : %v\n", key, value)
	}
}
