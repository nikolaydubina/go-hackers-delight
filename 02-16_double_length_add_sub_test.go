package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleAddDoubleLength() {
	x := [2]uint32{310, 410}
	y := [2]uint32{100, 200}
	fmt.Println(hd.AddDoubleLength(x, y))
	// Output: [410 610]
}

func ExampleSubDoubleLength() {
	x := [2]uint32{310, 405}
	y := [2]uint32{100, 200}
	fmt.Println(hd.SubDoubleLength(x, y))
	// Output: [210 205]
}

func FuzzDoubleLength(f *testing.F) {
	var vs = []uint64{
		0,
		1,
		math.MaxUint32,
		math.MaxUint32 / 2,
		math.MaxUint32 - 1,
		math.MaxUint32 + 1,
		math.MaxUint64,
		math.MaxUint64 / 2,
		math.MaxUint64 - 1,
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint64) {
		dx, dy := hd.DoubleLengthInt32FromUint64(x), hd.DoubleLengthInt32FromUint64(y)

		v := []struct {
			exp uint64
			got [2]uint32
		}{
			{x + y, hd.AddDoubleLength(dx, dy)},
			{x - y, hd.SubDoubleLength(dx, dy)},
		}
		for i, q := range v {
			if hd.DoubleLengthInt32FromUint64(q.exp) != q.got {
				t.Error(i, q, x, y)
			}
		}
	})
}
