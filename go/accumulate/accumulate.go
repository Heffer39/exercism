// Package accumulate provides an accessible way to perform operations
// on a slice of strings
package accumulate

// Accumulate performs the passed in converter function on a slice of strings
// Accumulate returns a new []string and does not modify the original slice
func Accumulate(words []string, converter func(string) string) []string {

	// This is the in-place solution, so it is more memory-efficient and faster
	// however, it modifies the "given []string" in the
	// test suite, which negatively impacts logging and debugging
	/*
		for i, word := range words {
			words[i] = converter(word)
		}
		return words
	*/

	// This solution requires extra memory, but it preserves the correct
	// initial test data when collecting logs during the tests
	modifiedWords := make([]string, len(words))
	copy(modifiedWords, words)

	for i, word := range modifiedWords {
		modifiedWords[i] = converter(word)
	}
	return modifiedWords
}
