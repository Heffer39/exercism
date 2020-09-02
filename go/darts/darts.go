// Package darts allows you to play a game of darts
package darts

import (
	"math"
)

const (
	innerCircle  = 1.0
	middleCircle = 5.0
	outerCircle  = 10.0
)

// Score calculates the score for a dart throw, given the coordinates of where it lands
// on a dart board. The dart board has a total radius of 10 units, a middle circle
// with a radius of 5 units, and an inner circle with a radius of 1 unit.
func Score(x float64, y float64) (score int) {
	pythagoreanDistance := math.Sqrt((x * x) + (y * y))

	switch {
	case innerCircle >= pythagoreanDistance:
		score = 10
	case middleCircle >= pythagoreanDistance:
		score = 5
	case outerCircle >= pythagoreanDistance:
		score = 1
	default:
		score = 0
	}

	return
}
