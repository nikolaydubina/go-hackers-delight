package ch2_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleAbs() {
	fmt.Print(ch2.Abs(-42))
	// Output: 42
}

func ExampleNAbs() {
	fmt.Print(ch2.NAbs(-42))
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
				ch2.Abs(x),
				ch2.Abs2(x),
				ch2.Abs3(x),
				ch2.AbsFastMul(x),
			}
			for i, v := range vs {
				if v != abs {
					t.Error(i, "x", x, "exp", abs, "got", v)
				}
			}
		})

		t.Run("nabs", func(t *testing.T) {
			vs := []int32{
				ch2.NAbs(x),
				ch2.NAbs2(x),
				ch2.NAbs3(x),
			}
			for i, v := range vs {
				if v != -abs {
					t.Error(i, "x", x, "exp", -abs, "got", v)
				}
			}
		})
	})
}
