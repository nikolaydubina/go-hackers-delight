package ch2_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleAbs() {
	fmt.Print(ch2.Abs(-42))
	// Output: 42
}

func ExampleAbs2() {
	fmt.Print(ch2.Abs2(-42))
	// Output: 42
}

func ExampleAbs3() {
	fmt.Print(ch2.Abs3(-42))
	// Output: 42
}

func ExampleNAbs() {
	fmt.Print(ch2.NAbs(-42))
	// Output: -42
}

func ExampleNAbs2() {
	fmt.Print(ch2.NAbs2(-42))
	// Output: -42
}

func ExampleNAbs3() {
	fmt.Print(ch2.NAbs3(-42))
	// Output: -42
}

func ExampleAbsFastMult() {
	fmt.Print(ch2.AbsFastMult(-42))
	// Output: 42
}
