package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzExchangeRegistersFull(f *testing.F) {
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

func ExampleExchangeRegistersMasked() {
	var x uint32 = 0xABCDEF12
	var y uint32 = 0x12345678
	var m uint32 = 0x0F0F0F0F
	x, y = hd.ExchangeRegistersMasked(x, y, m)
	fmt.Printf("\n%0X\n%0X", x, y)
	// Output:
	// A2C4E618
	// 1B3D5F72
}

func ExampleExchangeRegistersMasked_bits() {
	var x uint32 = 0b11110000111100001111000011110000
	var y uint32 = 0b00001111000011110000111100001111
	var m uint32 = 0b00001111111100000000000000000000

	x, y = hd.ExchangeRegistersMasked(x, y, m)
	fmt.Printf("\n%032b\n%032b", x, y)
	// Output:
	// 11111111000000001111000011110000
	// 00000000111111110000111100001111
}

func FuzzExchangeRegistersMasked(f *testing.F) {
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
			for _, m := range vs {
				f.Add(x, y, m)
			}
		}
	}
	f.Fuzz(func(t *testing.T, x, y, m int32) {
		ex, ey := hd.ExchangeRegistersMasked(x, y, m)

		if gx, gy := hd.ExchangeRegistersMasked2(x, y, m); ex != gx || ey != gy {
			t.Error(2, x, y, m, gx, gy)
		}

		if gx, gy := hd.ExchangeRegistersMasked3(x, y, m); ex != gx || ey != gy {
			t.Error(3, x, y, m, gx, gy)
		}

		if gx, gy := hd.ExchangeRegistersMasked4(x, y, m); ex != gx || ey != gy {
			t.Error(4, x, y, m, gx, gy)
		}
	})
}

func ExampleExchangeBitsInRegister() {
	var xx uint32 = 0xABCDEF12
	var mb uint32 = 0x0000F000
	var mo uint32 = 0xF0FF0FFF
	fmt.Printf("%0X", hd.ExchangeBitsInRegister(xx, mb, mo, 4*3))
	// Output: AECDBF12
}

func ExampleExchangeBitsInRegisterFast() {
	var xx uint32 = 0xABCDEF12
	var mb uint32 = 0x0000F000
	var mo uint32 = 0xF0FF0FFF
	fmt.Printf("%0X", hd.ExchangeBitsInRegisterFast(xx, mb, mo, 4*3))
	// Output: AECDBF12
}
