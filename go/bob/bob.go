// Package bob provides response options for conversations with bob
package bob

import (
	"strings"
	"unicode"
)

// Hey provides a response from bob based on the string passed in to the method
func Hey(remark string) (bobResponse string) {
	remark = strings.Trim(remark, " \t\n\r")

	var containsLetters bool
	for _, r := range remark {
		if unicode.IsLetter(r) {
			containsLetters = true
			break
		}
	}

	isAQuestion := strings.HasSuffix(remark, "?")
	containsAllUppercaseLetters := strings.ToUpper(remark) == remark && containsLetters

	if isAQuestion && containsAllUppercaseLetters {
		bobResponse = "Calm down, I know what I'm doing!"
	} else if isAQuestion {
		bobResponse = "Sure."
	} else if containsAllUppercaseLetters {
		bobResponse = "Whoa, chill out!"
	} else if remark == "" {
		bobResponse = "Fine. Be that way!"
	} else {
		bobResponse = "Whatever."
	}

	return bobResponse
}
