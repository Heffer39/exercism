// Package strand contains functionality revolving around dna and rna strings
package strand

var dnaToRnaMap = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA converts a DNA string to a RNA string
func ToRNA(dna string) string {
	rna := ""
	for _, char := range dna {
		rna += dnaToRnaMap[string(char)]
	}
	return rna
}

/* The benchmark results showed that this was a faster solution and used less memory,
but at the cost of being more difficult to maintain

func ToRNA(dna string) string {
	rna := ""
	for _, char := range dna {
		if char == 'G' {
			rna += "C"
		} else if char == 'C' {
			rna += "G"
		} else if char == 'T' {
			rna += "A"
		} else if char == 'A' {
			rna += "U"
		}
	}
	return rna
}
*/
