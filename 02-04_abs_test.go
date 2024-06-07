package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleAbs() {
	fmt.Print(hd.Abs(-42))
	// Output: 42
}

func ExampleNAbs() {
	fmt.Print(hd.NAbs(-42))
	// Output: -42
}

func ExampleAbsDiff() {
	fmt.Print(hd.AbsDiff(1, 100))
	// Output: 99
}

func FuzzAbsNormal(f *testing.F) {
	for _, x := range fuzzInt32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		abs := x
		if abs < 0 {
			abs = -abs
		}

		t.Run("abs", func(t *testing.T) {
			vs := []int32{
				hd.Abs(x),
				hd.Abs2(x),
				hd.Abs3(x),
				hd.Abs4(x),
				hd.AbsFastMul(x),
			}
			for i, v := range vs {
				if v != abs {
					t.Error(i, "x", x, "exp", abs, "got", v)
				}
			}
		})

		t.Run("nabs", func(t *testing.T) {
			vs := []int32{
				hd.NAbs(x),
				hd.NAbs2(x),
				hd.NAbs3(x),
			}
			for i, v := range vs {
				if v != -abs {
					t.Error(i, "x", x, "exp", -abs, "got", v)
				}
			}
		})
	})
}

func FuzzAbsDiffInt(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, y := range fuzzInt32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		vs := []int32{
			hd.AbsDiff2(x, y),
		}
		for i, v := range vs {
			if v != hd.AbsDiff(x, y) {
				t.Error(i, x, y, v)
			}
		}
	})
}

func FuzzAbsDiffUint(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint32) {
		ux, uy := uint64(x), uint64(y)
		exp := uint32(max(ux, uy) - min(ux, uy))

		if hd.AbsDiffUnsigned(x, y) != exp {
			t.Error(x, y, hd.AbsDiffUnsigned(x, y), exp)
		}
	})
}
