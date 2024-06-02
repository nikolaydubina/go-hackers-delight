package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleIntToNx16b_small() {
	fmt.Println(hd.IntToNx16b(5))
	// Output: [5]
}

func FuzzIntToNx16b_uint64(f *testing.F) {
	for _, x := range fuzzUint64 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint64) {
		if x != hd.Uint64FromNx16b(hd.IntToNx16b(x)) {
			t.Error("x", x)
		}
	})
}

func FuzzIntToNx16b_int64(f *testing.F) {
	for _, x := range fuzzInt64 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x int64) {
		if got := hd.Int64FromNx16b(hd.IntToNx16b(x)); got != x {
			t.Error("x", x, "got", got)
		}
	})
}
