package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleNLZ() {
	fmt.Println(hd.NLZ(255))
	// Output: 24
}

func ExampleNLZ_long() {
	fmt.Println(hd.NLZ(0b00111111111111111111111111101010))
	// Output: 2
}

func FuzzNLZ(f *testing.F) {
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
		if x == 0 {
			t.Skip()
		}

		// definition
		var n uint32 = 0
		for i := range 32 {
			if (x & (1 << (31 - i))) != 0 {
				n = uint32(i)
				break
			}
		}

		vs := []uint32{
			hd.NLZ(x),
		}
		for i, got := range vs {
			if n != uint32(got) {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", n, "got", got)
			}
		}
	})
}
