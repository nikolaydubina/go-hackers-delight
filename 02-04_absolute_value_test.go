package hd_test

import (
	"fmt"
	"math"
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

func FuzzAbs(f *testing.F) {
	var vs = []int32{
		0,
		1,
		-1,
		math.MinInt32,
		math.MaxInt32,
	}
	for _, x := range vs {
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
