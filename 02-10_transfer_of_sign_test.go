package ch2_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleISIGN() {
	fmt.Println(ch2.ISIGN(10, -100000), ch2.ISIGN(-10, 100000))
	// Output: -10 10
}

func FuzzISIGN(f *testing.F) {
	var vs = []int32{
		0,
		1,
		math.MaxInt32,
		math.MinInt32,
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		vs := []int32{
			ch2.ISIGN(x, y),
			ch2.ISIGN2(x, y),
			ch2.ISIGN3(x, y),
			ch2.ISIGN4(x, y),
		}
		for i, q := range vs {
			v := x
			if v < 0 {
				v = -v
			}
			if y < 0 {
				v = -v
			}

			if q != v {
				t.Error(i, x, y, "got", q, "exp", v)
			}
		}
	})
}
