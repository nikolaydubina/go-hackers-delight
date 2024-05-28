package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleParity_odd() {
	fmt.Println(hd.Parity(0b0000_1101))
	// Output: 1
}

func ExampleParity_even() {
	fmt.Println(hd.Parity(0b0010_1101))
	// Output: 0
}

func FuzzParity(f *testing.F) {
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
		n = n % 2

		vs := []int{
			hd.Parity(x),
			hd.Parity2(x),
		}
		if x < ((1 << 7) - 1) {
			vs = append(vs, hd.Parity3(x))
			vs = append(vs, hd.Parity4(x))
		}
		for i, got := range vs {
			if n != uint32(got) {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", n, "got", got)
			}
		}
	})
}
