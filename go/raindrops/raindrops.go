// Package raindrops provides functionality regarding raindrops and converting
// the potential they store into sounds
package raindrops

import (
	"strconv"
)

// Convert creates a string with the corresponding raindrop sounds,
// based on the potential factor of the passed in number
func Convert(potential int) string {

	var raindrops = ""

	if potential%3 == 0 {
		raindrops += "Pling"
	}
	if potential%5 == 0 {
		raindrops += "Plang"
	}
	if potential%7 == 0 {
		raindrops += "Plong"
	}
	if potential%3 != 0 && potential%5 != 0 && potential%7 != 0 {
		raindrops += strconv.Itoa(potential)
	}

	return raindrops

	/*
		var raindrops strings.Builder

		if potential%3 == 0 {
			raindrops.WriteString("Pling")
		}
		if potential%5 == 0 {
			raindrops.WriteString("Plang")
		}
		if potential%7 == 0 {
			raindrops.WriteString("Plong")
		}
		if potential%3 != 0 && potential%5 != 0 && potential%7 != 0 {
			raindrops.WriteString(strconv.Itoa(potential))
		}

		return raindrops.String()

	*/
}
