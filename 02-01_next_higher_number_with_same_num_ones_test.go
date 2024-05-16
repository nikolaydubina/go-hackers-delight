package ch2_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleNextHigherNumberWithSameNumberOfOnes() {
	fmt.Printf("%08b", ch2.NextHigherNumberWithSameNumberOfOnes(0b0011110000))
	// Output: 100000111
}
