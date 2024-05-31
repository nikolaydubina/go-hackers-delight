package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleNLZ() {
	fmt.Println(hd.NLZ(255))
	// Output: 24
}

func ExampleNLZ_long() {
	fmt.Println(hd.NLZ(0b00111111111111111111111111101010))
	// Output: 2
}

func FuzzNLZCompute(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		if x == 0 {
			t.Skip()
		}

		// definition
		var n uint32 = 0
		for i := range 32 {
			if (x & (1 << (31 - i))) != 0 {
				n = uint32(i)
				break
			}
		}

		vs := []uint{
			hd.NLZ(x),
			hd.NLZ2(x),
		}
		for i, got := range vs {
			if n != uint32(got) {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", n, "got", got)
			}
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
			{exp: hd.NLZ(x) == hd.NLZ(y), got: hd.NLZEq(x, y)},
			{exp: hd.NLZ(x) < hd.NLZ(y), got: hd.NLZLess(x, y)},
			{exp: hd.NLZ(x) <= hd.NLZ(y), got: hd.NLZLessEq(x, y)},
		}
		for i, tc := range vs {
			if tc.exp != tc.got {
				t.Error(i, tc)
			}
		}
	})
}

func ExampleBitSize_zero() {
	fmt.Println(hd.BitSize(0))
	// Output: 0
}

func ExampleBitSize_bit6() {
	fmt.Println(hd.BitSize(0b0000_1101))
	// Output: 5
}

func ExampleBitSize_bit7() {
	fmt.Println(hd.BitSize(0b0011_1101))
	// Output: 7
}

func ExampleBitSize_bit32() {
	fmt.Println(hd.BitSize(0xFFFF_FFFF >> 1))
	// Output: 32
}

func ExampleNTZ() {
	fmt.Println(hd.NTZ(0b0000_1101_0000))
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

		// definition
		var n uint32 = 0
		for i := range 32 {
			if (x & (1 << i)) != 0 {
				n = uint32(i)
				break
			}
		}

		vs := []int{
			hd.NTZ(x),
		}
		for i, got := range vs {
			if n != uint32(got) {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", n, "got", got)
			}
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
