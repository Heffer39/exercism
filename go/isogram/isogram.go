// Package isogram contains functionality that determines if a string
// is or is not an isogram
package isogram

import (
	"unicode"
)

// IsIsogram determines if the given word is an isogram
// An isogram is a word or phrase without a repeating letter
// Spaces and hyphens are considered exceptions
func IsIsogram(word string) bool {
	// The isogramMap keeps track of the total number of occurrences for
	// each rune in the word string
	repeated := make(map[rune]bool)

	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}
		// If the current rune has previously been added to the isogramMap,
		// then it is a repeating character, which means the word is not an isogram
		lowercaseR := unicode.ToLower(r)
		if repeated[lowercaseR] {
			return false
		}
		repeated[lowercaseR] = true
	}

	return true
}
