/*
Certain divisors have repetitive patterns in binary representation which allows to use few shift, add, subtract instructions
to approximate residual and get quotient. It is not trivial how to generate such optimal code.
*/
package hd

// DivShiftUnsigned3 (divu3) uses shifts and adds to approximate through residual. It is 18 instructions.
func DivShiftUnsigned3(n uint32) uint32 {
	q := (n >> 2) + (n >> 4) // q = n * 0.0101 (approx in binary)
	q += q >> 4              // q = n * 0.01010101
	q += q >> 8              //
	q += q >> 16             //
	r := n - (q * 3)         // 0 <= r <= 15
	return q + (11 * r >> 5) // returning q + r/3
}

// DivShiftUnsigned5 (divu5a)
func DivShiftUnsigned5(n uint32) uint32 {
	q := (n >> 3) + (n >> 4)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	r := n - (q * 5)
	return q + ((13 * r) >> 6)
}

// DivShiftUnsigned6 (divu6a)
func DivShiftUnsigned6(n uint32) uint32 {
	q := (n >> 3) + (n >> 5)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	r := n - (q * 6)
	return q + ((11 * r) >> 6)
}

// DivShiftUnsigned7 (divu7)
func DivShiftUnsigned7(n uint32) uint32 {
	q := (n >> 1) + (n >> 4)
	q += q >> 6
	q += (q >> 12) + (q >> 24)
	q >>= 2
	r := n - (7 * q)
	return q + ((r + 1) >> 3)
}

// DivShiftUnsigned9 (div9)
func DivShiftUnsigned9(n uint32) uint32 {
	q := n - (n >> 3)
	q += q >> 6
	q += (q >> 12) + (q >> 24)
	q >>= 3
	r := n - (9 * q)
	return q + ((r + 7) >> 4)
}

// DivShiftUnsigned10 (div10)
func DivShiftUnsigned10(n uint32) uint32 {
	q := (n >> 1) + (n >> 2)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 3
	r := n - (10 * q)
	return q + ((r + 6) >> 4)
}

// DivShiftUnsigned11 (div11)
func DivShiftUnsigned11(n uint32) uint32 {
	q := (n >> 1) + (n >> 2) - (n >> 5) + (n >> 7)
	q += q >> 10
	q += q >> 20
	q >>= 3
	r := n - (11 * q)
	return q + ((r + 5) >> 4)
}

// DivShiftUnsigned12 (div12)
func DivShiftUnsigned12(n uint32) uint32 {
	q := (n >> 1) + (n >> 3)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 3
	r := n - (12 * q)
	return q + ((r + 4) >> 4)
}

// DivShiftUnsigned13 (div13)
func DivShiftUnsigned13(n uint32) uint32 {
	q := (n >> 1) + (n >> 4)
	q += (q >> 4) + (q >> 5)
	q += (q >> 12) + (q >> 24)
	q >>= 3
	r := n - (13 * q)
	return q + ((r + 3) >> 4)
}

// DivShiftUnsigned100 (div100) is 25 instructions.
func DivShiftUnsigned100(n uint32) uint32 {
	q := (n >> 1) + (n >> 3) + (n >> 6) - (n >> 10) + (n >> 12) + (n >> 13) - (n >> 16)
	q += q >> 20
	q >>= 6
	r := n - (100 * q)
	return q + ((r + 28) >> 7)
}

// DivShiftUnsigned1000 (div1000) is 23 instructions.
func DivShiftUnsigned1000(n uint32) uint32 {
	t := (n >> 7) + (n >> 8) + (n >> 12)
	q := (n >> 1) + t + (n >> 15) + (t >> 11) + (t >> 14)
	q >>= 9
	r := n - (1000 * q)
	return q + ((r + 24) >> 10)
}

// DivShiftSigned3 (divs3)
func DivShiftSigned3(n int32) int32 {
	n += (n >> 31) & 2       // Add 2 if n < 0
	q := (n >> 2) + (n >> 4) // q = n * 0.101 (approx in binary)
	q += q >> 4              // q = n * 0.1010101
	q += q >> 8              //
	q += q >> 16             //
	r := n - (q * 3)         // 0 <= r <= 14
	return q + (11 * r >> 5)
}

// DivShiftSigned5 (divs5)
func DivShiftSigned5(n int32) int32 {
	n += (n >> 31) & 4
	q := (n >> 1) + (n >> 2)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 2
	r := n - (q * 5)
	return q + ((7 * r) >> 5)
}

// DivShiftSigned6 (divs6)
func DivShiftSigned6(n int32) int32 {
	n += (n >> 31) & 5
	q := (n >> 1) + (n >> 3)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 2
	r := n - (q * 6)
	return q + ((r + 2) >> 3)
}

// DivShiftSigned7 (divs7)
func DivShiftSigned7(n int32) int32 {
	n += (n >> 31) & 6
	q := (n >> 1) + (n >> 4)
	q += q >> 6
	q += (q >> 12) + (q >> 24)
	q >>= 2
	r := n - (7 * q)
	return q + ((r + 1) >> 3)
}

// DivShiftSigned9 (divs9)
func DivShiftSigned9(n int32) int32 {
	n += (n >> 31) & 8
	q := (n >> 1) + (n >> 2) + (n >> 3)
	q += q >> 6
	q += (q >> 12) + (q >> 24)
	q >>= 3
	r := n - (9 * q)
	return q + ((r + 7) >> 4)
}

// DivShiftSigned10 (divs10)
func DivShiftSigned10(n int32) int32 {
	n += (n >> 31) & 9
	q := (n >> 1) + (n >> 2)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 3
	r := n - (10 * q)
	return q + ((r + 6) >> 4)
}

// DivShiftSigned11 (divs11)
func DivShiftSigned11(n int32) int32 {
	n += (n >> 31) & 10
	q := (n >> 1) + (n >> 2) - (n >> 5) + (n >> 7)
	q += q >> 10
	q += q >> 20
	q >>= 3
	r := n - (11 * q)
	return q + ((r + 5) >> 4)
}

// DivShiftSigned12 (divs12)
func DivShiftSigned12(n int32) int32 {
	n += (n >> 31) & 11
	q := (n >> 1) + (n >> 3)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 3
	r := n - (12 * q)
	return q + ((r + 4) >> 4)
}

// DivShiftSigned13 (divs13)
func DivShiftSigned13(n int32) int32 {
	n += (n >> 31) & 12
	q := (n >> 1) + (n >> 4)
	q += (q >> 4) + (q >> 5)
	q += (q >> 12) + (q >> 24)
	q >>= 3
	r := n - (13 * q)
	return q + ((r + 3) >> 4)
}

// DivShiftSigned100 (divs100)
func DivShiftSigned100(n int32) int32 {
	n += (n >> 31) & 99
	q := (n >> 1) + (n >> 3) + (n >> 6) - (n >> 10) + (n >> 12) + (n >> 13) - (n >> 16)
	q += q >> 20
	q >>= 6
	r := n - (100 * q)
	return q + ((r + 28) >> 7)
}

// DivShiftSigned1000 (divs1000)
func DivShiftSigned1000(n int32) int32 {
	n += (n >> 31) & 999
	t := (n >> 7) + (n >> 8) + (n >> 12)
	q := (n >> 1) + t + (n >> 15) + (t >> 11) + (t >> 14) + (n >> 26) + (t >> 21)
	q >>= 9
	r := n - (1000 * q)
	return q + ((r + 24) >> 10)
}
