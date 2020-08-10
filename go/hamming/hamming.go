// The hamming package implements methods for determining the
// hamming distance, or number of differences, between two DNA strands

package hamming

import (
	"errors"
)

// Distance function returns the number of differences between two DNA strands
// (strings) of equal length. It returns an error if the strings are not
// of equal length
func Distance(a, b string) (int, error) {
	hammingDistance := 0

	if len(a) != len(b) {
		return 0, errors.New("the strings are not equal in length")
	}

	runesA := []rune(a)
	runesB := []rune(b)

	for i := range a {
		if runesA[i] != runesB[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
