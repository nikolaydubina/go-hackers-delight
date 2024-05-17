package hd_test

import (
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzInequalitiesAmongLogicAndArithmetic(f *testing.F) {
	var vs = []uint32{
		0,
		1,
		math.MaxInt32,
		math.MaxInt32 / 2,
		math.MaxInt32 - 1,
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint32) {
		v := []bool{
			(x ^ y) <= (x | y),
			(x | y) >= max(x, y),
			(x & y) <= min(x, y),
			hd.AbsDiff(x, y) <= (x ^ y),
		}
		for i, q := range v {
			if !q {
				t.Error(i, x, y)
			}
		}

		if isAddOverflow := hd.IsMostSignificantSet(hd.IsAddOverflowUnsigned(x, y)); isAddOverflow {
			if (x | y) <= (x + y) {
				t.Error(x, y)
			}
		} else {
			if (x | y) > (x + y) {
				t.Error(x, y)
			}
		}
	})
}
