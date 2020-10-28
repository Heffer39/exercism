// Package pangram determine if a sentence is a pangram
package pangram

import (
	"strings"
)

// IsPangram returns true if the given sentence is a pangram
func IsPangram(sentence string) bool {
	s := strings.ToLower(sentence)
	for r := 'a'; r <= 'z'; r++ {
		if !strings.ContainsRune(s, r) {
			return false
		}
	}
	return true
}
