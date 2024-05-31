package hd_test

import (
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzIsInRangeNormal(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, a := range fuzzInt32 {
			for _, b := range fuzzInt32 {
				f.Add(x, a, b)
			}
		}
	}
	f.Fuzz(func(t *testing.T, x, a, b int32) {
		if a >= b {
			t.Skip()
		}
		vs := []struct {
			exp bool
			got bool
		}{
			{exp: a <= x && x <= b, got: hd.IsInRange(x, a, b)},
			{exp: a <= x && x < b, got: hd.IsInRangeClosedOpen(x, a, b)},
			{exp: a < x && x <= b, got: hd.IsInRangeOpenClosed(x, a, b)},
			{exp: a < x && x < b, got: hd.IsInRangeOpen(x, a, b)},
			{exp: a < x && x < b, got: hd.IsInRangeOpen2(x, a, b)},
		}
		for _, v := range vs {
			if v.exp != v.got {
				t.Errorf("x=%d, a=%d, b=%d: expected %v, got %v", x, a, b, v.exp, v.got)
			}
		}
	})
}

func FuzzIsInRangePowerTwo(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, a := range fuzzUint32 {
			for ip := range hd.PowerOfTwo[:31] {
				f.Add(x, a, uint8(ip))
			}
		}
	}
	f.Fuzz(func(t *testing.T, x, a uint32, ip uint8) {
		if int(ip) >= 31 {
			t.Skip()
		}
		p := hd.PowerOfTwo[ip]

		t.Run("IsInRangePowerTwo", func(t *testing.T) {
			exp := x <= uint32(p-1)
			got := hd.IsInRangePowerTwo(x, int(ip))
			if exp != got {
				t.Errorf("IsInRangePowerTwo: %v %v %v %0X %0X %v %v", x, a, ip, x-a, uint(x-a)>>ip, got, exp)
			}
		})

		t.Run("IsInRangePowerTwoOffset", func(t *testing.T) {
			if hd.IsMostSignificantSet(hd.IsSubOverflowUnsigned(x, a)) {
				t.Skip("offset formula is not resistant to overflows")
			}

			// naive approach to rely on uint64 space to protect from overflow
			exp := uint64(a) <= uint64(x) && uint64(x) <= uint64(a)+uint64(p-1)
			got := hd.IsInRangePowerTwoOffset(x, a, int(ip))

			if exp != got {
				t.Errorf("%v %v %v %0X %0X %v %v", x, a, ip, x-a, uint(x-a)>>ip, got, exp)
			}
		})
	})
}
