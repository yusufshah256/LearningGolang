package main

import "fmt"

func main() {

	var array1 [3]string = [3]string{"A", "B", "C"}
	var array2 = [3]string{"D", "E", "F"}
	array3 := [3]string{"G", "H", "J"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	elementToDelete := "B"
	myNewArray := []string{}

	for _, x := range array1 {
		if x != elementToDelete {
			myNewArray = append(myNewArray, x)
		}
	}
	array1[2] = "H"

	fmt.Println(myNewArray)
	fmt.Println(array1)
}
