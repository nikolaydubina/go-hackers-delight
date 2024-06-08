package hd_test

import (
	"fmt"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleNextHigherNumberWithSameNumberOfOnes() {
	fmt.Printf("%08b", hd.NextHigherNumberWithSameNumberOfOnes(uint64(0b0011110000)))
	// Output: 100000111
}
