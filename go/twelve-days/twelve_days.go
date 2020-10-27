// Package twelve outputs the lyrics to 'The Twelve Days of Christmas'
package twelve

import "strings"

var christmasGifts = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

var dayNumbers = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

// Song returns a string of all twelve verses, for each day in the twelve days of christmas song
func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}

// Verse returns a string of the verse for a specific day in the twelve days of christmas song
func Verse(n int) string {
	var verse strings.Builder
	beginning := "On the " + dayNumbers[n] + " day of Christmas my true love gave to me: "
	verse.WriteString(beginning)
	for i := n; i > 0; i-- {
		if i == 1 {
			finalString := christmasGifts[i] + "."
			if n > 1 {
				finalString = "and " + finalString
			}
			verse.WriteString(finalString)
			break
		}
		day := christmasGifts[i] + ", "
		verse.WriteString(day)
	}
	return verse.String()
}
