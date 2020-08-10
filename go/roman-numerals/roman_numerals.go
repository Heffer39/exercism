package romannumerals

var romanNumeralsMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func ToRomanNumeral(arabic int) (string, bool){
	var romanNumeral = ""
	for arabic > 0 {

		arabic, romanNumeral = romanNumeralCount(arabic, 1000)
		arabic, romanNumeral = romanNumeralCount(arabic, 500)
		arabic, romanNumeral = romanNumeralCount(arabic, 100)
		arabic, romanNumeral = romanNumeralCount(arabic, 50)
		arabic, romanNumeral = romanNumeralCount(arabic, 10)
		arabic, romanNumeral = romanNumeralCount(arabic, 5)
		arabic, romanNumeral = romanNumeralCount(arabic, 1)

		/*
		if arabic >= 1000 {
			count := arabic / 1000
			for i := 0; i < count ; i++ {
				romanNumeral += "M"
				arabic -= 1000
			}
		} else if arabic >= 500 {
			count := arabic / 500
			for i := 0; i < count ; i++ {
				romanNumeral += "D"
				arabic -= 500
			}
		}
		 */
	}

	return romanNumeral, false
}

func romanNumeralCount(arabic int, value int) (int, string) {
	var romanNumeral = ""
	var letter = romanNumeralsMap[arabic]

	if arabic >= value {
		count := arabic / value
		for i := 0; i < count ; i++ {
			romanNumeral += letter
			arabic -= value
		}
	}

	return arabic, romanNumeral
}
