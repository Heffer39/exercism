// The triangle package contains functionality that supports determining
// if a shape is a triangle, as well as what type of triangle it may be.
package triangle

import (
	"math"
)

type Kind string

const (
    NaT Kind = "Not a Triangle" // not a triangle
    Equ Kind = "Equilateral" // equilateral
    Iso Kind = "Isosceles" // isosceles
    Sca Kind = "Scalene" // scalene
)

// KindFromSides returns a Kind type, which is a string representation
// of the type of Triangle (or lack thereof) based on the three length
// values (float64) passed in.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	//fmt.Printf("%g %g %g\n", a, b, c)

	// All sides must have a length greater than 0.
	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if math.IsNaN(a + b + c) || math.IsInf(a + b + c, 0) {
		// All sides must have a length greater than 0.
		k = NaT
	} else if a + b < c || a + c < b || b + c < a {
		// The sum of the lengths of any two sides must be greater
		// than or equal to the length of the third side.
		k = NaT
	} else if a == b && b == c && c == a {
		// An equilateral triangle has all three sides the same length.
		k = Equ
	} else if a == b || b == c || c == a {
		// An isosceles triangle has at least two sides the same length.
		k = Iso
	} else {
		// A scalene triangle has all sides of different lengths.
		k = Sca
	}

	return k
}
