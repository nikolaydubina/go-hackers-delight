package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleCountOnes() {
	fmt.Println(hd.CountOnes5(uint32(0b0000_0100_0100_01110)))
	// Output: 5
}

func FuzzCountOnes(f *testing.F) {
	var vs = []uint32{
		0,
		1,
		math.MaxInt32,
		math.MaxInt32 / 2,
		math.MaxInt32 - 1,
		math.MaxUint32,
		math.MaxUint32 / 2,
		math.MaxUint32 - 1,
	}
	for _, x := range vs {
		f.Add(x)
	}

	f.Fuzz(func(t *testing.T, x uint32) {
		// definition
		var n uint32 = 0
		for i := range 32 {
			if (x & (1 << i)) != 0 {
				n++
			}
		}

		vs := []uint32{
			hd.CountOnes(x),
			hd.CountOnes1(x),
			//hd.CountOnes2(x),
			hd.CountOnes3(x),
			hd.CountOnes4(x),
		}
		if x <= ((1 << 15) - 1) {
			vs = append(vs, hd.CountOnes5(x))
		}

		for i, v := range vs {
			if v != n {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", n, "got", v)
			}
		}
	})
}
