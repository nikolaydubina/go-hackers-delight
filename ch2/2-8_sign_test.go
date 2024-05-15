package ch2_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleSign() {
	fmt.Println(ch2.Sign(-10), ch2.Sign(10), ch2.Sign(0))
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
		q := ch2.Sign(x)

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
