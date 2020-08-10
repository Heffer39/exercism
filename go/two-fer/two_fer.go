// Package twofer implements "two-fer", or two for one messages
package twofer

// Returns a "Two-fer" string - "One for X, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
