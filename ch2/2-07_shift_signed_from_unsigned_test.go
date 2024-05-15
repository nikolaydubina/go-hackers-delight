package ch2_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleShiftRightSignedFromUnsigned() {
	fmt.Printf("%08b", ch2.ShiftRightSignedFromUnsigned(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned2() {
	fmt.Printf("%08b", ch2.ShiftRightSignedFromUnsigned2(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned3() {
	fmt.Printf("%08b", ch2.ShiftRightSignedFromUnsigned3(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned4() {
	fmt.Printf("%08b", ch2.ShiftRightSignedFromUnsigned4(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned5() {
	fmt.Printf("%08b", ch2.ShiftRightSignedFromUnsigned5(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}
