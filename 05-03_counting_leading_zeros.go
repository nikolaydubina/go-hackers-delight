package hd

const u = 99

var nlz_goryavsky = [...]uint32{
	32, 20, 19, u, u, 18, u, 7, 10, 17, u, u, 14, u, 6, u,
	u, 9, u, 16, u, u, 1, 26, u, 13, u, u, 24, 5, u, u,
	u, 21, u, 8, 11, u, 15, u, u, u, u, 2, 27, 0, 25, u,
	22, u, 12, u, u, 3, 28, u, 23, u, 4, 29, u, u, 30, 31,
}

// NLZ is Number of Leading Zeros.
// This is algorithm from Robert Harley.
// It consists of 14 instructions branch-free.
// It uses Julius Goryavsky variation for smaller lookup table size.
func NLZ(x uint32) uint32 {
	x = x | (x >> 1) // Propagate leftmost
	x = x | (x >> 2) // 1-bit to the right
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x & ^(x >> 16) // Goryavsky
	x = x * 0xFD7049FF // Multiplier is 7 * 255 ** 3, Gorvsky
	return nlz_goryavsky[(x >> 26)]
}

// NLZ2 uses binary search.
func NLZ2(x uint32) uint32 {
	var y uint32 = 0
	var n uint32 = 32

	y = x >> 16
	if y != 0 {
		n = n - 16
		x = y
	}
	y = x >> 8
	if y != 0 {
		n = n - 8
		x = y
	}
	y = x >> 4
	if y != 0 {
		n = n - 4
		x = y
	}
	y = x >> 2
	if y != 0 {
		n = n - 2
		x = y
	}
	y = x >> 1
	if y != 0 {
		return n - 2
	}
	return n - x
}

func NLZEq(x, y uint32) bool { return (x ^ y) <= (x & y) }

func NLZLess(x, y uint32) bool { return (x & ^y) > y }

func NLZLessEq(x, y uint32) bool { return (y & ^x) <= x }

// BitSize returns minimum number of bits requires to represent number in two's complement signed number.
// This function uses NLZ.
func BitSize(x int32) int { return int(32 - NLZ(uint32(x)^(uint32(x)<<1))) }
