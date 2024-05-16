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
