package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleEqual_equal() {
	fmt.Println(hd.Equal(10, 10) >> 31)
	// Output: -1
}

func ExampleEqual_not_equal() {
	fmt.Println(hd.Equal(10, 11) >> 31)
	// Output: 0
}

func FuzzCompareUint32(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y uint32) {
		v := []struct {
			got uint32
			exp bool
		}{
			{hd.Equal5(x, y), x == y},
			{hd.NotEqual3(x, y), x != y},
			{hd.EqualZero4(x), x == 0},
			{hd.NotEqualZero3(x), x != 0},
			{hd.Less4(x, y), x < y},
			{hd.LessUnsigned(x, y), x < y},
			{hd.LessUnsigned2(x, y), x < y},
			{hd.LessOrEqualUnsigned(x, y), x <= y},
		}
		for i, q := range v {
			if hd.IsMostSignificantSet(q.got) != q.exp {
				t.Error(i, x)
			}
		}
	})
}

func fuzzCompare[T hd.Signed](t *testing.T, x, y T) {
	v := []struct {
		got T
		exp bool
	}{
		{hd.Equal(x, y), x == y},
		{hd.Equal5(x, y), x == y},
		{hd.EqualZero5(x), x == 0},
		{hd.NotEqual(x, y), x != y},
		{hd.NotEqual3(x, y), x != y},
		{hd.NotEqualZero3(x), x != 0},
		{hd.Less(x, y), x < y},
		{hd.Less2(x, y), x < y},
		{hd.Less3(x, y), x < y},
		{hd.Less4(x, y), x < y},
		{hd.LessOrEqual(x, y), x <= y},
		{hd.LessOrEqual2(x, y), x <= y},
	}
	for i, q := range v {
		if (q.got < 0) != q.exp {
			t.Error(i, x)
		}
	}
}

func FuzzCompare_int32(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, y := range fuzzInt32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		fuzzCompare(t, x, y)

		// int32 only code
		got := []struct {
			got int32
			exp bool
		}{
			{hd.Equal2(x, y), x == y},
			{hd.Equal3(x, y), x == y},
			{hd.Equal4(x, y), x == y},
			{hd.NotEqual2(x, y), x != y},
			{hd.NotEqualZero4(x), x != 0},
		}
		for i, q := range got {
			if (q.got < 0) != q.exp {
				t.Error(i, x)
			}
		}
	})
}

func FuzzCompare_int16(f *testing.F) { f.Fuzz(fuzzCompare[int64]) }

func FuzzCompare_int64(f *testing.F) { f.Fuzz(fuzzCompare[int64]) }

func fuzzCompareZero[T hd.Signed](t *testing.T, x T) {
	got := []struct {
		got T
		exp bool
	}{
		{hd.EqualZero(x), x == 0},
		{hd.EqualZero4(x), x == 0},
		{hd.EqualZero5(x), x == 0},
		{hd.NotEqualZero(x), x != 0},
		{hd.NotEqualZero3(x), x != 0},
		{hd.LessZero(x), x < 0},
		{hd.LessOrEqualZero(x), x <= 0},
		{hd.LessOrEqualZero2(x), x <= 0},
		{hd.HigherZero(x), x > 0},
		{hd.HigherZero2(x), x > 0},
		{hd.HigherZero3(x), x > 0},
		{hd.HigherEqualZero(x), x >= 0},
	}
	for i, q := range got {
		if (q.got < 0) != q.exp {
			t.Error(i, x)
		}
	}
}

func FuzzCompareZero_int32(f *testing.F) {
	for _, v := range fuzzInt32 {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		fuzzCompareZero(t, x)

		// int32 only code
		got := []struct {
			got int32
			exp bool
		}{
			{hd.EqualZero2(x), x == 0},
			{hd.EqualZero3(x), x == 0},
			{hd.NotEqualZero2(x), x != 0},
			{hd.NotEqualZero4(x), x != 0},
		}
		for i, q := range got {
			if (q.got < 0) != q.exp {
				t.Error(i, x)
			}
		}
	})
}

func FuzzCompareZero_int16(f *testing.F) { f.Fuzz(fuzzCompareZero[int16]) }

func FuzzCompareZero_int64(f *testing.F) { f.Fuzz(fuzzCompareZero[int64]) }
