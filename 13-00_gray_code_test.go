package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleGrayCodeFromUint() {
	for i := range 10 {
		fmt.Printf("%08b -> %08b\n", i, hd.GrayCodeFromUint(uint64(i)))
	}
	// Output:
	// 00000000 -> 00000000
	// 00000001 -> 00000001
	// 00000010 -> 00000011
	// 00000011 -> 00000010
	// 00000100 -> 00000110
	// 00000101 -> 00000111
	// 00000110 -> 00000101
	// 00000111 -> 00000100
	// 00001000 -> 00001100
	// 00001001 -> 00001101
}

func FuzzGrayCode(f *testing.F) {
	for _, u := range fuzzUint64 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint64) {
		t.Run("uint16", func(t *testing.T) {
			x := uint16(x)
			if xx := hd.GrayCodeToUint16(hd.GrayCodeFromUint(x)); x != xx {
				t.Errorf("x=%d, xx=%d)", x, xx)
			}
		})

		t.Run("uint32", func(t *testing.T) {
			x := uint32(x)
			if xx := hd.GrayCodeToUint32(hd.GrayCodeFromUint(x)); x != xx {
				t.Errorf("x=%d, xx=%d)", x, xx)
			}
		})

		t.Run("uint64", func(t *testing.T) {
			if xx := hd.GrayCodeToUint64(hd.GrayCodeFromUint(x)); x != xx {
				t.Errorf("x=%d, xx=%d)", x, xx)
			}
		})
	})
}
