/*
# Chapter 9: Integer Division

  - Short Division is division of single word by another (e.g. 32bit / 32bit => 32bit). This is typical division operator available in high level languages.
  - Long Division is division of multi-word by single word (e.g. 64bit / 32bit => 32bit).
*/
package hd

// DivMultiWordUnsigned is Knuth algorithm for integer division.
// It stores quotient in q and remainder in r.
func DivMultiWordUnsigned(q, r, u, v []uint16) {
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

	s := NLZ16Basic(v[n-1]) // originally it was NLZ() - 16, but algorithm also expects s be in [0,16] thus only NLZ is taken

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

// DivLongUnsigned64b (aka divlu) performs long division of 64-bit unsigned integer by 32-bit unsigned integer.
// This algorithm is slightly modified to store both lower and higher 32 bits of dividend into 64-bit number.
// This algorithm uses shift-and-subtract operations. It illustrates how hardware is doing such division.
// It does not work for overflow cases.
func DivLongUnsigned64b(x uint64, y uint32) (q, r uint32) {
	xh := uint32(x >> 32)
	xl := uint32(x)

	for i := 1; i <= 32; i++ {
		t := int32(xh) >> 31 // All 1's if (x31) = 1
		xh = (xh << 1) | (xl >> 31)
		xl = xl << 1
		if (xh | uint32(t)) >= y {
			xh = xh - y
			xl = xl + 1
		}
	}
	return xl, xh
}
