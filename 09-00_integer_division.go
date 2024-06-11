/*
# Chapter 9: Integer Division

  - Short Division is division of single word by another (e.g. 32bit / 32bit => 32bit). This is typical division operator available in high level languages.
  - Long Division is division of multi-word by single word (e.g. 64bit / 32bit => 32bit).
*/
package hd

// DivModMultiWordUnsigned is Knuth algorithm for integer division.
// It stores quotient in q and remainder in r.
func DivModMultiWordUnsigned(q, r, u, v []uint16) {
	m, n := len(u), len(v)
	const b = 65536 // Number base (16 bits).

	if m < n || n <= 0 || v[n-1] == 0 {
		panic("wrong input")
	}

	if n == 1 {
		var k int32
		// Take care of the case of a single-digit divisor here.
		for j := m - 1; j >= 0; j-- {
			q[j] = uint16((uint32(k)*b + uint32(u[j])) / uint32(v[0]))
			k = int32(k*b+int32(u[j])) - (int32(q[j]) * int32(v[0]))
		}
		if r != nil {
			r[0] = uint16(k)
		}
		return
	}

	// Normalize by shifting v left just enough so that
	// its high-order bit is on, and shift u left the
	// same amount. We may have to append a high-order
	// digit on the dividend; we do that unconditionally.

	s := LeadingZerosUint32(uint32(v[n-1])) - 16 // number of leading zeroes for uint16

	// Normalized form of v
	vn := make([]uint16, 2*n)
	for i := n - 1; i > 0; i-- {
		vn[i] = (v[i] << s) | (v[i-1] >> (16 - s))
	}
	vn[0] = v[0] << s

	// Normalized form of u
	un := make([]uint16, (2 * (m + 1)))
	un[m] = u[m-1] >> (16 - s)
	for i := m - 1; i > 0; i-- {
		un[i] = (u[i] << s) | (u[i-1] >> (16 - s))
	}
	un[0] = u[0] << s

	// Main loop
	for j := m - n; j >= 0; j-- {
		// Compute estimate qhat of q[j].
		// Estimated quotient digit and reminder.
		qhat := (uint32(un[j+n])*b + uint32(un[j+n-1])) / uint32(vn[n-1])
		rhat := (uint32(un[j+n])*b + uint32(un[j+n-1])) - (qhat * uint32(vn[n-1]))

		for qhat >= b || (qhat*uint32(vn[n-2])) > ((rhat*b)+uint32(un[j+n-2])) {
			qhat -= 1
			rhat += uint32(vn[n-1])
			if rhat >= b {
				break
			}
		}

		// Multiply and subtract.
		var t, k int32
		var p uint32 // Product of two digits.
		for i := 0; i < n; i++ {
			p = uint32(qhat * uint32(vn[i]))
			t = int32(un[i+j]) - k - int32(p&0xFFFF)
			un[i+j] = uint16(t)
			k = int32(p>>16) - (t >> 16)
		}
		t = int32(un[j+n]) - k
		un[j+n] = uint16(t)

		// Store quotient digit.
		// If we subtracted too much, add back.
		// This occurs with probability 2/65536 = 0.00003
		q[j] = uint16(qhat)
		if t < 0 {
			q[j] -= 1
			k = 0
			for i := 0; i < n; i++ {
				t = int32(un[i+j]) + int32(vn[i]) + k
				un[i+j] = uint16(t)
				k = t >> 16
			}
			un[j+n] = uint16(int32(un[j+n]) + k)
		}
	}

	// If the caller wants the reminder, unnormalize it and pass it back.
	if r != nil {
		for i := 0; i < n; i++ {
			r[i] = (un[i] >> s) | (un[i+1] << (16 - s))
		}
	}
}

// TODO: unsigned short division from signed short division

// DivModLongUnsigned64b32b (divlu) performs long division of 64-bit unsigned integer by 32-bit unsigned integer.
// This algorithm is slightly modified to store both lower and higher 32 bits of dividend into 64-bit number.
// This algorithm uses shift-and-subtract operations. It illustrates how hardware is doing such division.
// It does not work for overflow cases.
// This executes in 321 to 385 RISC instructions.
func DivModLongUnsigned64b32b(x uint64, y uint32) (q, r uint32) {
	xh := uint32(x >> 32)
	xl := uint32(x)

	for i := 1; i <= 32; i++ {
		t := int32(xh) >> 31 // All 1's if (x31) = 1
		xh = (xh << 1) | (xl >> 31)
		xl <<= 1
		if (xh | uint32(t)) >= y {
			xh -= y
			xl++
		}
	}
	return xl, xh
}

// DivModLongUnsigned64b32b2 is alternative version based on multiword division.
// If overflow, it returns maximum quotient and reminder.
func DivModLongUnsigned64b32b2(x uint64, y uint32) (q, r uint32) {
	u1 := uint32(x >> 32)
	u0 := uint32(x)

	const b = 65536 // Number base (16 bits).

	// Overflow, return maximum quotient and reminder.
	if u1 >= y {
		return 0xFFFF_FFFF, 0xFFFF_FFFF
	}

	s := int32(LeadingZerosUint32(y)) // Shift amount. 0 <= s <= 31
	y <<= s                           // Normalize divisor.

	// Break divisor up into two 16-bit digits. Norm. divisor digits.
	vn1, vn0 := (y >> 16), (y & 0xFFFF)

	// Dividend digit pairs.
	un32 := (u1 << s) | ((u0 >> (32 - s)) & uint32((-s >> 31)))
	un10 := u0 << s // Shift dividend left.

	// Break dividend up into dividend into two digits. Norm. dividend LSD's.
	un1, un0 := (un10 >> 16), (un10 & 0xFFFF)

	// Compute the first quotient digit q1.
	q1 := un32 / vn1
	rhat := un32 - q1*vn1

	for q1 >= b || q1*vn0 > b*rhat+un1 {
		q1--
		rhat += vn1
		if rhat >= b {
			break
		}
	}

	// Multiply and subtract.
	un21 := un32*b + un1 - q1*y

	// Compute the second quotient digit q0.
	q0 := un21 / vn1
	rhat = un21 - q0*vn1

	for q0 >= b || q0*vn0 > b*rhat+un0 {
		q0--
		rhat += vn1
		if rhat >= b {
			break
		}
	}

	r = (un21*b + un0 - q0*y) >> s // Unnormalize the reminder.
	return q1*b + q0, r
}

// TODO: double word division from long division (64 / 64 => 64 from 64 / 32 => 32)
// TODO: signed double word division from unsigned double word division
