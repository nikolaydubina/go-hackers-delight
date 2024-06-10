package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleDivSignedPowTwo_fixed5() {
	fmt.Println(hd.DivSignedPowTwo_fixed5(96))
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

		Q := x / (1 << k)
		R := x % (1 << k)

		// preventing Go sign conversion bit.
		if k == 31 {
			Q = -Q
		}

		vs := []int32{
			hd.DivSignedPowTwo(x, int(k)),
			hd.DivSignedPowTwo2(x, int(k)),
		}
		for i, v := range vs {
			if v != Q {
				t.Errorf("%d: (%d, %d) = %d; want %d", i, x, k, v, Q)
			}
		}

		if q, r := hd.DivModSignedPowTwo(x, int(k)); Q != q || R != r {
			t.Errorf("DivModPowTwo(%d, %d) = (%d, %d); want (%d, %d)", x, k, q, r, Q, R)
		}
	})
}

func FuzzDivMod_int32(f *testing.F) {
	for _, u := range fuzzInt32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		t.Run("mod3", func(t *testing.T) {
			Q, R := x/3, x%3

			if q, r := hd.DivMod3Signed(x); q != Q || r != R {
				t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, q, r, Q, R)
			}

			if q, r := hd.DivMod3Signed2(x); q != Q || r != R {
				t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, q, r, Q, R)
			}

			if q, r := hd.DivModConstSigned(x, 3); q != Q || r != R {
				t.Errorf("DivMod3(%d) = got(%d, %d); want (%d, %d)", x, q, r, Q, R)
			}
		})

		t.Run("mod5", func(t *testing.T) {
			q, r := x/5, x%5

			if Q, R := hd.DivMod5Signed(x); q != Q || r != R {
				t.Errorf("DivMod5(%d) = got(%d, %d); want (%d, %d)", x, Q, R, q, r)
			}

			if Q, R := hd.DivModConstSigned(x, 5); q != Q || r != R {
				t.Errorf("DivMod5(%d) = got(%d, %d); want (%d, %d)", x, Q, R, q, r)
			}
		})

		t.Run("mod7", func(t *testing.T) {
			Q, R := x/7, x%7

			if q, r := hd.DivMod7Signed(x); q != Q || r != R {
				t.Errorf("DivMod7(%d) = got(%d, %d); want (%d, %d)", x, q, r, Q, R)
			}

			if q, r := hd.DivModConstSigned(x, 7); q != Q || r != R {
				t.Errorf("DivMod7(%d) = got(%d, %d); want (%d, %d)", x, q, r, Q, R)
			}
		})
	})
}

func int32FromHex(x uint32) int32 { return int32(x) }

func TestMagicSigned(t *testing.T) {
	type tc struct {
		M int32
		s int32
	}
	tests := map[int32]tc{
		-5:  {M: int32FromHex(0x9999_9999), s: 1},
		-3:  {M: 0x5555_5555, s: 1},
		3:   {M: 0x5555_5556, s: 0},
		5:   {M: 0x6666_6667, s: 1},
		6:   {M: 0x2AAA_AAAB, s: 0},
		7:   {M: int32FromHex(0x9249_2493), s: 2},
		9:   {M: 0x38E3_8E39, s: 1},
		10:  {M: 0x6666_6667, s: 2},
		11:  {M: 0x2E8B_A2E9, s: 1},
		12:  {M: 0x2AAA_AAAB, s: 1},
		25:  {M: 0x51EB_851F, s: 3},
		125: {M: 0x1062_4DD3, s: 3},
		625: {M: 0x68DB_8BAD, s: 8},
	}
	for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		tests[-1<<k] = tc{M: 0x7FFF_FFFF, s: int32(k - 1)}
		tests[+1<<k] = tc{M: int32FromHex(0x8000_0001), s: int32(k - 1)}
	}
	for d, tc := range tests {
		t.Run(fmt.Sprintf("%v", d), func(t *testing.T) {
			M, s := hd.DivModConstSignedMagic(d)
			if M != tc.M || s != tc.s {
				t.Errorf("MagicSigned(%d) = (%x, %d); want (%x, %d)", d, M, s, tc.M, tc.s)
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
		if y == 0 {
			t.Skip()
		}
		// TODO: why integer signed division by negative constants is not working?
		if y < 0 {
			t.Skip()
		}

		Q, R := x/y, x%y
		if q, r := hd.DivModConstSigned(x, y); Q != q || R != r {
			t.Errorf("DivModConst(%d, %d) = (%d, %d); want (%d, %d)", x, y, q, r, Q, R)
		}
	})
}
