package ch2_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/nikolaydubina/go-hackers-delight/ch2"
)

func ExampleEqual() {
	fmt.Println(ch2.Equal(10, 10)>>31, ch2.Equal(10, 11)>>31, ch2.Equal(-10, -10)>>31, ch2.Equal(-10, -11)>>31, ch2.Equal(-10, 10)>>31)
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
			{ch2.Less4(x, y), x < y},
			{ch2.LessUnsigned(x, y), x < y},
			{ch2.LessUnsigned2(x, y), x < y},
			{ch2.LessOrEqualUnsigned(x, y), x <= y},
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
			{ch2.Equal(x, y), x == y},
			{ch2.Equal2(x, y), x == y},
			{ch2.Equal3(x, y), x == y},
			{ch2.Equal4(x, y), x == y},
			{ch2.Equal5(x, y), x == y},
			{ch2.NotEqual(x, y), x != y},
			{ch2.NotEqual2(x, y), x != y},
			{ch2.NotEqual3(x, y), x != y},
			{ch2.Less(x, y), x < y},
			{ch2.Less2(x, y), x < y},
			{ch2.Less4(x, y), x < y},
			{ch2.LessOrEqual(x, y), x <= y},
			{ch2.LessOrEqual2(x, y), x <= y},
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
			{ch2.EqualZero(x), x == 0},
			{ch2.EqualZero2(x), x == 0},
			{ch2.EqualZero3(x), x == 0},
			{ch2.EqualZero4(x), x == 0},
			{ch2.EqualZero5(x), x == 0},
			{ch2.NotEqualZero(x), x != 0},
			{ch2.NotEqualZero2(x), x != 0},
			{ch2.NotEqualZero3(x), x != 0},
			{ch2.NotEqualZero4(x), x != 0},
			{ch2.LessZero(x), x < 0},
			{ch2.LessOrEqualZero(x), x <= 0},
			{ch2.LessOrEqualZero2(x), x <= 0},
			{ch2.HigherZero(x), x > 0},
			{ch2.HigherZero2(x), x > 0},
			{ch2.HigherZero3(x), x > 0},
			{ch2.HigherEqualZero(x), x >= 0},
		}
		for i, q := range v {
			if isMostSignificantSet(q.got) != q.exp {
				t.Error(i, x)
			}
		}
	})
}
