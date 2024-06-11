package hd_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleAbs() {
	fmt.Print(hd.Abs(-42))
	// Output: 42
}

func ExampleNAbs() {
	fmt.Print(hd.NAbs(-42))
	// Output: -42
}

func ExampleAbsDiff() {
	fmt.Print(hd.AbsDiff(1, 100))
	// Output: 99
}

func abs[T hd.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func fuzzAbs[T hd.Signed](t *testing.T, x T) {
	t.Run("abs", func(t *testing.T) {
		got := []T{
			hd.Abs(x),
			hd.Abs2(x),
			hd.Abs3(x),
			hd.Abs4(x),
		}
		for i, v := range got {
			if abs := abs(x); v != abs {
				t.Error(i, "x", x, "exp", abs, "got", v)
			}
		}
	})

	t.Run("nabs", func(t *testing.T) {
		got := []T{
			hd.NAbs(x),
			hd.NAbs2(x),
			hd.NAbs3(x),
		}
		for i, v := range got {
			if abs := abs(x); v != -abs {
				t.Error(i, "x", x, "exp", -abs, "got", v)
			}
		}
	})
}

func FuzzAbsNormal_int32(f *testing.F) {
	for _, x := range fuzzInt32 {
		f.Add(x)
	}
	f.Fuzz(fuzzAbs[int32])
}

func FuzzAbsNormal_int16(f *testing.F) { f.Fuzz(fuzzAbs[int16]) }

func FuzzAbsNormal_int64(f *testing.F) { f.Fuzz(fuzzAbs[int32]) }

func FuzzAbsDiffInt(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, y := range fuzzInt32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		vs := []int32{
			hd.AbsDiff2(x, y),
		}
		for i, v := range vs {
			if v != hd.AbsDiff(x, y) {
				t.Error(i, x, y, v)
			}
		}
	})
}

func FuzzAbsDiffUint(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint32) {
		ux, uy := uint64(x), uint64(y)
		exp := uint32(max(ux, uy) - min(ux, uy))

		if hd.AbsDiffUnsigned(x, y) != exp {
			t.Error(x, y, hd.AbsDiffUnsigned(x, y), exp)
		}
	})
}

func BenchmarkAbs(b *testing.B) {
	var out int32

	var vals []int32
	for i := 0; i < 1000; i++ {
		a := rand.Int32()
		b := rand.Int32()
		vals = append(vals, a-b)
	}

	// info on random numbers
	countPositive := 0
	for _, v := range vals {
		if v > 0 {
			countPositive++
		}
	}
	b.Logf("positive %d/%d = %.2f", countPositive, len(vals), float64(countPositive)/float64(len(vals)))

	vs := []struct {
		name string
		f    func(int32) int32
	}{
		{"basic", abs[int32]},
		{"Abs", hd.Abs[int32]},
		{"Abs2", hd.Abs2[int32]},
		{"Abs3", hd.Abs3[int32]},
		{"Abs4", hd.Abs4[int32]},
		{"AbsFastMul", hd.AbsFastMul},
	}
	for _, v := range vs {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i += len(vals) {
				for j := 0; j < len(vals); j++ {
					out = v.f(vals[j])
				}
			}
		})
	}

	if (out*2 - out - out) != 0 {
		b.Fatal("never")
	}
}
