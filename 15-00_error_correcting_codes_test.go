package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func TestCheckBits(t *testing.T) {
	cases := []struct {
		u, p uint32
	}{
		{0b0000_0000_0000_0000, 0b00_0000},
	}
	for _, c := range cases {
		if p := hd.CheckBits(c.u); p != c.p {
			t.Errorf("CheckBits(%032b) = %06b, expected %06b", c.u, p, c.p)
		}
	}
}

func TestSyndrome(t *testing.T) {
	cases := []struct {
		up, u, s uint32
	}{
		{up: 0xABCD_EF01, u: 0xABCD_EF01, s: 0},
		{up: 0b0000_0000_0000_0000_0000_0000_0000_0000, u: 0b0000_0000_0000_0000_0000_0000_0000_0000, s: 0},
		{up: 0b0000_0000_0000_0000_0000_0000_0000_0001, u: 0b0000_0000_0000_0000_0000_0000_0000_0000, s: 0b01_1111},
		{up: 0b0000_0000_0000_0000_0000_0000_0000_0010, u: 0b0000_0000_0000_0000_0000_0000_0000_0000, s: 0b10_0001},
		{up: 0b0000_0000_0000_0000_0000_0000_0000_0100, u: 0b0000_0000_0000_0000_0000_0000_0000_0000, s: 0b10_0010},
		{up: 0b1000_0000_0000_0000_0000_0000_0000_0000, u: 0b0000_0000_0000_0000_0000_0000_0000_0000, s: 0b11_1111},
	}
	for _, c := range cases {
		p := hd.CheckBits(c.up)
		if s := hd.Syndrome(p, c.u); s != c.s {
			t.Errorf("Syndrome(%032b, %032b) = %06b, expected %06b", c.up, c.u, s, c.s)
		}
	}
}

func ExampleCheckBits_no_error() {
	v := uint32(0xABCD_EF01)
	p := hd.CheckBits(v)
	u, errb := hd.Correct(p, v)
	fmt.Printf("%0x %v", u, errb)
	// Output: abcdef01 0
}

func FuzzCorrect_ok(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		if y, errb := hd.Correct(hd.CheckBits(x), x); y != x || errb != 0 {
			t.Errorf("Correct(%032b) = %032b, %v", x, y, errb)
		}
	})
}

func FuzzCorrect_error_one_bit(f *testing.F) {
	for _, x := range fuzzUint32 {
		for n := range 32 {
			f.Add(x, uint8(n))
		}
	}
	f.Fuzz(func(t *testing.T, x uint32, n uint8) {
		// flip one bit
		xerr := x ^ (1 << (n % 32))

		p := hd.CheckBits(x)
		xcor, errb := hd.Correct(p, xerr)

		if errb != 1 || xcor != x {
			t.Errorf("Correct(%032b, %032b) = %032b, %v", p, xerr, xcor, errb)
		}
	})
}
