package ch2_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleAvgUnsigned() {
	fmt.Print(ch2.AvgUnsigned(100, 200))
	// Output: 150
}

func ExampleAvgUnsignedCeil() {
	fmt.Print(ch2.AvgUnsignedCeil(100, 200))
	// Output: 150
}

func ExampleAvg() {
	fmt.Print(ch2.Avg(-100, -200))
	// Output: -150
}

func ExampleAvgCeil() {
	fmt.Print(ch2.AvgCeil(-100, -200))
	// Output: -150
}

func ExampleAvgUnsigned_noOverflow() {
	fmt.Print(ch2.AvgUnsigned(2147483647, 2147483645))
	// Output: 2147483646
}

func ExampleAvgUnsignedCeil_noOverflow() {
	fmt.Print(ch2.AvgUnsignedCeil(2147483647, 2147483646))
	// Output: 2147483647
}

func ExampleAvg_noOverflow() {
	fmt.Print(ch2.Avg(-2147483647, -2147483645))
	// Output: -2147483646
}

func ExampleAvgCeil_noOverflow() {
	fmt.Print(ch2.AvgCeil(-2147483647, -2147483646))
	// Output: -2147483646
}
