package hd_test

import (
	"fmt"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleShiftLeftDoubleLength() {
	x := [2]uint32{0b1111, 0b1111}
	y := hd.ShiftLeftDoubleLength(x, 3)
	fmt.Printf("%032b_%032b", y[1], y[0])
	// Output: 00000000000000000000000001111000_00000000000000000000000001111000
}

func ExampleShiftRightUnsignedDoubleLength() {
	x := [2]uint32{0b1111, 0b1111}
	y := hd.ShiftRightUnsignedDoubleLength(x, 3)
	fmt.Printf("%032b_%032b", y[1], y[0])
	// Output: 00000000000000000000000000000001_11100000000000000000000000000001
}

func ExampleShiftRightSignedDoubleLength() {
	x := [2]uint32{0b1111, 0b1111}
	y := hd.ShiftRightSignedDoubleLength(x, 3)
	fmt.Printf("%032b_%032b", y[1], y[0])
	// Output: 00000000000000000000000000000001_11100000000000000000000000000001
}

func ExampleShiftRightSignedDoubleLength_negative() {
	x := [2]uint32{0b1111, 0b10000000000000000000000000000001}
	y := hd.ShiftRightSignedDoubleLength(x, 3)
	fmt.Printf("%032b_%032b", y[1], y[0])
	// Output: 11110000000000000000000000000000_00100000000000000000000000000001
}
