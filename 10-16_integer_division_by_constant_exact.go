package hd

// DivModExactSeven uses multiplicative inverse. It expects exact division.
// Magic (5 * 2 ** 32 + 1) / 7 = binary 0xB6DB_6DB7
func DivModExactSeven(n int32) (q, r int32) { return -0x49249249 * n, 0 }

func MultiplicativeInverseEuclidInt(d int32) int32 {
	// main algorithm works with by interpreting bits of inputs and outputs as uint32
	return int32(MultiplicativeInverseEuclid(uint32(d)))
}

// MultiplicativeInverseEuclid uses extended Euclidian algorithm.
func MultiplicativeInverseEuclid(d uint32) uint32 {
	if (d % 2) == 0 {
		panic("even number")
	}
	var x1, x2 uint32 = 0xFFFF_FFFF, 1
	for v1, v2 := -d, d; v2 > 1; {
		q := v1 / v2
		x3 := x1 - (q * x2)
		v3 := v1 - (q * v2)
		x1, x2, v1, v2 = x2, x3, v2, v3
	}
	return x2
}
