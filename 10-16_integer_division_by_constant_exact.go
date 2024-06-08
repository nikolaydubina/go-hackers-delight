package hd

func DivExactSeven(n int32) int32 { return n * -0x49249249 }

// DivExact uses multiplicative inverse that can be computed at compile time.
// It relies on theorem "if a and m are relatively prime integers, then there exists an integer ā such that a*ā = 1 (mod m)".
func DivExact[T Integer](n, d T) T { return n * MultiplicativeInverseNewton(d) }

// IsDivExactUnsignedOdd tests if n is multiple of d, for odd d. It has single branch.
func IsDivExactUnsignedOdd(n, d uint32) bool {
	M := MultiplicativeInverseNewton(d)
	q := M * n
	c := 0xFFFF_FFFF / d
	return q <= c
}

// IsDivExactUnsigned is similar to odd version but transforms divisor d0 * 2^k and performs rotate right trick. It has single branch.
func IsDivExactUnsigned(n, d uint32) bool {
	k := TrailingZerosUint32(d)
	M := MultiplicativeInverseNewton(d >> k)
	q := M * n
	q = RotateRight(q, int(k))
	c := 0xFFFF_FFFF / d
	return q <= c
}

// IsDivExactSigned is simplified version with abs. it has single branch. however, it is not minimal number of instructions.
func IsDivExactSigned(n, d int32) bool { return IsDivExactUnsigned(uint32(Abs(n)), uint32(Abs(d))) }

// TODO: IsDivExactSigned does not work, has something to do with signed and unsigned conversion.

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

// MultiplicativeInverseNewton uses Newton method of multiplicative inverse.
// It follows from the well-known fact that sequence of xn_1 = xn * (2 - d * xn) converges to 1/d (mod d) given good starting xn.
// This also works with any inverse module of power of 2.
// Each iteration doubles number of correct bits.
func MultiplicativeInverseNewton[T Integer](d T) T {
	if (d % 2) == 0 {
		panic("even number")
	}
	for xn := d; ; xn *= 2 - (d * xn) {
		if (d * xn) == 1 {
			return xn
		}
	}
}
