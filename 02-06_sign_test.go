package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleExtendSign7_notExtended() {
	fmt.Printf("%08b", hd.ExtendSign7(0b0010_1010))
	// Output: 00101010
}

func ExampleExtendSign7_extended() {
	fmt.Printf("%08b", hd.ExtendSign7(0b1110_1010))
	// Output: 11111111111111111111111111101010
}

func ExampleExtendSign7Two_notExtended() {
	fmt.Printf("%08b", hd.ExtendSign7Two(0b0010_1010))
	// Output: 00101010
}

func ExampleExtendSign7Two_extended() {
	fmt.Printf("%08b", hd.ExtendSign7Two(0b1110_1010))
	// Output: 11111111111111111111111111101010
}

func ExampleExtendSign7Three_notExtended() {
	fmt.Printf("%08b", hd.ExtendSign7Three(0b0010_1010))
	// Output: 00101010
}

func ExampleExtendSign7Three_extended() {
	fmt.Printf("%08b", hd.ExtendSign7Three(0b1110_1010))
	// Output: 11111111111111111111111111101010
}

func ExampleShiftRightSignedFromUnsigned() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned2() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned2(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned3() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned3(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned4() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned4(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleShiftRightSignedFromUnsigned5() {
	fmt.Printf("%08b", hd.ShiftRightSignedFromUnsigned5(0b11111111111111111111111111101010, 2))
	// Output: 11111111111111111111111111111010
}

func ExampleSign() {
	fmt.Println(hd.Sign(-10), hd.Sign(10), hd.Sign(0))
	// Output: -1 1 0
}

func ExampleIsMostSignificantSet_int32() {
	fmt.Println(hd.IsMostSignificantSet(int32(-1)), hd.IsMostSignificantSet(int32(1)), hd.IsMostSignificantSet(int32(math.MaxInt32)))
	// Output: true false false
}

func ExampleIsMostSignificantSet_uint32() {
	fmt.Println(hd.IsMostSignificantSet(uint32(0xFFFFFFFF)), hd.IsMostSignificantSet(uint32(10)))
	// Output: true false
}

func FuzzSign(f *testing.F) {
	for _, x := range fuzzInt32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x int32) {
		q := hd.Sign(x)

		var v int32
		switch {
		case x > 0:
			v = 1
		case x < 0:
			v = -1
		default:
			v = 0
		}

		if q != v {
			t.Error("x", x, "got", q, "exp", v)
		}
	})
}

func ExampleISIGN() {
	fmt.Println(hd.ISIGN(10, -100000), hd.ISIGN(-10, 100000))
	// Output: -10 10
}

func FuzzISIGN(f *testing.F) {
	for _, x := range fuzzInt32 {
		for _, y := range fuzzInt32 {
			f.Add(x, y)
		}
	}
	f.Fuzz(func(t *testing.T, x, y int32) {
		vs := []int32{
			hd.ISIGN(x, y),
			hd.ISIGN2(x, y),
			hd.ISIGN3(x, y),
			hd.ISIGN4(x, y),
		}
		for i, q := range vs {
			v := x
			if v < 0 {
				v = -v
			}
			if y < 0 {
				v = -v
			}

			if q != v {
				t.Error(i, x, y, "got", q, "exp", v)
			}
		}
	})
}
