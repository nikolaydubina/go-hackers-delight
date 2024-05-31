package hd_test

import (
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzArithmeticBoundPropagateLogical(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			for _, minx := range fuzzUint32 {
				for _, maxx := range fuzzUint32 {
					for _, miny := range fuzzUint32 {
						for _, maxy := range fuzzUint32 {
							f.Add(x, y, minx, maxx, miny, maxy)
						}
					}
				}
			}
		}
	}
	f.Fuzz(func(t *testing.T, x, y, minx, maxx, miny, maxy uint32) {
		if maxx < minx {
			maxx, minx = minx, maxx
		}
		if maxy < miny {
			maxy, miny = miny, maxy
		}
		if !(minx <= x && x <= maxx && miny <= y && y <= maxy) {
			t.Skip()
		}

		t.Run("or", func(t *testing.T) {
			or := x | y
			minOR := hd.MinOR(minx, maxx, miny, maxy)
			maxOR := hd.MaxOR(minx, maxx, miny, maxy)

			if !((or >= minOR) && (or <= maxOR)) {
				t.Errorf("%v %v %v %v %v %v %v %v", x, y, minx, maxx, miny, maxy, minOR, maxOR)
			}
		})

		t.Run("and", func(t *testing.T) {
			and := x & y
			minAND := hd.MinAND(minx, maxx, miny, maxy)
			maxAND := hd.MaxAND(minx, maxx, miny, maxy)

			if !((and >= minAND) && (and <= maxAND)) {
				t.Errorf("%v %v %v %v %v %v %v %v", x, y, minx, maxx, miny, maxy, minAND, maxAND)
			}
		})
	})
}
