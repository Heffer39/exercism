// Package etl provides functionality to perform an "Extract-Transform-Load" operation
package etl

import (
	"strings"
)

// Transform takes Scrabble scores from a legacy system and converts them into a new format
func Transform(etl map[int][]string) map[string]int {
	transformedMap := make(map[string]int)

	for score, words := range etl {
		for _, word := range words {
			transformedMap[strings.ToLower(word)] = score
		}
	}

	return transformedMap
}
