package hd_test

import (
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzMod3Unsigned(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		exp := x % 3

		got := [...]uint32{
			hd.Mod3Unsigned(x),
			hd.Mod3Unsigned2(x),
			hd.Mod3Unsigned3(x),
			hd.Mod3Unsigned4(x),
		}
		for i, q := range got {
			if q != exp {
				t.Errorf("%d: (%d) = %d; want %d", i, x, q, exp)
			}
		}
	})
}
