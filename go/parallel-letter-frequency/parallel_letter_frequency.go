package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency uses goroutines to make concurrent calls to the Frequency function
// for each of the strings in the []string parameter and returns one combined FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)
	combinedFreq := FreqMap{}

	for _, s := range texts {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for range texts {
		for k, v := range <-c {
			combinedFreq[k] += v
		}
	}
	return combinedFreq
}
