package pythagorean

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var allTriplets []Triplet

	for i := min; i <= max; i++ {
		allTriplets = append(allTriplets, Sum(i)...)
	}

	return allTriplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet{

}