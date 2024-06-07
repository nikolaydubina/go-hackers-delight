package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleReverseBits() {
	fmt.Printf("0x%X", hd.ReverseBits(0x01234567))
	// Output: 0xE6A2C480
}

func FuzzReverseBitsEquality(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		exp := hd.ReverseBits(x)

		vs := []uint32{
			hd.ReverseBits2(x),
		}
		for i, got := range vs {
			if exp != got {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", exp, "got", got)
			}
		}
	})
}

func ExampleReverseBits64Knuth() {
	fmt.Printf("0x%X", hd.ReverseBits64Knuth(0x01234567_89ABCDEF))
	// Output: 0xF7B3D1D1E6A2C480
}

func FuzzReverseBits64(f *testing.F) {
	for _, x := range fuzzUint64 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint64) {
		var exp uint64
		for i := range 64 {
			exp = (exp << 1) | ((x >> i) & 1)
		}

		if got := hd.ReverseBits64Knuth(x); exp != got {
			t.Errorf("x=%0X exp=%0X got=%0X", x, exp, got)
		}
	})
}

func ExampleReverse8Bit() {
	fmt.Printf("%08b", hd.Reverse8Bit(0b1101_0101))
	// Output: 10101011
}

func FuzzIncrReversed(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		for i, ri := x, hd.ReverseBits(x); i < (x+100) && (i+1) < math.MaxUint32; i, ri = i+1, hd.IncrReversed(ri) {
			if exp := hd.ReverseBits(i); exp != ri {
				t.Errorf("i=%0X exp=%0X got=%0X", i, exp, ri)
			}
		}
	})
}
