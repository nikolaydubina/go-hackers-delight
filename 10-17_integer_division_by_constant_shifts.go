/*
Certain divisors have repetitive patterns in binary representation which allows to use few (~20) shift, add, subtract instructions
to approximate residual and get quotient. It is not trivial how to generate such optimal code.
*/
package hd

// Div3ShiftUnsigned (divu3) uses shifts and adds to approximate through residual. It is 18 instructions.
func Div3ShiftUnsigned(n uint32) uint32 {
	q := (n >> 2) + (n >> 4) // q = n * 0.0101 (approx in binary)
	q += q >> 4              // q = n * 0.01010101
	q += q >> 8              //
	q += q >> 16             //
	r := n - (q * 3)         // 0 <= r <= 15
	return q + (11 * r >> 5) // returning q + r/3
}

// Div5ShiftUnsigned (divu5a)
func Div5ShiftUnsigned(n uint32) uint32 {
	q := (n >> 3) + (n >> 4)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	r := n - (q * 5)
	return q + ((13 * r) >> 6)
}

// Div6ShiftUnsigned (divu6a)
func Div6ShiftUnsigned(n uint32) uint32 {
	q := (n >> 3) + (n >> 5)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	r := n - (q * 6)
	return q + ((11 * r) >> 6)
}

// Div7ShiftUnsigned (divu7)
func Div7ShiftUnsigned(n uint32) uint32 {
	q := (n >> 1) + (n >> 4)
	q += q >> 6
	q += (q >> 12) + (q >> 24)
	q >>= 2
	r := n - (7 * q)
	return q + ((r + 1) >> 3)
}

// Div9ShiftUnsigned (div9)
func Div9ShiftUnsigned(n uint32) uint32 {
	q := n - (n >> 3)
	q += q >> 6
	q += (q >> 12) + (q >> 24)
	q >>= 3
	r := n - (9 * q)
	return q + ((r + 7) >> 4)
}

// Div10ShiftUnsigned (div10)
func Div10ShiftUnsigned(n uint32) uint32 {
	q := (n >> 1) + (n >> 2)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 3
	r := n - (10 * q)
	return q + ((r + 6) >> 4)
}

// Div11ShiftUnsigned (div11)
func Div11ShiftUnsigned(n uint32) uint32 {
	q := (n >> 1) + (n >> 2) - (n >> 5) + (n >> 7)
	q += q >> 10
	q += q >> 20
	q >>= 3
	r := n - (11 * q)
	return q + ((r + 5) >> 4)
}

// Div12ShiftUnsigned (div12)
func Div12ShiftUnsigned(n uint32) uint32 {
	q := (n >> 1) + (n >> 3)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 3
	r := n - (12 * q)
	return q + ((r + 4) >> 4)
}

// Div13ShiftUnsigned (div13)
func Div13ShiftUnsigned(n uint32) uint32 {
	q := (n >> 1) + (n >> 4)
	q += (q >> 4) + (q >> 5)
	q += (q >> 12) + (q >> 24)
	q >>= 3
	r := n - (13 * q)
	return q + ((r + 3) >> 4)
}

// Div100ShiftUnsigned (div100) is 25 instructions.
func Div100ShiftUnsigned(n uint32) uint32 {
	q := (n >> 1) + (n >> 3) + (n >> 6) - (n >> 10) + (n >> 12) + (n >> 13) - (n >> 16)
	q += q >> 20
	q >>= 6
	r := n - (100 * q)
	return q + ((r + 28) >> 7)
}

// Div1000ShiftUnsigned (div1000) is 23 instructions.
func Div1000ShiftUnsigned(n uint32) uint32 {
	t := (n >> 7) + (n >> 8) + (n >> 12)
	q := (n >> 1) + t + (n >> 15) + (t >> 11) + (t >> 14)
	q >>= 9
	r := n - (1000 * q)
	return q + ((r + 24) >> 10)
}

// Div3ShiftSigned (divs3)
func Div3ShiftSigned(n int32) int32 {
	n += (n >> 31) & 2       // Add 2 if n < 0
	q := (n >> 2) + (n >> 4) // q = n * 0.101 (approx in binary)
	q += q >> 4              // q = n * 0.1010101
	q += q >> 8              //
	q += q >> 16             //
	r := n - (q * 3)         // 0 <= r <= 14
	return q + (11 * r >> 5)
}

// Div5ShiftSigned (divs5)
func Div5ShiftSigned(n int32) int32 {
	n += (n >> 31) & 4
	q := (n >> 1) + (n >> 2)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 2
	r := n - (q * 5)
	return q + ((7 * r) >> 5)
}

// Div6ShiftSigned (divs6)
func Div6ShiftSigned(n int32) int32 {
	n += (n >> 31) & 5
	q := (n >> 1) + (n >> 3)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 2
	r := n - (q * 6)
	return q + ((r + 2) >> 3)
}

// Div7ShiftSigned (divs7)
func Div7ShiftSigned(n int32) int32 {
	n += (n >> 31) & 6
	q := (n >> 1) + (n >> 4)
	q += q >> 6
	q += (q >> 12) + (q >> 24)
	q >>= 2
	r := n - (7 * q)
	return q + ((r + 1) >> 3)
}

// Div9ShiftSigned (divs9)
func Div9ShiftSigned(n int32) int32 {
	n += (n >> 31) & 8
	q := (n >> 1) + (n >> 2) + (n >> 3)
	q += q >> 6
	q += (q >> 12) + (q >> 24)
	q >>= 3
	r := n - (9 * q)
	return q + ((r + 7) >> 4)
}

// Div10ShiftSigned (divs10)
func Div10ShiftSigned(n int32) int32 {
	n += (n >> 31) & 9
	q := (n >> 1) + (n >> 2)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 3
	r := n - (10 * q)
	return q + ((r + 6) >> 4)
}

// Div11ShiftSigned (divs11)
func Div11ShiftSigned(n int32) int32 {
	n += (n >> 31) & 10
	q := (n >> 1) + (n >> 2) - (n >> 5) + (n >> 7)
	q += q >> 10
	q += q >> 20
	q >>= 3
	r := n - (11 * q)
	return q + ((r + 5) >> 4)
}

// Div12ShiftSigned (divs12)
func Div12ShiftSigned(n int32) int32 {
	n += (n >> 31) & 11
	q := (n >> 1) + (n >> 3)
	q += q >> 4
	q += q >> 8
	q += q >> 16
	q >>= 3
	r := n - (12 * q)
	return q + ((r + 4) >> 4)
}

// Div13ShiftSigned (divs13)
func Div13ShiftSigned(n int32) int32 {
	n += (n >> 31) & 12
	q := (n >> 1) + (n >> 4)
	q += (q >> 4) + (q >> 5)
	q += (q >> 12) + (q >> 24)
	q >>= 3
	r := n - (13 * q)
	return q + ((r + 3) >> 4)
}

// Div100ShiftSigned (divs100)
func Div100ShiftSigned(n int32) int32 {
	n += (n >> 31) & 99
	q := (n >> 1) + (n >> 3) + (n >> 6) - (n >> 10) + (n >> 12) + (n >> 13) - (n >> 16)
	q += q >> 20
	q >>= 6
	r := n - (100 * q)
	return q + ((r + 28) >> 7)
}

// Div1000ShiftSigned (divs1000)
func Div1000ShiftSigned(n int32) int32 {
	n += (n >> 31) & 999
	t := (n >> 7) + (n >> 8) + (n >> 12)
	q := (n >> 1) + t + (n >> 15) + (t >> 11) + (t >> 14) + (n >> 26) + (t >> 21)
	q >>= 9
	r := n - (1000 * q)
	return q + ((r + 24) >> 10)
}
