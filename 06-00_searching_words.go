package hd

// ZByteL finds index of left most zero byte.
// This is branch-free code in 11 RISC instructions.
// This is useful for `strlen` in C strings, which use 0 byte for string termination.
func ZByteL(x uint32) int {
	y := (x & 0x7F7F_7F7F) + 0x7F7F_7F7F
	y = ^(y | x | 0x7F7F_7F7F)
	return int(NLZ(y) >> 3)
}

// ZByteL1 is direct algorithm.
func ZByteL1(x uint32) int {
	switch {
	case ((x >> 24) == 0):
		return 0
	case ((x & 0x00FF0000) == 0):
		return 1
	case ((x & 0x0000FF00) == 0):
		return 2
	case ((x & 0x000000FF) == 0):
		return 3
	default:
		return 4
	}
}

func ZByteL64(x uint64) int {
	y := (x & 0x7F7F_7F7F_7F7F_7F7F) + 0x7F7F_7F7F_7F7F_7F7F
	y = ^(y | x | 0x7F7F_7F7F_7F7F_7F7F)
	return NLZ64Basic(y) >> 3
}

// FindInByte m in x illustrates that to find specific byte you need to XOR with repeated value.
func FindInByte(x uint32, m byte) int { return ZByteL(x ^ (uint32(m) * 0x0101_0101)) }

// FindInByteEq finds index of first byte that x and y are equal.
func FindInByteEq(x, y uint32) int { return ZByteL(x ^ y) }

// TODO: search first byte whose value is in range

// FindFirstStringOnes (aka FFStr1) finds index of first string of 1s of length n or 32 if none found.
// This uses divide and conquer and executes in 3 to 36 RISC instructions.
// This loop can be effectively unrolled.
func FindFirstStringOnes(x uint32, n int) int {
	s := 0
	for n > 1 {
		s = n >> 1
		x = x & (x << s)
		n = n - s
	}
	return int(NLZ(x))
}

// FindFirstStringOnes1 is is direct algorithm.
// For worst case this is not good, it is 178 RISC instructions for 0x5555_5555.
func FindFirstStringOnes1(x uint32, n int) int {
	p := 0
	for x != 0 {
		k := int(NLZ(x)) // Skip over initial 0's
		x = x << k       // (if any).
		p += k           //
		k = int(NLZ(^x)) // Count first/next group of 1's.
		if k >= n {      // If enough, return.
			return p
		}
		x = x << k // Not enough 1's, skip over them.
		p += k
	}
	return 32
}

// LenLongestStringOnes (aka maxstr1) returns longest length of string of ones.
// Executes in 131 RISC instructions in worst case.
// TODO: it has optimization to reduce to 39 RISC instructions in worst case.
func LenLongestStringOnes(x uint32) int {
	var k int
	for k = 0; x != 0; k++ {
		x = x & (2 * x)
	}
	return k
}

// LenShortestStringOnes (aka fminstr1) returns shortest length of string of ones.
func LenShortestStringOnes(x uint32) (p, n int) {
	if x == 0 {
		return 32, 0
	}
	b := ^(x >> 1) & x
	e := x & ^(x << 1)
	var k int
	for k = 1; (b & e) == 0; k++ {
		e = e << 1
	}
	return int(NLZ(b & e)), k
}
