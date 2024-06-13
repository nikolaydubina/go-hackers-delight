package hd_test

import (
	"fmt"
	"math"
	"math/bits"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleReverseByte() {
	fmt.Printf("%0b", hd.ReverseByte(0b1101_0101))
	// Output: 10101011
}

func reverseByte(x byte) byte {
	var out byte
	out |= x & 0b0000_0001 << 7
	out |= x & 0b0000_0010 << 5
	out |= x & 0b0000_0100 << 3
	out |= x & 0b0000_1000 << 1
	out |= x & 0b0001_0000 >> 1
	out |= x & 0b0010_0000 >> 3
	out |= x & 0b0100_0000 >> 5
	out |= x & 0b1000_0000 >> 7
	return out
}

func FuzzReverseByte(f *testing.F) {
	f.Fuzz(func(t *testing.T, x byte) {
		exp := reverseByte(x)
		got := hd.ReverseByte(x)
		if exp != got {
			t.Errorf("x=%0b exp=%0b got=%0b", x, exp, got)
		}
	})
}

func ExampleReverseBits() {
	fmt.Printf("0x%X", hd.ReverseBits(0x01234567))
	// Output: 0xE6A2C480
}

func FuzzReverseBitsEquality(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		exp := bits.Reverse32(x)
		got := []uint32{
			hd.ReverseBits(x),
			hd.ReverseBits2(x),
		}
		for i, got := range got {
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

func FuzzReverseBits_uint64(f *testing.F) {
	for _, x := range fuzzUint64 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint64) {
		exp := bits.Reverse64(x)
		got := hd.ReverseBits64Knuth(x)
		if exp != got {
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
