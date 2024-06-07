package hd_test

import (
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzDivShiftUnsigned(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		tests := map[uint32]uint32{
			3:    hd.DivShiftUnsigned3(x),
			5:    hd.DivShiftUnsigned5(x),
			6:    hd.DivShiftUnsigned6(x),
			7:    hd.DivShiftUnsigned7(x),
			9:    hd.DivShiftUnsigned9(x),
			10:   hd.DivShiftUnsigned10(x),
			11:   hd.DivShiftUnsigned11(x),
			12:   hd.DivShiftUnsigned12(x),
			13:   hd.DivShiftUnsigned13(x),
			100:  hd.DivShiftUnsigned100(x),
			1000: hd.DivShiftUnsigned1000(x),
		}
		for i, got := range tests {
			if exp := x / i; got != exp {
				t.Errorf("%d: (%d) = %d; want %d", i, x, got, exp)
			}
		}
	})
}

func FuzzDivShiftSigned(f *testing.F) {
	for _, u := range fuzzInt32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		tests := map[int32]int32{
			3:    hd.DivShiftSigned3(x),
			5:    hd.DivShiftSigned5(x),
			6:    hd.DivShiftSigned6(x),
			7:    hd.DivShiftSigned7(x),
			9:    hd.DivShiftSigned9(x),
			10:   hd.DivShiftSigned10(x),
			11:   hd.DivShiftSigned11(x),
			12:   hd.DivShiftSigned12(x),
			13:   hd.DivShiftSigned13(x),
			100:  hd.DivShiftSigned100(x),
			1000: hd.DivShiftSigned1000(x),
		}
		for i, got := range tests {
			if exp := x / i; got != exp {
				t.Errorf("%d: (%d) = %d; want %d", i, x, got, exp)
			}
		}
	})
}
