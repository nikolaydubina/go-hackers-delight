package hd_test

import (
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func doz[T hd.Integer](x, y T) T {
	if x > y {
		return x - y
	}
	return 0
}

func fuzzDOZ[T hd.Signed](t *testing.T, x, y T) {
	t.Run("doz", func(t *testing.T) {
		v := []struct {
			exp T
			got T
		}{
			{doz(x, y), hd.DifferenceOrZero(x, y)},
			{min(x, y), hd.Min(x, y)},
			{max(x, y), hd.Max(x, y)},
		}
		for i, q := range v {
			if q.exp != q.got {
				t.Error(i, q, x, y)
			}
		}
	})
}

func FuzzDOZ_int32(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, y := range fuzzInt32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		fuzzDOZ(t, x, y)

		t.Run("ranges", func(t *testing.T) {
			x &= (1<<31 - 1)
			y &= (1<<31 - 1)

			v := []struct {
				exp int32
				got int32
			}{
				{doz(x, y), hd.DifferenceOrZeroRanges(x, y)},
				{min(x, y), hd.MinRanges(x, y)},
				{max(x, y), hd.MaxRanges(x, y)},
			}
			for i, q := range v {
				if q.exp != q.got {
					t.Error(i, q, x, y)
				}
			}
		})
	})
}

func FuzzDOZ_int16(f *testing.F) { f.Fuzz(fuzzDOZ[int16]) }

func FuzzDOZ_int64(f *testing.F) { f.Fuzz(fuzzDOZ[int64]) }

func FuzzDOZ_uint32(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint32) {
		t.Run("ranges", func(t *testing.T) {
			x &= (1<<31 - 1)
			y &= (1<<31 - 1)

			v := []struct {
				exp uint32
				got uint32
			}{
				{doz(x, y), hd.DifferenceOrZeroRanges(x, y)},
				{min(x, y), hd.MinRanges(x, y)},
				{max(x, y), hd.MaxRanges(x, y)},
			}
			for i, q := range v {
				if q.exp != q.got {
					t.Error(i, q, x, y)
				}
			}
		})
	})
}
