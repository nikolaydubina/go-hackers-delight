package hd_test

import (
	"fmt"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleShiftRightSignedFromUnsigned() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned2() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned2(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned3() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned3(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned4() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned4(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned5() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned5(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}
