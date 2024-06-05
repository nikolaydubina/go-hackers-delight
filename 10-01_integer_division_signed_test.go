package hd_test

import (
	"fmt"
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
	})
}
