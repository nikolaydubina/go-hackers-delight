package hd_test

import (
	"math"
	"math/rand/v2"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func absFloat32(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func basicRSqrt(x float32) float32 { return float32(1 / math.Sqrt(float64(x))) }

func FuzzRSqrtFloat32(f *testing.F) {
	f.Fuzz(func(t *testing.T, x float32) {
		if x <= 0 || math.IsInf(float64(x), 1) || math.IsNaN(float64(x)) {
			t.Skip()
		}

		exp := basicRSqrt(x)
		got := hd.RSqrtFloat32(x)

		if absFloat32(exp-got)/got > hd.SqrtFloat32ErrorRate {
			t.Error("x", x, "exp", exp, "got", got)
		}
	})
}

func BenchmarkRSqrtFloat32(b *testing.B) {
	var out float32

	var vals []float32
	for i := 0; i < 10000; i++ {
		vals = append(vals, rand.Float32())
	}

	vs := []struct {
		name string
		f    func(x float32) float32
	}{
		{"basic", basicRSqrt},
		{"RSqrtFloat32", hd.RSqrtFloat32},
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
