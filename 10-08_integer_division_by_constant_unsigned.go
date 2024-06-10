package hd

func DivModUnsignedPowTwo(n uint32, k int) (q, r uint32) { return n >> k, n - (n >> k << k) }

func DivModUnsignedThree(n uint32) (q, r uint32) {
	const M uint32 = 0xAAAA_AAAB  // Magic (2 ** 33 + 3) / 3
	q = MultiplyHighOrder32(M, n) // mulhu
	q >>= 1
	return q, n - (q * 3)
}

func DivModUnsignedSeven(n uint32) (q, r uint32) {
	const M uint32 = 0x2492_4925  // Magic (2 ** 35 + 3) / 7  - 2 ** 32
	q = MultiplyHighOrder32(M, n) // mulhu
	t := n - q
	t >>= 1
	t += q
	q = t >> 2
	return q, n - (q * 7)
}

func DivModConstUnsigned(n, d uint32) (q, r uint32) {
	if d == 1 {
		return n, 0
	}

	M, a, s := DivModConstUnsignedMagic(d) // compile time
	q = MultiplyHighOrder32(M, n)          // mulhu

	// following branches are computed at compile time based on compile time constants d, M, a, s
	switch a {
	case 0:
		if s > 0 {
			q >>= s
		}
	case 1:
		t := n - q
		t >>= 1
		t += q
		if (s - 1) > 0 {
			q = t >> (s - 1)
		}
	}

	return q, n - (q * d)
}

// DivModConstUnsignedMagic computes magic for unsigned constant multiplication.
// This values can be computed at compile time.
// Code using big numbers can be simplified, such as in the case of Python or math.Big, it is no listed here.
func DivModConstUnsignedMagic(d uint32) (M uint32, a, s int32) {
	var nc uint32 = 0xFFFF_FFFF - ((-d) % d) // Unsigned arithmetic here. -1 changed to 0xFFFF_FFFF for unsigned arithmetic.
	var p int32 = 31

	q1 := 0x8000_0000 / nc
	r1 := 0x8000_0000 - (q1 * nc)
	q2 := 0x7FFF_FFFF / d
	r2 := 0x7FFF_FFFF - (q2 * d)

	for {
		p++
		if r1 >= (nc - r1) {
			q1 = (2 * q1) + 1
			r1 = (2 * r1) - nc
		} else {
			q1 *= 2
			r1 *= 2
		}
		if (r2 + 1) >= (d - r2) {
			if q2 >= 0x7FFF_FFFF {
				a = 1
			}
			q2 = (2 * q2) + 1
			r2 = (2 * r2) + 1 - d
		} else {
			if q2 >= 0x8000_0000 {
				a = 1
			}
			q2 *= 2
			r2 = (2 * r2) + 1
		}
		if delta := d - 1 - r2; !((p < 64) && (q1 < delta || ((q1 == delta) && (r1 == 0)))) {
			break
		}
	}

	return q2 + 1, a, p - 32
}
