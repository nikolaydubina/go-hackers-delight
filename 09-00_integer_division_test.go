package hd_test

import (
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzDivideMultiWord(f *testing.F) {
	for _, u := range fuzzUint64 {
		for _, v := range fuzzUint64 {
			f.Add(u, v)
		}
	}

	f.Add(uint64(123456), uint64(5))                               // single word case
	f.Add(uint64(0x7FFF_8000_0000_0000), uint64(0x8000_0000_0001)) // special case, trigger add-back logic that happens 0.003% of cases

	f.Fuzz(func(t *testing.T, u, v uint64) {
		if v == 0 {
			t.Skip()
		}
		if u < v {
			t.Skip()
		}

		expQ := u / v
		expR := u % v

		u16 := hd.IntToNx16b(u)
		v16 := hd.IntToNx16b(v)
		q16 := make([]uint16, len(u16))
		r16 := make([]uint16, len(u16))

		hd.DivMultiWordUnsigned(q16, r16, u16, v16)

		if got := hd.Uint64FromNx16b(q16); got != expQ {
			t.Errorf("u=%d %v v=%d %v: Q: exp=%d got=%d %v", u, u16, v, v16, expQ, got, q16)
		}

		if got := hd.Uint64FromNx16b(r16); got != expR {
			t.Errorf("u=%d %v v=%d %v: R: exp=%d got=%d %v", u, u16, v, v16, expR, got, r16)
		}
	})
}

func FuzzDivLongUnsigned64b(f *testing.F) {
	for _, u := range fuzzUint64 {
		for _, v := range fuzzUint32 {
			f.Add(u, v)
		}
	}

	f.Fuzz(func(t *testing.T, x uint64, y uint32) {
		if y == 0 {
			t.Skip()
		}

		if (x / uint64(y)) > math.MaxUint32 {
			t.Skip()
		}

		expQ := uint32(x / uint64(y))
		expR := uint32(x % uint64(y))

		t.Run("DivLongUnsigned64b32b", func(t *testing.T) {
			q, r := hd.DivLongUnsigned64b32b(x, y)

			if q != expQ {
				t.Errorf("x=%d y=%d: Q: exp=%d got=%d", x, y, expQ, q)
			}
			if r != expR {
				t.Errorf("x=%d y=%d: R: exp=%d got=%d", x, y, expR, r)
			}
		})

		t.Run("DivLongUnsigned64b32b2", func(t *testing.T) {
			q, r := hd.DivLongUnsigned64b32b2(x, y)

			if q != expQ {
				t.Errorf("x=%d y=%d: Q: exp=%d got=%d", x, y, expQ, q)
			}
			if r != expR {
				t.Errorf("x=%d y=%d: R: exp=%d got=%d", x, y, expR, r)
			}
		})

	})
}
