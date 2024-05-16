package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleAddFourUint8() {
	x := hd.FourUint8ToUint32([4]uint8{1, 2, 254, 4})
	y := hd.FourUint8ToUint32([4]uint8{2, 3, 1, 5})

	sum := hd.AddFourUint8(x, y)

	fmt.Println(hd.FourUint8FromUint32(sum))
	// Output: [3 5 255 9]
}

func ExampleSubFourUint8() {
	x := hd.FourUint8ToUint32([4]uint8{2, 3, 254, 17})
	y := hd.FourUint8ToUint32([4]uint8{2, 1, 253, 0})

	sum := hd.SubFourUint8(x, y)

	fmt.Println(hd.FourUint8FromUint32(sum))
	// Output: [0 2 1 17]
}

func FuzzFourUint8Arithmetics(f *testing.F) {
	var vs = [][4]uint8{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{8, 0, 1, 8},
		{8, math.MaxUint8, 1, 8},
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x[0], x[1], x[2], x[3], y[0], y[1], y[2], y[3])
		}
	}
	f.Fuzz(func(t *testing.T, xa, xb, xc, xd, ya, yb, yc, yd uint8) {
		vx := [4]uint8{xa, xb, xc, xd}
		vy := [4]uint8{ya, yb, yc, yd}

		add := func(x, y [4]uint8) [4]uint8 { return [4]uint8{x[0] + y[0], x[1] + y[1], x[2] + y[2], x[3] + y[3]} }
		sub := func(x, y [4]uint8) [4]uint8 { return [4]uint8{x[0] - y[0], x[1] - y[1], x[2] - y[2], x[3] - y[3]} }

		x := hd.FourUint8ToUint32(vx)
		y := hd.FourUint8ToUint32(vy)

		v := []struct {
			exp [4]uint8
			got uint32
		}{
			{add(vx, vy), hd.AddFourUint8(x, y)},
			{sub(vx, vy), hd.SubFourUint8(x, y)},
		}
		for i, q := range v {
			if q.exp != hd.FourUint8FromUint32(q.got) {
				t.Error(i, q, x, y)
			}
		}
	})
}

func ExampleAddTwoUint16() {
	x := hd.TwoUint16ToUint32([2]uint16{65535 - 1, 4})
	y := hd.TwoUint16ToUint32([2]uint16{1, 5})

	sum := hd.AddTwoUint16(x, y)

	fmt.Println(hd.TwoUint16FromUint32(sum))
	// Output: [65535 9]
}

func ExampleSubTwoUint16() {
	x := hd.TwoUint16ToUint32([2]uint16{65535 - 1, 5})
	y := hd.TwoUint16ToUint32([2]uint16{65535 - 5, 1})

	sum := hd.SubTwoUint16(x, y)

	fmt.Println(hd.TwoUint16FromUint32(sum))
	// Output: [4 4]
}

func FuzzTwoUint16Arithmetics(f *testing.F) {
	var vs = [][2]uint16{
		{8, 0},
		{0, math.MaxUint8},
	}
	for _, x := range vs {
		for _, y := range vs {
			f.Add(x[0], x[1], y[0], y[1])
		}
	}
	f.Fuzz(func(t *testing.T, xa, xb, ya, yb uint16) {
		vx := [2]uint16{xa, xb}
		vy := [2]uint16{ya, yb}

		add := func(x, y [2]uint16) [2]uint16 { return [2]uint16{x[0] + y[0], x[1] + y[1]} }
		sub := func(x, y [2]uint16) [2]uint16 { return [2]uint16{x[0] - y[0], x[1] - y[1]} }

		x := hd.TwoUint16ToUint32(vx)
		y := hd.TwoUint16ToUint32(vy)

		v := []struct {
			exp [2]uint16
			got uint32
		}{
			{add(vx, vy), hd.AddTwoUint16(x, y)},
			{sub(vx, vy), hd.SubTwoUint16(x, y)},
		}
		for i, q := range v {
			if q.exp != hd.TwoUint16FromUint32(q.got) {
				t.Error(i, q, x, y)
			}
		}
	})
}
