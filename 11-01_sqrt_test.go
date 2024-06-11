package hd_test

import (
	"math"
	"math/rand/v2"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func sqrtBasicFloat64(x uint32) uint32 { return uint32(math.Floor(math.Sqrt(float64(x)))) }

func FuzzSqrt(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		exp := sqrtBasicFloat64(x)
		got := [...]uint32{
			hd.SqrtNewton(x),
			hd.SqrtBinarySearch(x),
			hd.SqrtShiftAndSubtract(x),
		}
		for i, q := range got {
			if exp != q {
				t.Errorf("%d: exp(%v) q(%v) x(%x)", i, exp, q, x)
			}
		}
	})
}

func BenchmarkSqrt(b *testing.B) {
	var out uint32

	var vals []uint32
	for i := 0; i < 10000; i++ {
		vals = append(vals, rand.Uint32())
	}

	vs := []struct {
		name string
		f    func(x uint32) uint32
	}{
		{"basic", sqrtBasicFloat64},
		{"SqrtNewton", hd.SqrtNewton},
		{"SqrtBinarySearch", hd.SqrtBinarySearch},
		{"SqrtShiftAndSubtract", hd.SqrtShiftAndSubtract},
	}
	for _, v := range vs {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i += len(vals) {
				for j := 0; j < len(vals)-1; j++ {
					out = v.f(vals[j])
				}
			}
		})
	}

	if (out*2 - out - out) != 0 {
		b.Fatal("never")
	}
}
