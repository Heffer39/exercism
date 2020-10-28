// Package romannumerals converts normal numbers to Roman Numerals
package romannumerals

import (
	"fmt"
	"strings"
)

// This map of romanNumerals stores all the mappings we need to represent Roman Numerals
// up to 3000
var romanNumerals = []romanNumeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// romanNumeral stores the association between a Roman Numeral and it's corresponding
// Arabic number
type romanNumeral struct {
	character string
	value     int
}

// ToRomanNumeral converts a Roman Numeral into it's Arabic representation as a string
func ToRomanNumeral(arabic int) (string, error) {
	if arabic > 3000 || arabic <= 0 {
		return "", fmt.Errorf("unhandled inputs")
	}
	var result strings.Builder
	for _, v := range romanNumerals {
		diff := arabic / v.value
		s := strings.Repeat(v.character, diff)
		result.WriteString(s)
		arabic = arabic % v.value
	}
	return result.String(), nil
}
