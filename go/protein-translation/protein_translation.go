// Package protein contains methods that translate RNA sequences into proteins
package protein

import (
	"errors"
)

// ErrStop handles STOP codons when checking individual codons with FromCodon
var ErrStop error = errors.New("error stop")

// ErrInvalidBase handles codons with invalid base characters
var ErrInvalidBase error = errors.New("error invalid base")

var codonToProteinMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon translates a valid codon string into a single protein
func FromCodon(codon string) (string, error) {
	if protein, ok := codonToProteinMap[codon]; !ok {
		return protein, ErrInvalidBase
	} else if protein == "STOP" {
		return "", ErrStop
	} else {
		return protein, nil
	}
}

// FromRNA translates a valid RNA string into a collection of proteins
func FromRNA(rna string) (proteinSlice []string, e error) {
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		if protein, ok := codonToProteinMap[codon]; !ok {
			return proteinSlice, ErrInvalidBase
		} else if protein == "STOP" {
			break
		} else {
			proteinSlice = append(proteinSlice, protein)
		}
	}
	return proteinSlice, nil
}
