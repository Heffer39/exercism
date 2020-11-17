package pythagorean

// Triplet is an array of three integers, representing a Pythagorean triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (allTriplets []Triplet) {
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					allTriplets = append(allTriplets, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(sum int) (triplets []Triplet) {
	for a := 1; a <= sum/3; a++ {
		for b := a + 1; b <= sum/2; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return
}
