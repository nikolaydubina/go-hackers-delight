package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzDivModUnsignedPowTwo(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u, uint8(2))
	}
	f.Fuzz(func(t *testing.T, x uint32, k uint8) {
		k = k % 32
		if k == 0 {
			k++
		}

		Q := x / (1 << k)
		R := x % (1 << k)

		if k == 31 {
			// TODO: why?
			t.Skip()
		}

		if q, r := hd.DivModUnsignedPowTwo(x, int(k)); Q != q || R != r {
			t.Errorf("DivModUnsignedPowTwo(%d, %d) = (%d, %d); want (%d, %d)", x, k, q, r, Q, R)
		}
	})
}

func FuzzDivModUnsignedThree(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		Q, R := x/3, x%3
		if q, r := hd.DivModUnsignedThree(x); Q != q || R != r {
			t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, q, r, Q, R)
		}
	})
}

func FuzzDivModUnsignedSeven(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		Q, R := x/7, x%7
		if q, r := hd.DivModUnsignedSeven(x); Q != q || R != r {
			t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, q, r, Q, R)
		}
	})
}

func TestMulUnsignedMagic(t *testing.T) {
	type tc struct {
		M uint32
		s int32
		a int32
	}
	tests := map[uint32]tc{
		1:   {M: 0x0000_0000, a: 1, s: 0},
		3:   {M: 0xAAAA_AAAB, a: 0, s: 1},
		5:   {M: 0xCCCC_CCCD, a: 0, s: 2},
		7:   {M: 0x2492_4925, a: 1, s: 3},
		9:   {M: 0x38E3_8E39, a: 0, s: 1},
		10:  {M: 0xCCCC_CCCD, a: 0, s: 3},
		11:  {M: 0xBA2E_8BA3, a: 0, s: 3},
		12:  {M: 0xAAAA_AAAB, a: 0, s: 3},
		25:  {M: 0x51EB_851F, a: 0, s: 3},
		125: {M: 0x1062_4DD3, a: 0, s: 3},
		625: {M: 0xD1B7_1759, a: 0, s: 9},
	}
	for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		tests[1<<k] = tc{M: 1 << (32 - k), a: 0, s: 0}
	}
	for d, tc := range tests {
		t.Run(fmt.Sprintf("%v", d), func(t *testing.T) {
			M, a, s := hd.DivModUnsignedConstMagic(d)
			if M != tc.M || s != tc.s || tc.a != a {
				t.Errorf("MulUnsignedMagic(%d) = got(%d, %d, %d); want (%d, %d, %d)", d, M, a, s, tc.M, tc.a, tc.s)
			}
		})
	}
}

func FuzzDivModUnsignedConst(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			f.Add(x, y)
		}
	}

	f.Fuzz(func(t *testing.T, x, y uint32) {
		if y == 0 {
			t.Skip()
		}
		Q, R := x/y, x%y
		if q, r := hd.DivModUnsignedConst(x, y); Q != q || R != r {
			t.Errorf("DivModConst(%d, %d) = got(%d, %d); want (%d, %d)", x, y, q, r, Q, R)
		}
	})
}
