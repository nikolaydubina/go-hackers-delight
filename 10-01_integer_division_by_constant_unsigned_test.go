package hd_test

import (
	"strconv"
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

		expQ := x / (1 << k)
		expR := x % (1 << k)

		// preventing Go sign conversion bit.
		if k == 31 {
			expQ = -expQ
		}

		if q, r := hd.DivModUnsignedPowTwo(x, int(k)); expQ != q || expR != r {
			t.Errorf("DivModPow2(%d, %d) = (%d, %d); want (%d, %d)", x, k, q, r, expQ, expR)
		}
	})
}

func FuzzDivModUnsignedThree(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		expQ := x / 3
		expR := x % 3
		if q, r := hd.DivModUnsignedThree(x); expQ != q || expR != r {
			t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, q, r, expQ, expR)
		}
	})
}

func FuzzDivModUnsignedSeven(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		expQ := x / 7
		expR := x % 7
		if q, r := hd.DivModUnsignedSeven(x); expQ != q || expR != r {
			t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, q, r, expQ, expR)
		}
	})
}

func TestMulUnsignedMagic(t *testing.T) {
	type tc struct {
		d uint32
		M uint32
		s int32
		a int32
	}
	tests := []tc{
		{d: 1, M: 0x0000_0000, a: 1, s: 0},
		{d: 3, M: 0xAAAA_AAAB, a: 0, s: 1},
		{d: 5, M: 0xCCCC_CCCD, a: 0, s: 2},
		{d: 7, M: 0x2492_4925, a: 1, s: 3},
		{d: 9, M: 0x38E3_8E39, a: 0, s: 1},
		{d: 10, M: 0xCCCC_CCCD, a: 0, s: 3},
		{d: 11, M: 0xBA2E_8BA3, a: 0, s: 3},
		{d: 12, M: 0xAAAA_AAAB, a: 0, s: 3},
		{d: 25, M: 0x51EB_851F, a: 0, s: 3},
		{d: 125, M: 0x1062_4DD3, a: 0, s: 3},
		{d: 625, M: 0xD1B7_1759, a: 0, s: 9},
	}
	for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		tests = append(tests, tc{d: 1 << k, M: 1 << (32 - k), a: 0, s: 0})
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			M, a, s := hd.DivModUnsignedConstMagic(tc.d)
			if M != tc.M || s != tc.s || tc.a != a {
				t.Errorf("MulUnsignedMagic(%d) = got(%d, %d, %d); want (%d, %d, %d)", tc.d, M, a, s, tc.M, tc.a, tc.s)
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

		expQ := x / y
		expR := x % y

		if q, r := hd.DivModUnsignedConst(x, y); expQ != q || expR != r {
			M, a, s := hd.DivModUnsignedConstMagic(y)
			t.Errorf("DivModConst(%d, %d) = got(%d, %d); want (%d, %d) M=%d, a=%d, s=%d", x, y, q, r, expQ, expR, M, a, s)
		}
	})
}
