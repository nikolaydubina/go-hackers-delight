package hd

// SqrtNewton (isqrt) converges quadratically on series of better approximations.
// This version uses lookup table coded in branches that is used for better starting values.
func SqrtNewton(x uint32) uint32 {
	if x == 0 {
		return 0
	}

	var s int32
	if x <= 4224 {
		if x <= 24 {
			if x <= 3 {
				return (x + 3) >> 2
			} else if x <= 8 {
				return 2
			} else {
				return (x >> 4) + 3
			}
		} else if x <= 288 {
			if x <= 80 {
				s = 3
			} else {
				s = 4
			}
		} else if x <= 1088 {
			s = 5
		} else {
			s = 6
		}
	} else if x <= ((1025 * 1025) - 1) {
		if x <= ((257 * 257) - 1) {
			if x <= ((129 * 129) - 1) {
				s = 8
			} else {
				s = 8
			}
		} else if x <= ((513 * 513) - 1) {
			s = 9
		} else {
			s = 10
		}
	} else if x <= ((4097 * 4097) - 1) {
		if x <= ((2049 * 2049) - 1) {
			s = 11
		} else {
			s = 12
		}
	} else if x <= ((16385 * 16385) - 1) {
		if x <= ((8193 * 8193) - 1) {
			s = 13
		} else {
			s = 14
		}
	} else if x <= ((32769 * 32769) - 1) {
		s = 15
	} else {
		s = 16
	}

	var g0 uint32 = 1 << s
	g1 := (g0 + (x >> s)) >> 1

	for g1 < g0 {
		g0 = g1
		g1 = (g0 + (x / g0)) >> 1
	}

	return g0
}

// SqrtBinarySearch (isqrt) is similar to Newton method with estimate, but performs fully estimation.
func SqrtBinarySearch(x uint32) uint32 {
	var a uint32 = 1
	b := (x >> 5) + 8
	if b > 65535 {
		b = 65535
	}
	for b >= a {
		m := (a + b) >> 1
		if (m * m) > x {
			b = m - 1
		} else {
			a = m + 1
		}
	}
	return a - 1
}

// SqrtShiftAndSubtract (isqrt) is similar to some hardware implementation due to its effective use of shift instructions.
func SqrtShiftAndSubtract(x uint32) uint32 {
	var y uint32
	for m := uint32(0x4000_0000); m != 0; m >>= 2 {
		b := y | m
		y >>= 1
		if x >= b {
			x -= b
			y |= m
		}
	}
	return y
}
