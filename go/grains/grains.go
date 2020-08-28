// Package grains has two functions that help to determine the number of grains on a 64 square chessboard,
// given that the number of grains on each square doubles
package grains

import (
	"fmt"
)

// Square calculates the number of grains in a given square on a chess board
func Square(position int) (uint64, error) {
	if position <= 0 || position > 64 {
		return 0, fmt.Errorf("position %v does not fit within a 64 square chessboard", position)
	}
	return 1 << (position - 1), nil
}

// Total is a deterministic method that will compute the sum of all squares on a chessboard with 64 squares
// The answer is always equal to 18446744073709551615
// Errors are unhandled because we will never capture the "expectedError"
func Total() uint64 {
	return 1<<(64) - 1
}

/*
func Total() uint64 {
	total, err := Square(64)

	if err != nil {
		panic("error blah")
	}

	return total*2 - 1
}
*/
