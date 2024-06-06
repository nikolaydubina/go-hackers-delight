package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzDivModExactSeven(f *testing.F) {
	for _, u := range fuzzInt32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		q, r := x/7, x%7
		if r != 0 {
			t.Skip()
		}
		if gotQ, gotR := hd.DivModExactSeven(x); q != gotQ || r != gotR {
			t.Errorf("DivModExactSeven(%d) = got(%d, %d); want (%d, %d)", x, gotQ, gotR, q, r)
		}
	})
}

func ExampleMultiplicativeInverseEuclidInt() {
	fmt.Printf("%X", hd.MultiplicativeInverseEuclidInt(7))
	// Output: -49249249
}

func ExampleMultiplicativeInverseEuclid() {
	fmt.Printf("%X", hd.MultiplicativeInverseEuclid(7))
	// Output: B6DB6DB7
}
