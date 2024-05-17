package hd_test

import (
	"fmt"
	"math"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzCycleThreeValues(f *testing.F) {
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

func ExampleCycleThreeValues() {
	var c int32 = 0b10101 // 21
	var a int32 = 0b11111 // 31
	var b int32 = 0b10100 // 20

	out := []int32{a}
	for range 10 {
		out = append(out, hd.CycleThreeValues(a, b, c, out[len(out)-1]))
	}

	fmt.Println(out)
	// Output: [31 20 21 31 20 21 31 20 21 31 20]
}

func TestSetupCycleThreeValuesN1N2(t *testing.T) {
	var c int32 = 0b10101
	var a int32 = 0b11111
	var b int32 = 0b10100

	na, nb, nc, n1, n2 := hd.SetupCycleThreeValuesN1N2(c, a, b)
	if na != 0b11111 || nb != 0b10100 || nc != 0b10101 || n1 != 1 || n2 != 0 {
		t.Errorf("%05b %05b %05b %d %d", na, nb, nc, n1, n2)
	}
}

func ExampleFirstOneOffDifferentBits() {
	var a int32 = 0b11111
	var b int32 = 0b10100
	var c int32 = 0b10101
	fmt.Println(hd.FirstOneOffDifferentBits(a, b, c))
	// Output: 1 0 -1
}
