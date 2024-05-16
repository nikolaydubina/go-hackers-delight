package hd_test

import (
	"fmt"
	"math"
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
	var vs = []int32{
		0,
		1,
		-1,
		math.MaxInt32,
		math.MinInt32,
	}
	for _, x := range vs {
		for _, y := range vs {
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
	var vs = []uint32{
		0,
		1,
		math.MaxUint32,
	}
	for _, x := range vs {
		for _, y := range vs {
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
