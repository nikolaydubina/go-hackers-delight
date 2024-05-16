package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleISIGN() {
	fmt.Println(hd.ISIGN(10, -100000), hd.ISIGN(-10, 100000))
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
			hd.ISIGN(x, y),
			hd.ISIGN2(x, y),
			hd.ISIGN3(x, y),
			hd.ISIGN4(x, y),
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
