package hd_test

import (
	"fmt"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleTurnOffRightMostBit() {
	fmt.Printf("%08b", hd.TurnOffRightMostBit(0b01011000))
	// Output: 01010000
}

func ExampleTurnOnRightMostBit() {
	fmt.Printf("%08b", hd.TurnOnRightMostBit(0b10100111))
	// Output: 10101111
}

func ExampleTurnOffTrailingOnes() {
	fmt.Printf("%08b", hd.TurnOffTrailingOnes(0b10100111))
	// Output: 10100000
}

func ExampleTurnOnTrailingZeros() {
	fmt.Printf("%08b", hd.TurnOnTrailingZeros(0b10101000))
	// Output: 10101111
}

func ExampleSetBitLastZero() {
	fmt.Printf("%08b", hd.SetBitLastZero(0b10100111))
	// Output: 00001000
}

func ExampleSetTrailingZeros() {
	fmt.Printf("%08b", hd.SetTrailingZeros(0b01011000))
	// Output: 00000111
}

func ExampleSetTrailingZeros2() {
	fmt.Printf("%08b", hd.SetTrailingZeros2(0b01011000))
	// Output: 00000111
}

func ExampleSetTrailingZeros3() {
	fmt.Printf("%08b", hd.SetTrailingZeros3(0b01011000))
	// Output: 00000111
}

func ExampleIsolateRightmostOneBit() {
	fmt.Printf("%08b", hd.IsolateRightmostOneBit(0b01011000))
	// Output: 00001000
}

func ExampleSetTrailingZerosWithRightMostOne() {
	fmt.Printf("%08b", hd.SetTrailingZerosWithRightMostOne(0b01011000))
	// Output: 00001111
}

func ExampleSetTrailingOnesWithRightMostOne() {
	fmt.Printf("%08b", hd.SetTrailingOnesWithRightMostOne(0b01010111))
	// Output: 00001111
}

func ExampleTurnOffRightmostOnes() {
	fmt.Printf("%08b", hd.TurnOffRightmostOnes(0b01011100))
	// Output: 01000000
}

func ExampleTurnOffRightmostOnes2() {
	fmt.Printf("%08b", hd.TurnOffRightmostOnes2(0b01011100))
	// Output: 01000000
}
