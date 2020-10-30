// Package encode implements run-length encoding and decoding
package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode compresses data in a lossless format
func RunLengthEncode(data string) string {
	dataRunes := []rune(data)
	var encoded strings.Builder
	var count int

	for i := 0; i <= len(dataRunes); i++ {
		if i == 0 {
			count++
			continue
		}
		if i == len(dataRunes) || dataRunes[i-1] != dataRunes[i] {
			s := string(dataRunes[i-1])
			if count > 1 {
				s = strconv.Itoa(count) + s
			}
			encoded.WriteString(s)
			count = 0
		}
		count++
	}
	return encoded.String()
}

// RunLengthDecode reconstructs compressed data without loss
func RunLengthDecode(data string) string {
	dataRunes := []rune(data)
	var decoded strings.Builder
	var count int
	var countRunes string

	for i := 0; i < len(dataRunes); i++ {
		if len(dataRunes) == 0 {
			return ""
		}
		v := dataRunes[i]
		if unicode.IsNumber(v) {
			countRunes += string(v)
			continue
		}
		if unicode.IsLetter(v) || unicode.IsSpace(v) {
			count, _ = strconv.Atoi(countRunes)
			s := string(v)
			if count > 0 {
				s = strings.Repeat(string(v), count)
			}
			decoded.WriteString(s)
			countRunes = ""
		}
	}
	return decoded.String()
}
