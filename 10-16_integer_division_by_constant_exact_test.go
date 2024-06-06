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
	tests := []struct {
		d uint32
		i uint32
	}{
		{d: uint32FromHex(-7), i: 0x4924_9249},
		{d: uint32FromHex(-5), i: 0x3333_3333},
		{d: uint32FromHex(-3), i: 0x5555_5555},
		{d: uint32FromHex(-1), i: 0xFFFF_FFFF},
		{d: 1, i: 1},
		{d: 3, i: 0xAAAA_AAAB},
		{d: 5, i: 0xCCCC_CCCD},
		{d: 7, i: 0xB6DB_6DB7},
		{d: 9, i: 0x38E3_8E39},
		{d: 11, i: 0xBA2E_8BA3},
		{d: 13, i: 0xC4EC_4EC5},
		{d: 15, i: 0xEEEE_EEEF},
		{d: 25, i: 0xC28F_5C29},
		{d: 125, i: 0x26E9_78D5},
		{d: 625, i: 0x3AFB_7E91},
	}
	for _, tt := range tests {
		vs := []uint32{
			hd.MultiplicativeInverseEuclid(tt.d),
			hd.MultiplicativeInverseNewton(tt.d),
		}
		for i, v := range vs {
			if tt.i != v {
				t.Errorf("d(%d) = %d; want %d; method %d", tt.d, v, tt.i, i)
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
		if d == 0 {
			t.Skip()
		}
		if d%2 == 0 {
			t.Skip()
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
