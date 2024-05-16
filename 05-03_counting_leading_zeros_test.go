package hd_test

import (
	"fmt"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleNLZ() {
	fmt.Println(hd.NLZ(255))
	// Output: 24
}

func ExampleNLZ_long() {
	fmt.Println(hd.NLZ(0b00111111111111111111111111101010))
	// Output: 2
}
