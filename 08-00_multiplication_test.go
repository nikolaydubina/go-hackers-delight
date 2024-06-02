package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleMultiplyMultiWord() {
	u := []uint16{1, 2, 3, 4}
	v := []uint16{5, 6, 7}
	w := make([]uint16, len(u)+len(v))

	hd.MultiplyMultiWord(w, u, v)

	fmt.Println(w)
	// Output: [5 16 34 52 45 28 0]
}

func FuzzMultiplyMultiWord(f *testing.F) {
	for _, u := range fuzzInt32 {
		for _, v := range fuzzInt32 {
			f.Add(u, v)
		}
	}
	f.Fuzz(func(t *testing.T, u, v int32) {
		iu := int64(u)
		iv := int64(v)
		exp := iu * iv

		u16 := hd.IntToNx16b(iu)
		v16 := hd.IntToNx16b(iv)
		w16 := make([]uint16, len(u16)+len(v16))
		hd.MultiplyMultiWord(w16, u16, v16)

		if got := hd.Int64FromNx16b(w16); got != exp {
			t.Errorf("u=%d v=%d exp=%d got=%d", u, v, exp, got)
		}
	})
}

func FuzzMultiplyHighOrderSigned(f *testing.F) {
	for _, u := range fuzzInt32 {
		for _, v := range fuzzInt32 {
			f.Add(u, v)
		}
	}
	f.Fuzz(func(t *testing.T, u, v int32) {
		iu := int64(u)
		iv := int64(v)
		exp := int32(uint64(iu*iv) >> 32)

		if got := hd.MultiplyHighOrderSigned(u, v); got != exp {
			t.Errorf("u=%d v=%d exp=%d got=%d", u, v, exp, got)
		}
	})
}
