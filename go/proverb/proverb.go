// Package proverb will output the full text of a proverbial rhyme, given a list of inputs
package proverb

// Proverb returns a relevant proverb, based on the list of inputs
func Proverb(rhyme []string) (proverb []string) {
	if len(rhyme) > 0 {
		for i, v := range rhyme {
			if i+1 >= len(rhyme) {
				break
			}
			s := "For want of a " + v + " the " + rhyme[i+1] + " was lost."
			proverb = append(proverb, s)
		}
		finale := "And all for the want of a " + rhyme[0] + "."
		proverb = append(proverb, finale)
	}
	return proverb
}
