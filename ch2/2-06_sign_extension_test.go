package ch2_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleExtendSign7_notExtended() {
	fmt.Printf("%08b", ch2.ExtendSign7(0b00101010))
	// Output: 00101010
}

func ExampleExtendSign7_extended() {
	fmt.Printf("%08b", ch2.ExtendSign7(0b11101010))
	// Output: 11111111111111111111111111101010
}

func ExampleExtendSign7Two_notExtended() {
	fmt.Printf("%08b", ch2.ExtendSign7Two(0b00101010))
	// Output: 00101010
}

func ExampleExtendSign7Two_extended() {
	fmt.Printf("%08b", ch2.ExtendSign7Two(0b11101010))
	// Output: 11111111111111111111111111101010
}

func ExampleExtendSign7Three_notExtended() {
	fmt.Printf("%08b", ch2.ExtendSign7Three(0b00101010))
	// Output: 00101010
}

func ExampleExtendSign7Three_extended() {
	fmt.Printf("%08b", ch2.ExtendSign7Three(0b11101010))
	// Output: 11111111111111111111111111101010
}
