// Package isbn validates book identification numbers
package isbn

import (
	"unicode"
)

// IsValidISBN validates if the given string adheres to the ISBN-10 verification process
func IsValidISBN(isbn string) bool {
	sum, count := 0, 10
	for _, r := range isbn {
		if unicode.IsNumber(r) {
			sum += int(r-'0') * count
			count--
		}
		if r == 'X' && count == 1 {
			sum += 10 * count
			count--
		}
	}
	return sum%11 == 0 && count == 0
}
