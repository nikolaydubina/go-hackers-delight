package hd_test

import (
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzExchangeRegisters(f *testing.F) {
	var vs = []int32{
		0,
		1,
		-1,
		math.MinInt32,
		math.MinInt32 / 2,
		math.MinInt32 + 1,
		math.MaxInt32,
		math.MaxInt32 / 2,
		math.MaxInt32 - 1,
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		ex, ey := y, x
		gx, gy := hd.ExchangeRegisters(x, y)
		if ex != gx || ey != gy {
			t.Error(x, y, gx, gy)
		}
	})
}
