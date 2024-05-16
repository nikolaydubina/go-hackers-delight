package ch2_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleTurnOffRightMostBit() {
	fmt.Printf("%08b", ch2.TurnOffRightMostBit(0b01011000))
	// Output: 01010000
}

func ExampleTurnOnRightMostBit() {
	fmt.Printf("%08b", ch2.TurnOnRightMostBit(0b10100111))
	// Output: 10101111
}

func ExampleTurnOffTrailingOnes() {
	fmt.Printf("%08b", ch2.TurnOffTrailingOnes(0b10100111))
	// Output: 10100000
}

func ExampleTurnOnTrailingZeros() {
	fmt.Printf("%08b", ch2.TurnOnTrailingZeros(0b10101000))
	// Output: 10101111
}

func ExampleSetBitLastZero() {
	fmt.Printf("%08b", ch2.SetBitLastZero(0b10100111))
	// Output: 00001000
}

/*
func ExampleSetZeroBitLastOne() {
	fmt.Printf("%08b", ch2.SetZeroBitLastOne(0b10101000))
	// Output: 11110111
}
*/

func ExampleSetTrailingZeroes() {
	fmt.Printf("%08b", ch2.SetTrailingZeroes(0b01011000))
	// Output: 00000111
}

func ExampleSetTrailingZeroes2() {
	fmt.Printf("%08b", ch2.SetTrailingZeroes2(0b01011000))
	// Output: 00000111
}

func ExampleSetTrailingZeroes3() {
	fmt.Printf("%08b", ch2.SetTrailingZeroes3(0b01011000))
	// Output: 00000111
}

/*
func ExampleSetTrailingOnes() {
	fmt.Printf("%08b", ch2.SetTrailingOnes(0b10100111))
	// Output: 11111000
}
*/

func ExampleIsolateRightmostOneBit() {
	fmt.Printf("%08b", ch2.IsolateRightmostOneBit(0b01011000))
	// Output: 00001000
}

func ExampleSetTrailingZeroesWithRightMostOne() {
	fmt.Printf("%08b", ch2.SetTrailingZeroesWithRightMostOne(0b01011000))
	// Output: 00001111
}

func ExampleSetTrailingOnesWithRightMostOne() {
	fmt.Printf("%08b", ch2.SetTrailingOnesWithRightMostOne(0b01010111))
	// Output: 00001111
}

func ExampleTurnOffRightmostOnes() {
	fmt.Printf("%08b", ch2.TurnOffRightmostOnes(0b01011100))
	// Output: 01000000
}

func ExampleTurnOffRightmostOnes2() {
	fmt.Printf("%08b", ch2.TurnOffRightmostOnes2(0b01011100))
	// Output: 01000000
}
