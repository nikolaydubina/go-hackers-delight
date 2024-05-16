package hd_test

import (
	"fmt"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleRotateLeft() {
	fmt.Printf("%032b", hd.RotateLeft(0b11101111111100001111111111111101, 8))
	// Output: 11110000111111111111110111101111
}

func ExampleRotateRight() {
	fmt.Printf("%032b", hd.RotateRight(0b11101111111100001111111111111101, 8))
	// Output: 11111101111011111111000011111111
}
