package lsproduct

import (
	"fmt"
	"strconv"
)

// LargestSeriesProduct calculates the largest product for a contiguous substring of digits of length n
func LargestSeriesProduct(input string, span int) (int, error) {
	if span == 0 {
		return 1, nil
	}
	if span < 0 {
		return 0, fmt.Errorf("span must be greater than zero")
	}
	if span > len(input) {
		return 0, fmt.Errorf("span must be smaller than string length")
	}

	numbers := make([]int, len(input))
	for i, v := range input {
		var err error
		numbers[i], err = strconv.Atoi(string(v))
		if err != nil {
			return 0, fmt.Errorf("digits input must only contain digits")
		}
	}

	var largestSeriesProduct int

	for i := range numbers {
		if i+span > len(numbers) {
			break
		}
		seriesProduct := 0
		for j := i; j < span+i; j++ {
			if j == i {
				seriesProduct += numbers[j]
			} else {
				seriesProduct *= numbers[j]
			}
		}
		if seriesProduct > largestSeriesProduct {
			largestSeriesProduct = seriesProduct
		}
	}

	return largestSeriesProduct, nil
}
