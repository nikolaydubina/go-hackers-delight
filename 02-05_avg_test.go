package hd_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleAvgFloor() {
	fmt.Print(hd.AvgFloor[int32](-101, -200))
	// Output: -151
}

func ExampleAvgCeil() {
	fmt.Print(hd.AvgCeil[int32](-101, -200))
	// Output: -150
}

func FuzzAvgInt32(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, y := range fuzzInt32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		sum := int64(x) + int64(y)

		t.Run("ceil", func(t *testing.T) {
			var v int32 = int32(sum / 2)
			if sum%2 == 1 {
				v += 1
			}
			if avg := hd.AvgCeil(x, y); avg != v {
				t.Error("x", x, "y", y, "got", avg, "exp", v)
			}
		})

		t.Run("floor", func(t *testing.T) {
			var v int32 = int32(sum / 2)
			if sum%2 == -1 {
				v -= 1
			}
			if avg := hd.AvgFloor(x, y); avg != v {
				t.Error("x", x, "y", y, "got", avg, "exp", v)
			}
		})
	})
}

func FuzzAvgUint32(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint32) {
		sum := int64(x) + int64(y)

		t.Run("ceil", func(t *testing.T) {
			var v uint32 = uint32(sum / 2)
			if sum%2 == 1 {
				v += 1
			}
			if avg := hd.AvgCeil(x, y); avg != v {
				t.Error("x", x, "y", y, "got", avg, "exp", v)
			}
		})

		t.Run("floor", func(t *testing.T) {
			var v uint32 = uint32(sum / 2)
			if avg := hd.AvgFloor(x, y); avg != v {
				t.Error("x", x, "y", y, "got", avg, "exp", v)
			}
		})
	})
}

func avg[T hd.Integer](x, y T) T { return (x + y) / 2 }

func BenchmarkAvg(b *testing.B) {
	var out int32

	var vals []int32
	for i := 0; i < 1000; i++ {
		a := rand.Int32()
		b := rand.Int32()
		vals = append(vals, a-b)
	}

	vs := []struct {
		name string
		f    func(x, y int32) int32
	}{
		{"basic", avg[int32]},
		{"AvgFloor", hd.AvgFloor[int32]},
		{"AvgCeil", hd.AvgCeil[int32]},
	}
	for _, v := range vs {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i += len(vals) {
				for j := 0; j < len(vals)-1; j++ {
					out = v.f(vals[j], vals[j+1])
				}
			}
		})
	}

	if (out*2 - out - out) != 0 {
		b.Fatal("never")
	}
}
