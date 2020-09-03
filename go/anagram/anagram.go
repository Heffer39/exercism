// Package anagram can help determine if words are anagrams of each other
package anagram

import (
	"strings"
)

// Detect selects the sublist of anagrams for the given word out of the list of candidates
func Detect(subject string, candidates []string) (anagrams []string) {
	for _, candidate := range candidates {
		if compareTwoStrings(strings.ToLower(subject), strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}
	return
}

// compareTwoStrings checks two individual words to determine if they are anagrams
func compareTwoStrings(subject, candidate string) bool {
	if len(subject) != len(candidate) || subject == candidate {
		return false
	}
	subjectRuneMap := map[rune]int{}
	candidateRuneMap := map[rune]int{}

	for _, c := range subject {
		subjectRuneMap[c]++
	}
	for _, c := range candidate {
		candidateRuneMap[c]++
	}

	for c := range subjectRuneMap {
		if subjectRuneMap[c] != candidateRuneMap[c] {
			return false
		}
	}

	return true
}
