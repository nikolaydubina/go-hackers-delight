package hd_test

import (
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func FuzzModUnsigned(f *testing.F) {
	for _, u := range fuzzUint32 {
		f.Add(u)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		got := [...]struct {
			exp uint32
			got uint32
		}{
			{exp: x % 3, got: hd.Mod3Unsigned(x)},
			{exp: x % 3, got: hd.Mod3Unsigned2(x)},
			{exp: x % 3, got: hd.Mod3Unsigned3(x)},
			{exp: x % 3, got: hd.Mod3Unsigned4(x)},
			{exp: x % 5, got: hd.Mod5Unsigned(x)},
			{exp: x % 7, got: hd.Mod7Unsigned(x)},
			{exp: x % 9, got: hd.Mod9Unsigned(x)},
		}
		for i, tc := range got {
			if tc.exp != tc.got {
				t.Errorf("%d: tc(%v) x(%x)", i, tc, x)
			}
		}
	})
}
