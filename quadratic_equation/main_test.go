package main

import (
	"math"
	"testing"
)

// TestSolveQuadratic runs test cases for the quadratic equation solver
func TestSolveQuadratic(t *testing.T) {
	// Define test cases
	tests := []struct {
		name         string
		a, b, c      float64
		wantX1       float64
		wantX2       float64
		wantRealRoot bool
	}{
		{
			name:         "Two distinct real roots (x-1)(x-2)",
			a:            1,
			b:            -3,
			c:            2,
			wantX1:       2,
			wantX2:       1,
			wantRealRoot: true,
		},
		{
			name:         "One repeated root (x-1)²",
			a:            1,
			b:            -2,
			c:            1,
			wantX1:       1,
			wantX2:       1,
			wantRealRoot: true,
		},
		{
			name:         "No real roots x²+1",
			a:            1,
			b:            0,
			c:            1,
			wantX1:       0,
			wantX2:       0,
			wantRealRoot: false,
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x1, x2, hasRealRoots := solveQuadratic(tt.a, tt.b, tt.c)

			// Check if hasRealRoots matches expected
			if hasRealRoots != tt.wantRealRoot {
				t.Errorf("solveQuadratic(%v, %v, %v) hasRealRoots = %v, want %v",
					tt.a, tt.b, tt.c, hasRealRoots, tt.wantRealRoot)
			}

			// If real roots expected, check values
			if tt.wantRealRoot {
				// Use small epsilon for floating point comparison
				const epsilon = 1e-10

				// Check both roots
				if math.Abs(x1-tt.wantX1) > epsilon || math.Abs(x2-tt.wantX2) > epsilon {
					t.Errorf("solveQuadratic(%v, %v, %v) = (%v, %v), want (%v, %v)",
						tt.a, tt.b, tt.c, x1, x2, tt.wantX1, tt.wantX2)
				}
			}
		})
	}
}
