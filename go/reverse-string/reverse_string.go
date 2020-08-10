// Package reverse is able to reverse the order of a string
package reverse

// Reverse will return a string in the reverse order of the passed-in string
func Reverse(input string) (reverse string) {
	for _, r := range input {
		reverse = string(r) + reverse
	}
	return
}