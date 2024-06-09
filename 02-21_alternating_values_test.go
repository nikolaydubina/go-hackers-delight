package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzCycleThreeValues(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b, c int32, n uint8) {
		count := map[int32]int{a: 0, b: 0, c: 0}

		// three different values
		if len(count) != 3 {
			t.Skip()
		}

		if n < 3 {
			n = 3
		}

		x := a
		for i := range int(n) {
			x = hd.CycleThreeValues(a, b, c, x)

			if _, ok := count[x]; !ok {
				t.Errorf("unexpected value %d", x)
			}

			if i >= len(count) && i-count[x] != len(count) {
				t.Errorf("unexpected cycle length %d", i-count[x])
			}

			count[x] = i
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
	fmt.Println(hd.FirstOneOffDifferentBits([3]int32{a, b, c}))
	// Output: [1 0 -1]
}
