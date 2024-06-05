package hd

import (
	"fmt"
	"sync"
)

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
	t := int32(uint32(n) >> 31) // shift right unsigned is not available in Go, so simulating with conversion to uint32, Add 1 to q if negative
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
	t := int32(uint32(n) >> 31) // shift right unsigned is not available in Go, so simulating with conversion to uint32, Add 1 to q if negative
	q += t                      // n is negative.
	return q, n - (q * 7)
}

// DivModSignedConst performs division by constant.
// This code should be generated at compile time depending on the value of compile time constant k and result of MagicSigned execution.
func DivModSignedConst(n, d int32) (q, r int32) {
	M, s := MulSignedMagicCached(d) // compile time

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

	t := int32(uint32(n) >> 31) // shift right unsigned is not available in Go, so simulating with conversion to uint32, Add 1 to q if negative
	q += t                      // n is negative.
	return q, n - (q * d)
}

type magicSignedMul struct {
	M int32
	s int32
}

var mulSignedMagicCache = make(map[int32]magicSignedMul)
var mulSignedMagicCacheMtx = sync.Mutex{}

func MulSignedMagicCached(d int32) (M, s int32) {
	mulSignedMagicCacheMtx.Lock()
	defer mulSignedMagicCacheMtx.Unlock()

	if _, ok := mulSignedMagicCache[d]; !ok {
		M, s = MulSignedMagic(d)
		mulSignedMagicCache[d] = magicSignedMul{M: M, s: s}
	}
	return mulSignedMagicCache[d].M, mulSignedMagicCache[d].s
}

// MulSignedMagic computes magic number and shift amount for signed integer division.
// This values can be computed at compile time.
// Code using big numbers can be simplified, such as in the case of Python or math.Big, it is no listed here.
func MulSignedMagic(d int32) (M, s int32) {
	if d > -2 && d < 2 { // Original condition was testing against boundaries of int32, which is not necessary.
		panic(fmt.Errorf("d(%v) is out of range", d))
	}

	const two31 uint32 = 0x8000_0000 // 2 ** 31
	var ad uint32 = uint32(Abs(d))

	t := two31 + (uint32(d) >> 31)
	anc := t - 1 - (t % ad)  // Absolute value of nc.
	q1 := two31 / anc        // Init q1 = 2 ** p / |nc|.
	r1 := two31 - (q1 * anc) // Init r1 = rem(2 ** p, |nc|).
	q2 := two31 / ad         // Init q2 = 2 ** p / |d|.
	r2 := two31 - (q2 * ad)  // Init r2 = rem(2 ** p, |d|).

	var p int32 = 31
	for {
		p++
		q1 *= 2 // Update q1 = 2 ** p / |nc|.
		r1 *= 2 // Update r1 = rem(2 ** p, |nc|).
		if r1 >= anc {
			q1++
			r1 -= anc
		}

		q2 *= 2 // Update q2 = 2 ** p / |d|.
		r2 *= 2 // Update r2 = rem(2 ** p, |d|).
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
