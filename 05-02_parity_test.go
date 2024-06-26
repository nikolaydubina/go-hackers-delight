package hd_test

import (
	"fmt"
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

func parity(x uint32) uint32 {
	var n uint32
	for i := range 32 {
		if (x & (1 << i)) != 0 {
			n++
		}
	}
	return n % 2
}

func FuzzParity(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		n := parity(x)

		vs := []uint32{
			hd.Parity(x),
			hd.Parity2(x),
		}
		if x < ((1 << 7) - 1) {
			vs = append(vs, hd.Parity3(x))
			vs = append(vs, hd.Parity4(x))
		}
		for i, got := range vs {
			if n != got {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", n, "got", got)
			}
		}
	})
}
