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
		if u < v {
			u, v = v, u
		}
		if v == 0 {
			t.Skip()
		}

		Q := u / v
		R := u % v

		u16 := hd.IntToNx16b(u)
		v16 := hd.IntToNx16b(v)
		q16 := make([]uint16, len(u16))
		r16 := make([]uint16, len(u16))

		hd.DivModMultiWordUnsigned(q16, r16, u16, v16)

		if q := hd.Uint64FromNx16b(q16); q != Q {
			t.Errorf("u=%d %v v=%d %v: Q: exp=%d got=%d %v", u, u16, v, v16, Q, q, q16)
		}

		if r := hd.Uint64FromNx16b(r16); r != R {
			t.Errorf("u=%d %v v=%d %v: R: exp=%d got=%d %v", u, u16, v, v16, R, r, r16)
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

		Q := uint32(x / uint64(y))
		R := uint32(x % uint64(y))

		t.Run("DivLongUnsigned64b32b", func(t *testing.T) {
			// Overflow. This code does not work. Skip.
			if (x / uint64(y)) > math.MaxUint32 {
				t.Skip()
			}
			if q, r := hd.DivModLongUnsigned64b32b(x, y); q != Q || r != R {
				t.Errorf("x=%d y=%d: Q(exp=%d got=%d) R(exp=%d, got=%d)", x, y, Q, q, R, r)
			}
		})

		t.Run("DivLongUnsigned64b32b2", func(t *testing.T) {
			// Overflow. Special values returned.
			if (x / uint64(y)) > math.MaxUint32 {
				Q = math.MaxUint32
				R = math.MaxUint32
			}
			if q, r := hd.DivModLongUnsigned64b32b2(x, y); q != Q || r != R {
				t.Errorf("x=%d y=%d: Q(exp=%d got=%d) R(exp=%d, got=%d)", x, y, Q, q, R, r)
			}
		})
	})
}
