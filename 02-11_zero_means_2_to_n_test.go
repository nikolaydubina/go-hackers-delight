package hd_test

import "testing"

// Encoding number from range [1, 2^n] in lower n bits and treating zero as 2^n
// Here is example of n=3, and three bit fields.
// Encoding is just masking with 2^n - 1
// Decoding without branching is possible. Here are examples that should take three instructions.
func TestZeroMeansTwoThreeBit(t *testing.T) {
	encode := func(x int32) int32 { return x & 7 } // 7 == 0x111
	decode := []func(x int32) int32{
		func(x int32) int32 { return ((x - 1) & 7) + 1 },
		func(x int32) int32 { return ((x + 7) & 7) + 1 },
		func(x int32) int32 { return ((x + 7) | -8) + 9 },
		func(x int32) int32 { return ((x + 7) | 8) - 7 },
		func(x int32) int32 { return 8 - (-x & 7) },
		func(x int32) int32 { return -(-x | -8) },
		func(x int32) int32 { return ((x - 1) & 8) + x },
		func(x int32) int32 { return ((x - 1) | -8) + 9 },
	}

	t.Run("zero means 8", func(t *testing.T) {
		for i, d := range decode {
			if d(0) != 8 {
				t.Error(i)
			}
		}
	})

	t.Run("arithmetics of encoded numbers is equivalent to arithmetics on raw numbers", func(t *testing.T) {
		v := []struct {
			exp int32
			got int32
		}{
			{7, encode(8) - 1},
			{3, encode(2) + 1},
			{2, encode(5) / 2},
			{1, encode(0) + 1},
		}
		for i, q := range v {
			for j, d := range decode {
				if d(q.got) != q.exp {
					t.Error(i, q, j)
				}
			}
		}
	})
}
