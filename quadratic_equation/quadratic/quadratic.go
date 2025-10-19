package quadratic

import "math"

// SolveQuadratic calculates the roots of a quadratic equation
// Returns real roots and a boolean indicating if roots are real
func SolveQuadratic(a, b, c float64) (float64, float64, bool) {
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return 0, 0, false
	}

	x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	return x1, x2, true
}
