package raindrops

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	for _, test := range tests {
		if actual := Convert(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		} else {
			fmt.Printf("Convert(%d) = %q, expected %q.\n",
				test.input, actual, test.expected)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Convert(test.input)
		}
	}
}
