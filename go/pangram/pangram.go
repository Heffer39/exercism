// Package pangram determine if a sentence is a pangram
package pangram

import (
	"unicode"
)

// IsPangram returns true if the given sentence is a pangram
func IsPangram(sentence string) bool {
	var pangramCheck = map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}

	for _, v := range sentence {
		if _, ok := pangramCheck[v]; !ok {
			if string(v) == " " || string(v) == "_" || unicode.IsNumber(v) {
				continue
			}
		}
		pangramCheck[unicode.ToLower(v)]++
	}
	for _, v := range pangramCheck {
		if v == 0 {
			return false
		}
	}
	return true
}
