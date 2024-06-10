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
			3:    hd.Div3ShiftUnsigned(x),
			5:    hd.Div5ShiftUnsigned(x),
			6:    hd.Div6ShiftUnsigned(x),
			7:    hd.Div7ShiftUnsigned(x),
			9:    hd.Div9ShiftUnsigned(x),
			10:   hd.Div10ShiftUnsigned(x),
			11:   hd.Div11ShiftUnsigned(x),
			12:   hd.Div12ShiftUnsigned(x),
			13:   hd.Div13ShiftUnsigned(x),
			100:  hd.Div100ShiftUnsigned(x),
			1000: hd.Div1000ShiftUnsigned(x),
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
			3:    hd.Div3ShiftSigned(x),
			5:    hd.Div5ShiftSigned(x),
			6:    hd.Div6ShiftSigned(x),
			7:    hd.Div7ShiftSigned(x),
			9:    hd.Div9ShiftSigned(x),
			10:   hd.Div10ShiftSigned(x),
			11:   hd.Div11ShiftSigned(x),
			12:   hd.Div12ShiftSigned(x),
			13:   hd.Div13ShiftSigned(x),
			100:  hd.Div100ShiftSigned(x),
			1000: hd.Div1000ShiftSigned(x),
		}
		for i, got := range tests {
			if exp := x / i; got != exp {
				t.Errorf("%d: (%d) = %d; want %d", i, x, got, exp)
			}
		}
	})
}
