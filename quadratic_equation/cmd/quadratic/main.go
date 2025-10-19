package main

import (
	"fmt"
	"quadratic_equation/quadratic"
)

func main() {
	var a, b, c float64

	fmt.Println("Enter coefficients for quadratic equation (ax² + bx + c = 0)")

	fmt.Print("Enter a: ")
	fmt.Scan(&a)

	if a == 0 {
		fmt.Println("'a' cannot be 0 in a quadratic equation")
		return
	}

	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	fmt.Print("Enter c: ")
	fmt.Scan(&c)

	x1, x2, hasRealRoots := quadratic.SolveQuadratic(a, b, c)

	if hasRealRoots {
		if x1 == x2 {
			fmt.Printf("The equation has one real root: x = %.2f\n", x1)
		} else {
			fmt.Printf("The equation has two real roots:\n")
			fmt.Printf("x₁ = %.2f\n", x1)
			fmt.Printf("x₂ = %.2f\n", x2)
		}
	} else {
		fmt.Println("The equation has no real roots (complex roots exist)")
	}
}
