package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleUint64ToNx16b_small() {
	fmt.Println(hd.Uint64ToNx16b(5))
	// Output: [5]
}

func FuzzUint64ToNx16b(f *testing.F) {
	for _, x := range fuzzUint64 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint64) {
		if x != hd.Uint64FromNx16b(hd.Uint64ToNx16b(x)) {
			t.Error("x", x)
		}
	})
}

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

		u16 := hd.Uint64ToNx16b(u)
		v16 := hd.Uint64ToNx16b(v)
		q16 := make([]uint16, len(u16))
		r16 := make([]uint16, len(u16))

		hd.DivideMultiWord(q16, r16, u16, v16)

		if got := hd.Uint64FromNx16b(q16); got != expQ {
			t.Errorf("u=%d %v v=%d %v: Q: exp=%d got=%d %v", u, u16, v, v16, expQ, got, q16)
		}

		if got := hd.Uint64FromNx16b(r16); got != expR {
			t.Errorf("u=%d %v v=%d %v: R: exp=%d got=%d %v", u, u16, v, v16, expR, got, r16)
		}
	})
}
