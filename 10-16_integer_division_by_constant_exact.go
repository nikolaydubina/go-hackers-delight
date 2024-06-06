package hd

func DivExactSeven(n int32) int32 { return -0x49249249 * n }

// DivExact uses multiplicative inverse that can be computed at compile time.
func DivExact(n, d int32) int32 {
	var M int32 = MultiplicativeInverseNewtonInt(d)
	return M * n
}

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

func MultiplicativeInverseNewtonInt(d int32) int32 {
	// main algorithm works with by interpreting bits of inputs and outputs as uint32
	return int32(MultiplicativeInverseNewton(uint32(d)))
}

// MultiplicativeInverseNewton uses Newton method of multiplicative inverse.
// It follows from the well-known fact that sequence of xn_1 = xn * (2 - d * xn) converges to 1/d (mod d) given good starting xn.
// This also works with any inverse module of power of 2.
// Each iteration doubles number of correct bits.
func MultiplicativeInverseNewton(d uint32) uint32 {
	if (d % 2) == 0 {
		panic("even number")
	}
	for xn := d; ; xn = xn * (2 - d*xn) {
		if (d * xn) == 1 {
			return xn
		}
	}
}
