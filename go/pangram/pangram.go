// Package pangram determine if a sentence is a pangram
package pangram

import (
	"strings"
)

// IsPangram returns true if the given sentence is a pangram
func IsPangram(sentence string) bool {
	/* Map approach is O(n) for speed, but is much slower for the benchmark cases
	var pangramCheck = make(map[rune]bool, 26)
	s := strings.ToLower(sentence)
	for _, v := range s {
		if unicode.IsLetter(v) {
			pangramCheck[v] = true
		}
	}
	return len(pangramCheck) == 26
	*/
	// ContainsRune approach is O(n^2) for speed, but is much faster for the benchmark cases
	s := strings.ToLower(sentence)
	for r := 'a'; r <= 'z'; r++ {
		if !strings.ContainsRune(s, r) {
			return false
		}
	}
	return true
}
