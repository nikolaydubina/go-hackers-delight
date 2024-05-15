package ch5_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch5"
)

func ExampleNLZ() {
	fmt.Println(ch5.NLZ(255))
	// Output: 24
}

func ExampleNLZ_long() {
	fmt.Println(ch5.NLZ(0b00111111111111111111111111101010))
	// Output: 2
}
