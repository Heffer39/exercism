// Package summultiples sums up multiples up to a limit
package summultiples

// SumMultiples finds the sum of all the unique multiples of particular
// numbers up to but not including that number
func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d > 0 && i%d == 0 {
				sum += i
				break
			}
		}
	}
	return
}
