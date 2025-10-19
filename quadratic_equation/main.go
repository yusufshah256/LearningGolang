package main

import (
	"fmt"
	"math"
)

// solveQuadratic calculates the roots of a quadratic equation
// Returns real roots and a boolean indicating if roots are real
func solveQuadratic(a, b, c float64) (float64, float64, bool) {
	// Calculate discriminant
	discriminant := b*b - 4*a*c

	// Check if roots are real
	if discriminant < 0 {
		return 0, 0, false
	}

	// Calculate roots using quadratic formula
	// x = (-b ± √(b² - 4ac)) / 2a
	x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	return x1, x2, true
}

func main() {
	var a, b, c float64

	// Get coefficients from user
	fmt.Println("Enter coefficients for quadratic equation (ax² + bx + c = 0)")

	fmt.Print("Enter a: ")
	fmt.Scan(&a)

	// Check if a is zero (makes equation linear)
	if a == 0 {
		fmt.Println("'a' cannot be 0 in a quadratic equation")
		return
	}

	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	fmt.Print("Enter c: ")
	fmt.Scan(&c)

	// Solve the equation
	x1, x2, hasRealRoots := solveQuadratic(a, b, c)

	// Display results
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
