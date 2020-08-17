// Package luhn determines whether or not a number is valid per
// the Luhn formula
package luhn

import (
	"strings"
)

// Valid performs a simple checksum to validate an identification number
// using the Luhn algorithm
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	var sum int

	for i := range input {
		// Iterating through runes in the input slice from right to left
		r := input[len(input)-1-i]
		// Check to see if the rune is not a digit; only digits are allowed
		if r < '0' || r > '9' {
			return false
		}
		n := int(r - '0')
		if i%2 != 0 {
			n = n * 2
			if n > 9 {
				n = n - 9
			}
		}
		sum += n
	}

	return sum%10 == 0
}
