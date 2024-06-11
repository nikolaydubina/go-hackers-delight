package hd_test

import (
	"fmt"
	"math"
	"math/rand/v2"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleCbrt() {
	fmt.Println(hd.Cbrt(64))
	// Output: 4
}

func cbrtBasicFloat64(x uint32) uint32 { return uint32(math.Floor(math.Pow(float64(x), 1./3))) }

func FuzzCbrt(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		// detecting underflow
		if v := math.Pow(float64(x), 1./3); math.Abs(math.Round(v)-v) < 0.0000001 {
			t.Skip()
		}

		exp := cbrtBasicFloat64(x)
		got := hd.Cbrt(x)
		if exp != got {
			t.Errorf("exp(%v, float=%v) got(%v) x(%d)", exp, math.Pow(float64(x), 1./3), got, x)
		}
	})
}

func BenchmarkCbrt(b *testing.B) {
	var out uint32

	var vals []uint32
	for i := 0; i < 10000; i++ {
		vals = append(vals, rand.Uint32())
	}

	vs := []struct {
		name string
		f    func(x uint32) uint32
	}{
		{"basic", cbrtBasicFloat64},
		{"Cbrt", hd.Cbrt},
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

func powBasic[T hd.Integer](x T, n uint) T { return T(math.Floor(math.Pow(float64(x), float64(n)))) }

func FuzzPow(f *testing.F) {
	for _, u := range fuzzInt32 {
		f.Add(u, uint(2))
	}
	f.Fuzz(func(t *testing.T, x int32, n uint) {
		if v := math.Pow(float64(x), float64(n)); v > math.MaxInt32 || v < math.MinInt32 {
			t.Skip()
		}

		exp := powBasic(x, n)
		got := hd.Pow(x, n)
		if exp != got {
			t.Errorf("exp(%v) got(%v) x(%x) n(%x)", exp, got, x, n)
		}
	})
}

func BenchmarkPow(b *testing.B) {
	var out int32

	type tc struct {
		x int32
		n uint
	}

	var vals []tc
	for x, n := int32(0), uint(0); len(vals) < 10000; x, n = rand.Int32(), uint(rand.Uint32()) {
		if v := math.Pow(float64(x), float64(n)); !(v > math.MaxInt32 || v < math.MinInt32) {
			continue
		}
		vals = append(vals, tc{x, n})
	}
	b.Logf("num_vals(%d)", len(vals))

	vs := []struct {
		name string
		f    func(x int32, n uint) int32
	}{
		{"basic", powBasic[int32]},
		{"Pow", hd.Pow[int32]},
	}
	for _, v := range vs {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i += len(vals) {
				for j := 0; j < len(vals)-1; j++ {
					out = v.f(vals[j].x, vals[j].n)
				}
			}
		})
	}

	if (out*2 - out - out) != 0 {
		b.Fatal("never")
	}
}

func ExampleLog2_zero() {
	// This is special condition. There is a case to extend this definition to be valid.
	// However, it is not universally accepted.
	fmt.Println(hd.Log2(0))
	// Output: 4294967295
}

func ExampleLog2_one() {
	fmt.Println(hd.Log2(4))
	// Output: 2
}

func ExampleLog2_four() {
	fmt.Println(hd.Log2(4))
	// Output: 2
}

func ExampleLog2_nine() {
	fmt.Println(hd.Log2(9))
	// Output: 3
}

func log2Basic(x uint32) uint32 { return uint32(math.Floor(math.Log2(float64(x)))) }

func log10Basic(x uint32) uint32 { return uint32(math.Floor(math.Log10(float64(x)))) }

func FuzzLog_uint32(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		if x == 0 {
			t.Skip()
		}
		tests := []struct {
			name string
			exp  uint32
			got  uint32
		}{
			{"log2", log2Basic(x), hd.Log2(x)},
			{"log10", log10Basic(x), hd.Log10x32(x)},
		}
		for _, test := range tests {
			if test.exp != test.got {
				t.Errorf("%s: exp(%v) got(%v) x(%v)", test.name, test.exp, test.got, x)
			}
		}
	})
}

func log2Basic64(x uint64) uint64 { return uint64(math.Floor(math.Log2(float64(x)))) }

func log10Basic64(x uint64) uint64 { return uint64(math.Floor(math.Log10(float64(x)))) }

func FuzzLog_uint64(f *testing.F) {
	for _, u := range fuzzUint64 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint64) {
		if x == 0 {
			t.Skip()
		}
		// float64 based code can is not correct with such large uint64 numbers for this function. so skipping.
		// this was verified with big arithmetics in wolfram alpha.
		if x > (math.MaxUint32 >> 1) {
			t.Skip()
		}
		tests := []struct {
			name string
			exp  uint64
			got  uint64
		}{
			{name: "log2", exp: log2Basic64(x), got: hd.Log2x64(x)},
			{name: "log10", exp: log10Basic64(x), got: hd.Log10x64(x)},
		}
		for _, test := range tests {
			if test.exp != test.got {
				t.Errorf("%s: exp(%v) got(%v) x(%v, %x) x_leading_zeros(%d)", test.name, test.exp, test.got, x, x, hd.LeadingZerosUint64(x))
			}
		}
	})
}

func BenchmarkLog(b *testing.B) {
	b.Run("uint32", func(b *testing.B) {
		var out uint32

		var vals []uint32
		for x := uint32(0); len(vals) < 10000; x = rand.Uint32() {
			vals = append(vals, x)
		}
		b.Logf("num_vals(%d)", len(vals))

		vs := []struct {
			name string
			f    func(x uint32) uint32
		}{
			{"2/basic", log2Basic},
			{"2/Log2", hd.Log2},
			{"10/basic", log10Basic},
			{"10/Log10", hd.Log10x32},
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
	})

	b.Run("uint64", func(b *testing.B) {
		var out uint64

		var vals []uint64
		for x := uint64(0); len(vals) < 10000; x = rand.Uint64() {
			vals = append(vals, x)
		}
		b.Logf("num_vals(%d)", len(vals))

		vs := []struct {
			name string
			f    func(x uint64) uint64
		}{
			{"2/basic", log2Basic64},
			{"2/Log2", hd.Log2x64},
			{"10/basic", log10Basic64},
			{"10/Log10", hd.Log10x64},
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
	})
}
