// Package luhn determines whether or not a number is valid per
// the Luhn formula
package luhn

import (
	"fmt"
	"strconv"
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
		if n, err := strconv.Atoi(string(input[len(input)-1-i])); err == nil {
			if i%2 != 0 {
				n = n * 2
				if n > 9 {
					n = n - 9
				}
			}
			sum += n
		} else {
			fmt.Println("not a number")
			return false
		}
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
