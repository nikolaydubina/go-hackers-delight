package hd_test

import "testing"

// De Morgan's Laws can be thought of as distributing the _not_ sign.
func FuzzDeMorganLawsExtended(f *testing.F) {
	f.Fuzz(func(t *testing.T, x, y int32) {
		v := []bool{
			^(x & y) == ^x|^y,
			^(x | y) == ^x & ^y,
			^(x + 1) == ^x-1,
			^(x - 1) == ^x+1,
			^-x == x-1,
			^(x ^ y) == ^x^y, // x === y
			// ^(x === y) == ^x===y == x^y
			^(x + y) == ^x-y,
			^(x - y) == ^x+y,
		}
		for i, q := range v {
			if !q {
				t.Error(i, x, y)
			}
		}
	})
}
