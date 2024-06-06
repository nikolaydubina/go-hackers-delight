package hd_test

import (
	"fmt"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleExtendSign7_notExtended() {
	fmt.Printf("%08b", hd.ExtendSign7(0b0010_1010))
	// Output: 00101010
}

func ExampleExtendSign7_extended() {
	fmt.Printf("%08b", hd.ExtendSign7(0b1110_1010))
	// Output: 11111111111111111111111111101010
}

func ExampleExtendSign7Two_notExtended() {
	fmt.Printf("%08b", hd.ExtendSign7Two(0b0010_1010))
	// Output: 00101010
}

func ExampleExtendSign7Two_extended() {
	fmt.Printf("%08b", hd.ExtendSign7Two(0b1110_1010))
	// Output: 11111111111111111111111111101010
}

func ExampleExtendSign7Three_notExtended() {
	fmt.Printf("%08b", hd.ExtendSign7Three(0b0010_1010))
	// Output: 00101010
}

func ExampleExtendSign7Three_extended() {
	fmt.Printf("%08b", hd.ExtendSign7Three(0b1110_1010))
	// Output: 11111111111111111111111111101010
}
