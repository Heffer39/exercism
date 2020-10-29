// Package wordcount counts the occurrences of each word in a phrase
package wordcount

import (
	"strings"
	"unicode"
)

//Frequency counts the number of occurrences for a given word
type Frequency map[string]int

// WordCount parses a phrase to count each of the numbers, simple words, or contractions in a phrase
// The count is case insensitive and unordered
// Other than contractions, all punctuation is ignored
// Separation can be any form of whitespace
func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	words := strings.FieldsFunc(phrase, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	})
	freq := make(Frequency, len(words))
	for _, v := range words {
		v = strings.Trim(v, "'")
		freq[v]++
	}
	return freq
}
