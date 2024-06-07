package hd

// DivSignedPowTwo returns n / 2 ** k for 1 <= k < 31.
// This is illustration, for performance k should be fixed at compile time.
// This is four branch-free instructions.
func DivSignedPowTwo(n int32, k int) int32 {
	t := n >> (k - 1)
	t = ShiftRightUnsigned32(t, (32 - k))
	t += n
	return t >> k
}

func DivSignedPowTwo_fixed5(n int32) int32 {
	const k = 5
	t := n >> (k - 1)
	t = ShiftRightUnsigned32(t, (32 - k))
	t += n
	return t >> k
}

// DivSignedPowTwo2 is alternative version that uses branch.
func DivSignedPowTwo2(n int32, k int) int32 {
	if n >= 0 {
		return n >> k
	}
	return (n + (1 << k) - 1) >> k
}

// DivModSignedPowTwo takes 7 instructions, 5 instructions for division and 2 instructions for reminder.
func DivModSignedPowTwo(n int32, k int) (q, r int32) {
	q = DivSignedPowTwo(n, k)
	return q, n - (q << k)
}

// TODO: dedicated 5 instruction version of just a reminder.

// DivModSignedThree this computes in 6 instructions.
// This is 9 or 10 cycles. Meanwhile, divide operation can take 20 cycles.
func DivModSignedThree(n int32) (q, r int32) {
	const M int32 = 0x5555_5556
	q = MultiplyHighOrderSigned(M, n)
	t := ShiftRightUnsigned32(n, 31)
	q += t
	return q, n - (q * 3)
}

// DivModSignedFive is similar to DivMod3, but error terms is too large, and thus it needs slight variation of magic constant and correction with shift right.
func DivModSignedFive(n int32) (q, r int32) {
	const M int32 = 0x6666_6667
	q = MultiplyHighOrderSigned(M, n)
	q >>= 1
	t := ShiftRightUnsigned32(n, 31)
	q += t
	return q, n - (q * 5)
}

// DivModSignedSeven is similar to DivMod3, but error terms is too large, and thus it needs slight variation of magic constant and correction with shift right.
// The magic is still too large, and thus it needs flipping sign and addition.
func DivModSignedSeven(n int32) (q, r int32) {
	const M int32 = -1840700269
	q = MultiplyHighOrderSigned(M, n)
	q += n
	q >>= 2
	t := ShiftRightUnsigned32(n, 31)
	q += t
	return q, n - (q * 7)
}

// DivModSignedConst performs division by constant.
// This code should be generated at compile time depending on the value of compile time constant k and result of MagicSigned execution.
// The basic trick is to multiply by magic number 2**32/d and then extract leftmost 32 bits of the product.
func DivModSignedConst(n, d int32) (q, r int32) {
	if d < 0 {
		panic("TODO: why integer signed division by negative constants is not working?")
	}

	M, s := DivModSignedConstMagic(d) // compile time

	q = MultiplyHighOrderSigned(M, n)

	// branches bellow generated at compile time conditional on compile constants k, M, s.
	if d > 0 && M < 0 {
		q += n
	}
	if d < 0 && M > 0 {
		q -= n
	}
	if s > 0 {
		q >>= s
	}

	t := ShiftRightUnsigned32(n, 31)
	q += t
	return q, n - (q * d)
}

// DivModSignedConstMagic computes magic number and shift amount for signed integer division.
// This values can be computed at compile time.
// Code using big numbers can be simplified, such as in the case of Python or math.Big, it is no listed here.
func DivModSignedConstMagic(d int32) (M, s int32) {
	if d > -2 && d < 2 {
		panic("d is out of range")
	}

	const two31 uint32 = 0x8000_0000 // 2 ** 31
	var ad uint32 = uint32(Abs(d))

	t := two31 + (uint32(d) >> 31)
	anc := t - 1 - (t % ad) // Absolute value of nc.
	q1 := two31 / anc
	r1 := two31 - (q1 * anc)
	q2 := two31 / ad
	r2 := two31 - (q2 * ad)

	var p int32 = 31
	for {
		p++
		q1 *= 2
		r1 *= 2
		if r1 >= anc {
			q1++
			r1 -= anc
		}

		q2 *= 2
		r2 *= 2
		if r2 >= ad {
			q2++
			r2 -= ad
		}

		if delta := ad - r2; !((q1 < delta) || ((q1 == delta) && (r1 == 0))) {
			break
		}
	}

	M = int32(q2 + 1)
	if d < 0 {
		M = -M
	}

	return M, p - 32
}
