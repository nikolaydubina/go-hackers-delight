package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleIsAddOverflow() {
	fmt.Println(hd.IsAddOverflow(math.MaxInt32, 1)>>31, hd.IsAddOverflow(10, 1)>>31)
	// Output: -1 0
}

func ExampleIsSubOverflow() {
	fmt.Println(hd.IsSubOverflow2(math.MinInt32, 1)>>31, hd.IsSubOverflow2(-10, 1)>>31)
	// Output: -1 0
}

func ExampleIsAddOverflowUnsigned() {
	fmt.Println(hd.IsAddOverflowUnsigned(math.MaxUint32, 1)>>31, hd.IsAddOverflowUnsigned(10, 1)>>31)
	// Output: 1 0
}

func ExampleIsSubOverflowUnsigned() {
	fmt.Println(hd.IsSubOverflowUnsigned(0, 1)>>31, hd.IsSubOverflowUnsigned(10, 1)>>31)
	// Output: 1 0
}

func ExampleIsAddOverflowUnsigned4() {
	var x uint32 = 1
	fmt.Println(hd.IsAddOverflowUnsigned4(math.MaxUint32, 1, math.MaxUint32+x), hd.IsAddOverflowUnsigned4(10, 1, 10+x))
	// Output: true false
}

func ExampleIsSubOverflowUnsigned4() {
	var x uint32 = 1
	fmt.Println(hd.IsSubOverflowUnsigned4(0, 1, 0-x), hd.IsSubOverflowUnsigned4(10, 1, 10-x))
	// Output: true false
}

func FuzzOverflowInt32(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, y := range fuzzInt32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		a, b := float64(x), float64(y)
		sum := a + b
		sub := a - b
		mul := a * b
		div := a / b
		sumOverflow := sum > math.MaxInt32 || sum < math.MinInt32
		subOverflow := sub > math.MaxInt32 || sub < math.MinInt32
		mulOverflow := mul > math.MaxInt32 || mul < math.MinInt32
		divOverflow := div > math.MaxInt32 || div < math.MinInt32

		v := []struct {
			exp bool
			got int32
		}{
			{sumOverflow, hd.IsAddOverflow(x, y)},
			{sumOverflow, hd.IsAddOverflow2(x, y)},
			{subOverflow, hd.IsSubOverflow(x, y)},
			{subOverflow, hd.IsSubOverflow2(x, y)},
		}
		for i, q := range v {
			if hd.IsMostSignificantSet(q.got) != q.exp {
				t.Error(i, x, y)
			}
		}

		if mulOverflow != hd.IsMulOverflow(x, y) {
			t.Error("mul", x, y)
		}

		// in Go 0/0 is panic at runtime, therefore overflow value is not determined for this case
		if x > 0 && divOverflow != hd.IsDivOverflow(x, y) {
			t.Error("div", x, y)
		}
	})
}

func FuzzOverflowUint32(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, y := range fuzzUint32 {
			f.Add(x, y)
		}
	}

	f.Add(uint32(0x0000FFFF), uint32(0x0001FFFF)) // 4 * 4 + 3 * 4 + 3 = 31 bits leading zeros n + m

	f.Fuzz(func(t *testing.T, x, y uint32) {
		a, b := uint64(x), uint64(y)
		sum := a + b
		sub := a - b
		mul := a * b
		div := uint64(0)
		if b > 0 {
			div = a / b
		}
		sumOverflow := sum > math.MaxUint32
		subOverflow := sub > math.MaxUint32
		mulOverflow := mul > math.MaxUint32
		divOverflow := div > math.MaxUint32

		v := []struct {
			exp bool
			got bool
		}{
			{sumOverflow, hd.IsMostSignificantSet(hd.IsAddOverflowUnsigned(x, y))},
			{sumOverflow, hd.IsMostSignificantSet(hd.IsAddOverflowUnsigned2(x, y))},
			{sumOverflow, hd.IsAddOverflowUnsigned3(x, y)},
			{sumOverflow, hd.IsAddOverflowUnsigned4(x, y, x+y)},
			{subOverflow, hd.IsMostSignificantSet(hd.IsSubOverflowUnsigned(x, y))},
			{subOverflow, hd.IsMostSignificantSet(hd.IsSubOverflowUnsigned2(x, y))},
			{subOverflow, hd.IsSubOverflowUnsigned3(x, y)},
			{subOverflow, hd.IsSubOverflowUnsigned4(x, y, x-y)},
			{mulOverflow, hd.IsMulOverflowUnsigned(x, y)},
			{mulOverflow, hd.IsMulOverflowUnsignedNLZ(x, y)},
		}
		for i, q := range v {
			if q.got != q.exp {
				t.Error(i, x, y)
			}
		}

		t.Run("IsDivOverflowUnsigned", func(t *testing.T) {
			if y == 0 {
				t.Skip("cannot divide by 0")
			}

			if divOverflow != hd.IsDivOverflowUnsigned(x, y) {
				t.Error("div", x, y)
			}
		})
	})
}
