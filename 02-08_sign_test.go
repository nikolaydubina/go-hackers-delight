package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleSign() {
	fmt.Println(hd.Sign(-10), hd.Sign(10), hd.Sign(0))
	// Output: -1 1 0
}

func ExampleIsMostSignificantSet_int32() {
	fmt.Println(hd.IsMostSignificantSet(int32(-1)), hd.IsMostSignificantSet(int32(1)), hd.IsMostSignificantSet(int32(math.MaxInt32)))
	// Output: true false false
}

func ExampleIsMostSignificantSet_uint32() {
	fmt.Println(hd.IsMostSignificantSet(uint32(0xFFFFFFFF)), hd.IsMostSignificantSet(uint32(10)))
	// Output: true false
}

func FuzzSign(f *testing.F) {
	var vs = []int32{
		0,
		1,
		math.MaxInt32,
		math.MinInt32,
	}
	for _, x := range vs {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		q := hd.Sign(x)

		var v int32
		switch {
		case x > 0:
			v = 1
		case x < 0:
			v = -1
		default:
			v = 0
		}

		if q != v {
			t.Error("x", x, "got", q, "exp", v)
		}
	})
}
