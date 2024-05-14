package ch2_test

import (
	"fmt"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleSign() {
	fmt.Println(ch2.Sign(-10), ch2.Sign(10), ch2.Sign(0))
	// Output: -1 1 0
}

func ExampleSign_high() {
	fmt.Println(ch2.Sign(-2147483647), ch2.Sign(2147483647))
	// Output: -1 1
}
