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
		if gotQ := hd.DivExactSeven(x); q != gotQ {
			t.Errorf("DivExactSeven(%d) = %d; want %d", x, gotQ, q)
		}
	})
}

func ExampleMultiplicativeInverseEuclidInt() {
	fmt.Printf("%X", hd.MultiplicativeInverseEuclidInt(7))
	// Output: -49249249
}

func uint32FromHex(x int32) uint32 { return uint32(x) }

func TestMultiplicativeInverse(t *testing.T) {
	tests := map[uint32]uint32{
		uint32FromHex(-7): 0x4924_9249,
		uint32FromHex(-5): 0x3333_3333,
		uint32FromHex(-3): 0x5555_5555,
		uint32FromHex(-1): 0xFFFF_FFFF,
		1:                 0x0000_0001,
		3:                 0xAAAA_AAAB,
		5:                 0xCCCC_CCCD,
		7:                 0xB6DB_6DB7,
		9:                 0x38E3_8E39,
		11:                0xBA2E_8BA3,
		13:                0xC4EC_4EC5,
		15:                0xEEEE_EEEF,
		25:                0xC28F_5C29,
		125:               0x26E9_78D5,
		625:               0x3AFB_7E91,
	}
	for d, i := range tests {
		vs := []uint32{
			hd.MultiplicativeInverseEuclid(d),
			hd.MultiplicativeInverseNewton(d),
		}
		for j, v := range vs {
			if i != v {
				t.Errorf("%d: d(%d) = %d; want %d", j, d, v, i)
			}
		}
	}
}

func FuzzDivExact(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, d := range fuzzInt32 {
			f.Add(x, d)
		}
	}
	f.Fuzz(func(t *testing.T, x, d int32) {
		if (d % 2) == 0 {
			d++
		}
		q, r := x/d, x%d
		if r != 0 {
			t.Skip()
		}
		if gotQ := hd.DivExact(x, d); q != gotQ {
			t.Errorf("DivExact(%d) = %d; want %d", x, gotQ, q)
		}
	})
}
