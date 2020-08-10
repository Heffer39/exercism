// Package scrabble contains functions that assist with scoring
// for the scrabble game
package scrabble

import (
	"unicode"
)

// The scrabbleMap holds all of the characters and their associated
// scores
var scrabbleMap = map[rune]int{
	'a': 1,
	'e': 1,
	'i': 1,
	'o': 1,
	'u': 1,
	'l': 1,
	'n': 1,
	'r': 1,
	's': 1,
	't': 1,
	'd': 2,
	'g': 2,
	'b': 3,
	'c': 3,
	'm': 3,
	'p': 3,
	'f': 4,
	'h': 4,
	'v': 4,
	'w': 4,
	'y': 4,
	'k': 5,
	'j': 8,
	'x': 8,
	'q': 10,
	'z': 10,
}

// Score calculates the total score for a given word,
// based on the score of each letter in the word
func Score(scrabble string) int {
	var score = 0

	for _, letterRune := range scrabble {
		var lowercaseLetter = unicode.ToLower(letterRune)
		if _, ok := scrabbleMap[lowercaseLetter]; ok {
			score += scrabbleMap[lowercaseLetter]
		}
	}

	return score
}
