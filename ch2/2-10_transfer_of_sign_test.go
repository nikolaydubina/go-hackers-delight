package ch2_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleISIGN() {
	fmt.Println(ch2.ISIGN(10, -100000))
	// Output: -10
}

func ExampleISIGN_positive() {
	fmt.Println(ch2.ISIGN(-10, 100000))
	// Output: 10
}

func ExampleISIGNTwo() {
	fmt.Println(ch2.ISIGNTwo(10, -100000))
	// Output: -10
}

func ExampleISIGNThree() {
	fmt.Println(ch2.ISIGNThree(10, -100000))
	// Output: -10
}

func ExampleISIGNFour() {
	fmt.Println(ch2.ISIGNFour(10, -100000))
	// Output: -10
}
