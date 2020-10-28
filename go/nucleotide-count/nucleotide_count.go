package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, v := range d {
		if _, ok := h[v]; !ok {
			return h, fmt.Errorf("invalid nucleotide")
		}
		h[v]++
	}
	return h, nil
}
