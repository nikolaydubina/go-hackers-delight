package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleDOZ() {
	fmt.Println(hd.DOZ(10, 5), hd.DOZ(5, 10), hd.DOZ(-5, -10), hd.DOZ(-10, -5))
	// Output: 5 0 5 0
}

func ExampleDOZU() {
	fmt.Println(hd.DOZU(10, 5), hd.DOZ(5, 10))
	// Output: 5 0
}

func FuzzDOZ(f *testing.F) {
	var vs = []int32{
		0,
		1,
		-1,
		math.MinInt32,
		math.MinInt32 / 2,
		math.MinInt32 + 1,
		math.MaxInt32,
		math.MaxInt32 / 2,
		math.MaxInt32 - 1,
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		var doz int32
		if x >= y {
			doz = x - y
		}

		v := []struct {
			exp int32
			got int32
		}{
			{doz, hd.DOZ(x, y)},
			{min(x, y), hd.Min(x, y)},
			{max(x, y), hd.Max(x, y)},
		}

		if x >= (-1<<30) && x <= (1<<30-1) && y >= (-1<<30) && y <= (1<<30-1) {
			v = append(v, []struct {
				exp int32
				got int32
			}{
				{doz, hd.DOZRanges(x, y)},
				{min(x, y), hd.MinRanges(x, y)},
				{max(x, y), hd.MaxRanges(x, y)},
			}...)
		}

		for i, q := range v {
			if q.exp != q.got {
				t.Error(i, q, x, y)
			}
		}
	})
}
