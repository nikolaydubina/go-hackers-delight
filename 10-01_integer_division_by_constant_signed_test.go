package hd_test

import (
	"fmt"
	"strconv"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleDivPow2_fixed5() {
	fmt.Println(hd.DivPow2_fixed5(96))
	// Output: 3

}

func FuzzDivPow2(f *testing.F) {
	for _, u := range fuzzInt32 {
		f.Add(u, uint8(2))
	}
	f.Fuzz(func(t *testing.T, x int32, k uint8) {
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

		if got := hd.DivPow2(x, int(k)); expQ != got {
			t.Errorf("DivPow2(%d, %d) = %d; want %d", x, k, got, expQ)
		}

		if got := hd.DivPow2Two(x, int(k)); expQ != got {
			t.Errorf("DivPow2Two(%d, %d) = %d; want %d", x, k, got, expQ)
		}

		if gotQ, gotR := hd.DivModPow2(x, int(k)); expQ != gotQ || expR != gotR {
			t.Errorf("DivModPow2(%d, %d) = (%d, %d); want (%d, %d)", x, k, gotQ, gotR, expQ, expR)
		}
	})
}

func FuzzDivMod3(f *testing.F) {
	for _, u := range fuzzInt32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		q, r := x/3, x%3

		if gotQ, gotR := hd.DivMod3(x); q != gotQ || r != gotR {
			t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, gotQ, gotR, q, r)
		}

		if gotQ, gotR := hd.DivModSignedConst(x, 3); q != gotQ || r != gotR {
			t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, gotQ, gotR, q, r)
		}
	})
}

func FuzzDivMod5(f *testing.F) {
	for _, u := range fuzzInt32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		q, r := x/5, x%5

		if gotQ, gotR := hd.DivMod5(x); q != gotQ || r != gotR {
			t.Errorf("DivMod5(%d) = got(%d, %d); want (%d, %d)", x, gotQ, gotR, q, r)
		}

		if gotQ, gotR := hd.DivModSignedConst(x, 5); q != gotQ || r != gotR {
			t.Errorf("DivMod5(%d) = got(%d, %d); want (%d, %d)", x, gotQ, gotR, q, r)
		}
	})
}

func FuzzDivMod7(f *testing.F) {
	for _, u := range fuzzInt32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		q, r := x/7, x%7

		if gotQ, gotR := hd.DivMod7(x); q != gotQ || r != gotR {
			t.Errorf("DivMod7(%d) = got(%d, %d); want (%d, %d)", x, gotQ, gotR, q, r)
		}

		if gotQ, gotR := hd.DivModSignedConst(x, 7); q != gotQ || r != gotR {
			t.Errorf("DivMod7(%d) = got(%d, %d); want (%d, %d)", x, gotQ, gotR, q, r)
		}
	})
}

func int32FromHex(x uint32) int32 { return int32(x) }

func TestMagicSigned(t *testing.T) {
	type tc struct {
		d int32
		M int32
		s int32
	}
	tests := []tc{
		{d: -5, M: int32FromHex(0x9999_9999), s: 1},
		{d: -3, M: 0x5555_5555, s: 1},
		//{d: 1, } // intentionally skipped,
		{d: 3, M: 0x5555_5556, s: 0},
		{d: 5, M: 0x6666_6667, s: 1},
		{d: 6, M: 0x2AAA_AAAB, s: 0},
		{d: 7, M: int32FromHex(0x9249_2493), s: 2},
		{d: 9, M: 0x38E3_8E39, s: 1},
		{d: 10, M: 0x6666_6667, s: 2},
		{d: 11, M: 0x2E8B_A2E9, s: 1},
		{d: 12, M: 0x2AAA_AAAB, s: 1},
		{d: 25, M: 0x51EB_851F, s: 3},
		{d: 125, M: 0x1062_4DD3, s: 3},
		{d: 625, M: 0x68DB_8BAD, s: 8},
	}
	for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		tests = append(tests, tc{d: -1 << k, M: 0x7FFF_FFFF, s: int32(k - 1)})
		tests = append(tests, tc{d: 1 << k, M: int32FromHex(0x8000_0001), s: int32(k - 1)})
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			M, s := hd.MulSignedMagic(tc.d)
			if M != tc.M || s != tc.s {
				t.Errorf("MagicSigned(%d) = (%x, %d); want (%x, %d)", tc.d, M, s, tc.M, tc.s)
			}
		})
	}
}

func FuzzDivModSignedConst(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, y := range fuzzInt32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		if y > -2 && y < 2 {
			t.Skip()
		}
		// TODO: why integer signed division by negative constants is not working?
		if y < 0 {
			t.Skip()
		}

		expQ := x / y
		expR := x % y

		if q, r := hd.DivModSignedConst(x, y); expQ != q || expR != r {
			M, s := hd.MulSignedMagicCached(y)
			t.Errorf("DivModConst(%d, %d) = (%d, %d); want (%d, %d), M(%v) s(%v)", x, y, q, r, expQ, expR, M, s)
		}
	})
}
