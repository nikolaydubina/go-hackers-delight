package hd_test

import (
	"fmt"
	"math/bits"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleMulMultiWord() {
	u := []uint16{1, 2, 3, 4}
	v := []uint16{5, 6, 7}
	w := make([]uint16, len(u)+len(v))

	hd.MulMultiWord(w, u, v)

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
		hd.MulMultiWord(w16, u16, v16)

		if got := hd.Int64FromNx16b(w16); got != exp {
			t.Errorf("u=%d v=%d exp=%d got=%d", u, v, exp, got)
		}
	})
}

func FuzzMultiplyHighOrder32_int32(f *testing.F) {
	for _, u := range fuzzInt32 {
		for _, v := range fuzzInt32 {
			f.Add(u, v)
		}
	}
	f.Fuzz(func(t *testing.T, u, v int32) {
		exp := int32((int64(u) * int64(v)) >> 32)
		got := hd.MultiplyHighOrder32(u, v)
		if got != exp {
			t.Errorf("u=%d v=%d exp=%d got=%d", u, v, exp, got)
		}
	})
}

func FuzzMultiplyHighOrder32_uint32(f *testing.F) {
	for _, u := range fuzzUint32 {
		for _, v := range fuzzUint32 {
			f.Add(u, v)
		}
	}
	f.Fuzz(func(t *testing.T, u, v uint32) {
		exp, _ := bits.Mul32(u, v)
		if got := hd.MultiplyHighOrder32(u, v); got != exp {
			t.Errorf("u=%d v=%d exp=%d got=%d", u, v, exp, got)
		}
	})
}

func FuzzMultiplyHighOrder64_uint64(f *testing.F) {
	for _, u := range fuzzUint64 {
		for _, v := range fuzzUint64 {
			f.Add(u, v)
		}
	}
	f.Fuzz(func(t *testing.T, u, v uint64) {
		exp, _ := bits.Mul64(u, v)
		if got := hd.MultiplyHighOrder64(u, v); got != exp {
			t.Errorf("u=%d v=%d exp=%d got=%d", u, v, exp, got)
		}
	})
}
