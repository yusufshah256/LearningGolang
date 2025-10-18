package main

import (
	"fmt"
)

type Cars struct {
	name  string
	model int
	color string
	price int
}

func main() {
	var car1 Cars = Cars{"BMW", 543456543, "Red", 65454}
	car2 := Cars{"Hunday", 2024, "white", 65435}
	fmt.Printf("The type of this variable is %T\nit's value is :%v\n", car1, car1)

	fmt.Println(car1.name)
	fmt.Println(car1.model)
	fmt.Println(car1.color)
	fmt.Println(car1.price)

	fmt.Printf("The type of this variable is %T\nit's value is :%v\n", car2, car2)

	// printing values uing pointers

	valueRefrence := &car1
	fmt.Println("Name fo the Car is:", (*valueRefrence).name)
	fmt.Println("Name fo the Car is:", (*valueRefrence).price)
	fmt.Println("Name fo the Car is:", (*valueRefrence).color)
	fmt.Println("Name fo the Car is:", (*valueRefrence).model)
}
