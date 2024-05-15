package ch2_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleAvg() {
	fmt.Print(ch2.Avg[int32](-100, -200))
	// Output: -150
}

func ExampleAvgCeil() {
	fmt.Print(ch2.Avg[int32](-100, -200))
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
		t.Run("ceil", func(t *testing.T) {
			avg := ch2.AvgCeil(x, y)

			var a, b int64 = int64(x), int64(y)
			sum := a + b
			var v int32 = int32(sum / 2)
			if sum%2 == 1 {
				v += 1
			}

			if avg != v {
				t.Error("x", x, "y", y, "got", avg, "exp", v)
			}
		})

		t.Run("floor", func(t *testing.T) {
			avg := ch2.Avg(x, y)

			var a, b int64 = int64(x), int64(y)
			sum := a + b
			var v int32 = int32(sum / 2)
			if sum%2 == -1 {
				v -= 1
			}

			if avg != v {
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
		t.Run("ceil", func(t *testing.T) {
			avg := ch2.AvgCeil(x, y)

			var a, b int64 = int64(x), int64(y)
			sum := a + b
			var v uint32 = uint32(sum / 2)
			if sum%2 == 1 {
				v += 1
			}

			if avg != v {
				t.Error("x", x, "y", y, "got", avg, "exp", v)
			}
		})

		t.Run("floor", func(t *testing.T) {
			avg := ch2.Avg(x, y)

			var a, b int64 = int64(x), int64(y)
			sum := a + b
			var v uint32 = uint32(sum / 2)

			if avg != v {
				t.Error("x", x, "y", y, "got", avg, "exp", v)
			}
		})
	})
}
