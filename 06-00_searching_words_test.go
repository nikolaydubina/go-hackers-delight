package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleZByteL() {
	fmt.Println(hd.ZByteL(0x12_00_FF_00))
	// Output: 1
}

func FuzzZByteL(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		exp := hd.ZByteL1(x)

		vs := []int{
			hd.ZByteL(x),
		}
		for i, got := range vs {
			if exp != got {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", exp, "got", got)
			}
		}
	})
}

func ExampleFindInByte() {
	fmt.Println(hd.FindInByte(0x12_00_FF_00, 0xFF))
	// Output: 2
}

func ExampleFindInByteEq() {
	fmt.Println(hd.FindInByteEq(0x12_00_F9_00, 0x23_01_F9_00))
	// Output: 2
}

func ExampleFindFirstStringOnes() {
	fmt.Println(hd.FindFirstStringOnes(0b0000_0000_0000_0000_1100_1110_1111_1000, 5))
	// Output: 24
}

func FuzzFindFirstStringOnes(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x, uint8(3))
	}
	f.Fuzz(func(t *testing.T, x uint32, n uint8) {
		n = n % 32
		exp := hd.FindFirstStringOnes(x, int(n))

		vs := []int{
			hd.FindFirstStringOnes1(x, int(n)),
		}
		for i, got := range vs {
			if exp != got {
				t.Error(i, "x", fmt.Sprintf("%032b", x), "exp", exp, "got", got)
			}
		}
	})
}

func ExampleLenLongestStringOnes() {
	fmt.Println(hd.LenLongestStringOnes(0b0000_0000_0000_0000_1100_1110_1111_1000))
	// Output: 5
}

func ExampleLenShortestStringOnes() {
	fmt.Println(hd.LenShortestStringOnes(0xFF0FF0))
	// Output: 8 8
}
