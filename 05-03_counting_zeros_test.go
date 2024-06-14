package hd_test

import (
	"fmt"
	"math/bits"
	"math/rand/v2"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleLeadingZerosUint32() {
	fmt.Println(hd.LeadingZerosUint32(255))
	// Output: 24
}

func ExampleLeadingZerosUint32_long() {
	fmt.Println(hd.LeadingZerosUint32(0b00111111111111111111111111101010))
	// Output: 2
}

func FuzzNLZCompute32(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		if x == 0 {
			t.Skip()
		}

		n := bits.LeadingZeros32(x)

		vs := []uint8{
			hd.LeadingZerosUint32(x),
			hd.LeadingZerosUint32BinarySearch(x),
		}
		for i, got := range vs {
			if uint8(n) != got {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", n, "got", got)
			}
		}
	})
}

func FuzzNLZCompute64(f *testing.F) {
	for _, x := range fuzzUint64 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint64) {
		if x == 0 {
			t.Skip()
		}
		exp := bits.LeadingZeros64(x)
		got := hd.LeadingZerosUint64(x)
		if uint8(exp) != got {
			t.Error("x", fmt.Sprintf("%064b", x), "exp", exp, "got", got)
		}
	})
}

func FuzzNLZCompare(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint32) {
		if x == 0 || y == 0 {
			t.Skip()
		}
		vs := []struct {
			exp bool
			got bool
		}{
			{exp: bits.LeadingZeros32(x) == bits.LeadingZeros32(y), got: hd.LeadingZerosEqual(x, y)},
			{exp: bits.LeadingZeros32(x) < bits.LeadingZeros32(y), got: hd.LeadingZerosLess(x, y)},
			{exp: bits.LeadingZeros32(x) <= bits.LeadingZeros32(y), got: hd.LeadingZerosLessOrEqual(x, y)},
		}
		for i, tc := range vs {
			if tc.exp != tc.got {
				t.Error(i, tc)
			}
		}
	})
}

func ExampleBitSize32_zero() {
	fmt.Println(hd.BitSize32(0))
	// Output: 0
}

func ExampleBitSize32_bit6() {
	fmt.Println(hd.BitSize32(0b0000_1101))
	// Output: 5
}

func ExampleBitSize32_bit7() {
	fmt.Println(hd.BitSize32(0b0011_1101))
	// Output: 7
}

func ExampleBitSize32_bit32() {
	fmt.Println(hd.BitSize32(0xFFFF_FFFF >> 1))
	// Output: 32
}

func ExampleTrailingZerosUint32() {
	fmt.Println(hd.TrailingZerosUint32(0b0000_1101_0000))
	// Output: 4
}

func FuzzNTZCompute(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		if x == 0 {
			t.Skip()
		}
		exp := bits.TrailingZeros32(x)
		got := hd.TrailingZerosUint32(x)
		if uint8(exp) != got {
			t.Error("x", fmt.Sprintf("%032b", x), "exp", exp, "got", got)
		}
	})
}

func ExampleLoopDetectionGosper() {
	f := func(v int) int {
		if v < 100 {
			return v + 1
		}
		return 1
	}
	fmt.Println(hd.LoopDetectionGosper(f, 0))
	// Output: 0 63 100
}

func leadingZeros32(x uint32) uint8 { return uint8(bits.LeadingZeros32(x)) }

func leadingZeros64(x uint64) uint8 { return uint8(bits.LeadingZeros64(x)) }

func BenchmarkLeadingZeros(b *testing.B) {
	b.Run("uint32", func(b *testing.B) {
		var out uint8

		var vals []uint32
		for i := 0; i < 1000; i++ {
			vals = append(vals, rand.Uint32())
		}

		vs := []struct {
			name string
			f    func(x uint32) uint8
		}{
			{"basic", leadingZeros32},
			{"LeadingZerosUint32", hd.LeadingZerosUint32},
		}
		for _, v := range vs {
			b.Run(v.name, func(b *testing.B) {
				for i := 0; i < b.N; i += len(vals) {
					for j := 0; j < len(vals); j++ {
						out = v.f(vals[j])
					}
				}
			})
		}

		if (out*2 - out - out) != 0 {
			b.Fatal("never")
		}
	})

	b.Run("uint64", func(b *testing.B) {
		var out uint8

		var vals []uint64
		for i := 0; i < 1000; i++ {
			vals = append(vals, rand.Uint64())
		}

		vs := []struct {
			name string
			f    func(x uint64) uint8
		}{
			{"basic", leadingZeros64},
			{"LeadingZerosUint32", hd.LeadingZerosUint64},
		}
		for _, v := range vs {
			b.Run(v.name, func(b *testing.B) {
				for i := 0; i < b.N; i += len(vals) {
					for j := 0; j < len(vals); j++ {
						out = v.f(vals[j])
					}
				}
			})
		}

		if (out*2 - out - out) != 0 {
			b.Fatal("never")
		}
	})
}
