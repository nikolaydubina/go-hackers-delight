package hd

// DivPow2 returns n / 2 ** k for 1 <= k < 31.
// This is illustration, for performance k should be fixed at compile time.
// This is four branch-free instructions.
func DivPow2(n int32, k int) int32 {
	t := n >> (k - 1)                // form the integer
	t = int32(uint32(t) >> (32 - k)) // shift right unsigned is not available directly in Go, so simulating with conversion to uint32, 2 ** k - 1 if n < 0 else 0.
	t += n                           // Add it to n,
	return t >> k                    // and shift right (signed).
}

func DivPow2_fixed5(n int32) int32 {
	const k = 5
	t := n >> (k - 1)
	t = int32(uint32(t) >> (32 - k))
	t += n
	return t >> k
}

// DivPow2Two is alternative version that uses branch.
func DivPow2Two(n int32, k int) int32 {
	if n >= 0 {
		return n >> k
	}
	return (n + (1 << k) - 1) >> k
}

// DivModPow2 takes 7 instructions, 5 instructions for division and 2 instructions for reminder.
func DivModPow2(n int32, k int) (q, r int32) {
	q = DivPow2(n, k)
	return q, n - (q << k)
}

// TODO: dedicated 5 instruction version of just a reminder.

// DivMod3 this computes in 6 instructions.
// This is 9 or 10 cycles. Meanwhile, divide operation can take 20 cycles.
func DivMod3(n int32) (q, r int32) {
	const M int32 = 0x5555_5556       // Magic (2 ** 32 + 2) / 3
	q = MultiplyHighOrderSigned(M, n) // (mulhs), multiply signed, q = floor(M*n/2**32), this can be a single instruction on many architectures.
	t := int32(uint32(n) >> 31)       // shift right unsigned is not available in Go, so simulating with conversion to uint32, Add 1 to q if
	q += t                            // n is negative.
	return q, n - (q * 3)
}

// DivMod5 is similar to DivMod3, but error terms is too large, and thus it needs slight variation of magic constant and correction with shift right.
func DivMod5(n int32) (q, r int32) {
	const M int32 = 0x6666_6667       // Magic (2 ** 33 + 3) / 5
	q = MultiplyHighOrderSigned(M, n) // (mulhs), multiply signed, q = floor(M*n/2**32), this can be a single instruction on many architectures.
	q >>= 1
	t := int32(uint32(n) >> 31) // shift right unsigned is not available in Go, so simulating with conversion to uint32, Add 1 to q if
	q += t                      // n is negative.
	return q, n - (q * 5)
}

// DivMod7 is similar to DivMod3, but error terms is too large, and thus it needs slight variation of magic constant and correction with shift right.
// The magic is still too large, and thus it needs flipping sign and addition.
func DivMod7(n int32) (q, r int32) {
	const M int32 = -1840700269       // Magic (2 ** 34 + 5) / 7 - 2 ** 32. Original text uses 0x92492493, but this gives wrong values and is different from formula.
	q = MultiplyHighOrderSigned(M, n) // (mulhs), multiply signed, q = floor(M*n/2**32), this can be a single instruction on many architectures.
	q += n
	q >>= 2
	t := int32(uint32(n) >> 31) // shift right unsigned is not available in Go, so simulating with conversion to uint32, Add 1 to q if
	q += t                      // n is negative.
	return q, n - (q * 7)
}
