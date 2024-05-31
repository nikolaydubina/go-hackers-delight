package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleTransposeMatrix8bx8b() {
	A := []byte{
		0b1010_1011,
		0b0101_1110,
		0b0001_0101,
		0b0000_1000,
		0b0001_0000,
		0b0011_0000,
		0b0101_0000,
		0b1001_0000,
	}
	B := make([]byte, 8)

	hd.TransposeMatrix8bx8b(A, B, 1, 1)

	for _, q := range B {
		fmt.Printf("%08b\n", q)
	}
	// Output:
	// 10000001
	// 01000010
	// 10000100
	// 01101111
	// 11010000
	// 01100000
	// 11000000
	// 10100000
}

func FuzzTransposeMatrix8bx8b(f *testing.F) {
	f.Fuzz(func(t *testing.T, v0, v1, v2, v3, v4, v5, v6, v7 byte) {
		A := []byte{v0, v1, v2, v3, v4, v5, v6, v7}
		B := make([]byte, 8)

		hd.TransposeMatrix8bx8b(A, B, 1, 1)

		for i := range 8 {
			for j := range 8 {
				a := (A[i] & (1 << (7 - j))) == 0
				b := (B[j] & (1 << (7 - i))) == 0

				if a != b {
					t.Error(i, j, a, b)
				}
			}
		}
	})
}
