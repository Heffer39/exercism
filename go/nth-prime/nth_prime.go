// Package prime calculates prime numbers
package prime

import (
	"math"
)

// Nth returns the nth prime number, given a number n
func Nth(index int) (int, bool) {
	var val, count int
	for i := 2; count < index; i++ {
		if isPrime(i) {
			count++
			val = i
		}
	}
	return val, count > 0
}

// isPrime determines if the given number is a prime
func isPrime(j int) bool {
	sq := int(math.Sqrt(float64(j)))
	for i := 2; i <= sq; i++ {
		if j%i == 0 {
			return false
		}
	}
	return true
}
