// Package strain contains "keep" and "discard" operations on collections
package strain

// Ints represents a slice of integers
type Ints []int

// Lists represents two dimensional slice of integers (slice of slices)
type Lists [][]int

// Strings represents a slice of strings
type Strings []string

// Keep returns a subset of integers from the Ints pointer receiver where
// the function parameter predicate is true
func (integers Ints) Keep(x func(int) bool) Ints {
	var cp Ints
	for _, v := range integers {
		if x(v) {
			cp = append(cp, v)
		}
	}
	return cp
}

// Discard returns a subset of integers from the Ints pointer receiver where
// the function parameter predicate is false
func (integers Ints) Discard(x func(int) bool) Ints {
	var cp Ints
	for _, v := range integers {
		if !x(v) {
			cp = append(cp, v)
		}
	}
	return cp
}

// Keep returns a subset of integers from the Lists pointer receiver where
// the function parameter predicate is true
func (lists Lists) Keep(x func([]int) bool) Lists {
	var cp Lists
	for _, v := range lists {
		if x(v) {
			cp = append(cp, v)
		}
	}
	return cp
}

// Keep returns a subset of integers from the Strings pointer receiver where
// the function parameter predicate is true
func (strings Strings) Keep(x func(string) bool) Strings {
	var cp Strings
	for _, v := range strings {
		if x(v) {
			cp = append(cp, v)
		}
	}
	return cp
}
