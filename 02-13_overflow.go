// Overflow means result of operation is too large or too small to fit into the variable.
// Many hardware supply bits for overflow, but high-lever languages may not have access to it (rejected Go [overflow proposal]).
//
// [overflow proposal]: https://github.com/golang/go/issues/31500
package hd

// Unless stated opposite, formulas bellow assume carry to be 0 (version when carry is one does not work; and cannot get carry in Go).

// IsAddOverflow sets most significant bit if overflow occurs.
// This version does not use carry bit and is efficient.
func IsAddOverflow(x, y int32) int32 { return ^(x ^ y) & ((x + y) ^ x) }

func IsAddOverflow2(x, y int32) int32 { return ((x + y) ^ x) & ((x + y) ^ y) }

// IsSubOverflow sets most significant bit if overflow occurs.
// This version does not use carry bit and is efficient.
func IsSubOverflow(x, y int32) int32 { return (x ^ y) & ((x - y) ^ x) }

func IsSubOverflow2(x, y int32) int32 { return ((x - y) ^ x) & (^((x - y) ^ y)) }

func IsAddOverflowUnsigned(x, y uint32) uint32 { return (x & y) | ((x | y) & ^(x + y)) }

func IsAddOverflowUnsigned2(x, y uint32) uint32 { return (x >> 1) + (y >> 1) + ((x & y) & 1) }

func IsAddOverflowUnsigned3(x, y uint32) bool { return ^x < y }

func IsAddOverflowUnsigned4(x, y, sum uint32) bool { return sum < x }

func IsSubOverflowUnsigned(x, y uint32) uint32 { return (^x & y) | (^(x &^ y) & (x - y)) }

func IsSubOverflowUnsigned2(x, y uint32) uint32 { return (x >> 1) - (y >> 1) - ((^x & y) & 1) }

func IsSubOverflowUnsigned3(x, y uint32) bool { return x < y }

func IsSubOverflowUnsigned4(x, y, sub uint32) bool { return sub > x }

func IsMulOverflow(x, y int32) bool {
	// changed to boolean version, since ternary nor "conditional-and" are not supported in Go.
	if y == 0 {
		return false
	}
	c := uint32((^(x ^ y) >> 31)) + (1 << 31)
	return uint32(Abs(x)) > (c / uint32((Abs(y))))
}

// IsDivOverflow uses around seven instructions with branches, and it is not easy to improve this.
func IsDivOverflow(x, y int32) bool { return y == 0 || ((x == (-1 << 31)) && y == -1) }

func IsMulOverflowUnsigned(x, y uint32) bool {
	// changed to boolean version, since ternary nor "conditional-and" are not supported in Go.
	if y == 0 {
		return false
	}
	return x > (0xFFFFFFFF / y)
}

// IsMulOverflowUnsignedNLZ counts number of leading zeros to detect overflow.
func IsMulOverflowUnsignedNLZ(x, y uint32) bool {
	if m, n := LeadingZerosUint32(x), LeadingZerosUint32(y); (m + n) <= 30 {
		return true
	}
	t := x * (y >> 1)
	if int32(t) < 0 {
		return true
	}
	z := t * 2
	if (y & 1) != 0 {
		z += x
		if z < x {
			return true
		}
	}
	return false
}

func IsDivOverflowUnsigned(x, y uint32) bool { return y == 0 }

// TODO: division overflow with long-division instructions
// TODO: with overflow interrupts
// TODO: carry with 2^31 version
