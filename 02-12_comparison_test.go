package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleEqual() {
	fmt.Println(hd.Equal(10, 10)>>31, hd.Equal(10, 11)>>31, hd.Equal(-10, -10)>>31, hd.Equal(-10, -11)>>31, hd.Equal(-10, 10)>>31)
	// Output: -1 0 -1 0 0
}

func isMostSignificantSet[T int32 | uint32](x T) bool { return !(x>>31 == 0) }

func FuzzCompareUint32(f *testing.F) {
	var vs = []uint32{
		0,
		1,
		math.MaxInt32,
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint32) {
		v := []struct {
			got uint32
			exp bool
		}{
			{hd.Less4(x, y), x < y},
			{hd.LessUnsigned(x, y), x < y},
			{hd.LessUnsigned2(x, y), x < y},
			{hd.LessOrEqualUnsigned(x, y), x <= y},
		}
		for i, q := range v {
			if isMostSignificantSet(q.got) != q.exp {
				t.Error(i, x)
			}
		}
	})
}

func FuzzCompareInt32(f *testing.F) {
	var vs = []int32{
		0,
		1,
		-1,
		math.MinInt32,
		math.MaxInt32,
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		v := []struct {
			got int32
			exp bool
		}{
			{hd.Equal(x, y), x == y},
			{hd.Equal2(x, y), x == y},
			{hd.Equal3(x, y), x == y},
			{hd.Equal4(x, y), x == y},
			{hd.Equal5(x, y), x == y},
			{hd.NotEqual(x, y), x != y},
			{hd.NotEqual2(x, y), x != y},
			{hd.NotEqual3(x, y), x != y},
			{hd.Less(x, y), x < y},
			{hd.Less2(x, y), x < y},
			{hd.Less4(x, y), x < y},
			{hd.LessOrEqual(x, y), x <= y},
			{hd.LessOrEqual2(x, y), x <= y},
		}
		for i, q := range v {
			if isMostSignificantSet(q.got) != q.exp {
				t.Error(i, x)
			}
		}
	})
}

func FuzzCompareZeroInt32(f *testing.F) {
	var vs = []int32{
		0,
		1,
		-1,
		math.MinInt32,
		math.MaxInt32,
	}
	for _, v := range vs {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		v := []struct {
			got int32
			exp bool
		}{
			{hd.EqualZero(x), x == 0},
			{hd.EqualZero2(x), x == 0},
			{hd.EqualZero3(x), x == 0},
			{hd.EqualZero4(x), x == 0},
			{hd.EqualZero5(x), x == 0},
			{hd.NotEqualZero(x), x != 0},
			{hd.NotEqualZero2(x), x != 0},
			{hd.NotEqualZero3(x), x != 0},
			{hd.NotEqualZero4(x), x != 0},
			{hd.LessZero(x), x < 0},
			{hd.LessOrEqualZero(x), x <= 0},
			{hd.LessOrEqualZero2(x), x <= 0},
			{hd.HigherZero(x), x > 0},
			{hd.HigherZero2(x), x > 0},
			{hd.HigherZero3(x), x > 0},
			{hd.HigherEqualZero(x), x >= 0},
		}
		for i, q := range v {
			if isMostSignificantSet(q.got) != q.exp {
				t.Error(i, x)
			}
		}
	})
}
