package hd

// CountOnes (aka pop) uses divide and conquer to count number set bits.
// Folklore says counting ones is important for NSA, but nobody knows why.
// Applications of this function include Hamming Distance.
func CountOnes(x uint32) uint32 {
	x = (x & 0x5555_5555) + ((x >> 1) & 0x5555_5555)
	x = (x & 0x3333_3333) + ((x >> 2) & 0x3333_3333)
	x = (x & 0x0F0F_0F0F) + ((x >> 4) & 0x0F0F_0F0F)
	x = (x & 0x00FF_00FF) + ((x >> 8) & 0x00FF_00FF)
	x = (x & 0x0000_FFFF) + ((x >> 16) & 0x0000_FFFF)
	return x
}

// CountOnes1 is optimized version of divide and conquer that executes in 21 instructions and is branch-free.
func CountOnes1(x uint32) uint32 {
	x = x - ((x >> 1) & 0x5555_5555)
	x = (x & 0x3333_3333) + ((x >> 2) & 0x3333_3333)
	x = (x + (x >> 4)) & 0x0F0F_0F0F
	x = x + (x >> 8)
	x = x + (x >> 16)
	return x & 0x0000_003F
}

// CountOnes2 is algorithm from HAKMEM (HAK #169) that executes in 13 RISC instructions.
// TODO: why this does not work? unsigned % 63 does not work? uint64 conversions?
/*
func CountOnes2(x uint32) uint32 {
	n, u := uint64(x), uint64(x)
	n = (n >> 1) & 0x333_3333_3333          // Count bits in
	u = u - n                               // each 3-bit
	n = (n >> 1) & 0x333_3333_3333          // field.
	u = u - n                               // each 3-bit
	u = ((u + (u >> 3)) & 0x0307_0707_0707) // 6-bit sums.
	return uint32(u % 63)                   // Add 6-bit sums.
}
*/

// CountOnes3 it executes in 19 RISC instructions, but works well on machines with two addresses.
func CountOnes3(x uint32) uint32 {
	n := x
	n = (x >> 1) & 0x7777_7777       // Count bits in
	x = x - n                        // each 4-bit
	n = (n >> 1) & 0x7777_7777       // field.
	x = x - n                        //
	n = (n >> 1) & 0x7777_7777       //
	x = x - n                        //
	x = (x + (x >> 4)) & 0x0F0F_0F0F // Get byte sums.
	x = x * 0x0101_0101              // Add the bytes.
	return x >> 24
}

// CountOnes4 is very fast for if number of 1 bits is small.
func CountOnes4(x uint32) uint32 {
	var n uint32
	for x != 0 {
		n++
		x &= x - 1
	}
	return n
}

// CountOnes5 uses uint64 registers. It works only with uint15.
func CountOnes5(x uint32) uint32 {
	y := uint64(x)
	y = y * 0x0002_0004_0008_0010
	y = y & 0x1111_1111_1111_1111
	y = y * 0x1111_1111_1111_1111
	y = y >> 60
	return uint32(y)
}

func CompareCountOnes(x, y uint32) int {
	xp := x & ^y
	yp := y &^ x
	for {
		if xp == 0 {
			return int(yp) | -int(yp)
		}
		if yp == 0 {
			return 1
		}
		xp = xp & (xp - 1)
		yp = yp & (yp - 1)
	}
}

// CSA is Carry Save Adder
func CSA(a, b, c uint32) (h, l uint32) {
	u := a ^ b
	v := c
	h = (a & b) | (u & v)
	l = u ^ v
	return h, l
}

// CountOnesArrayGroup8 utilizes CSA for faster count ones computation.
// It is executing with 3x less instructions than direct version with CountOnes per uint32.
func CountOnesArrayGroup8(A []uint32) uint32 {
	var tot uint32
	var ones, twos, twosA, twosB, fours, foursA, foursB, eights uint32

	var i int = 0
	for ; i <= (len(A) - 8); i += 8 {
		twosA, ones = CSA(ones, A[i], A[i+1])
		twosB, ones = CSA(ones, A[i+2], A[i+3])
		foursA, twos = CSA(twos, twosA, twosB)
		twosA, ones = CSA(ones, A[i+4], A[i+5])
		twosB, ones = CSA(ones, A[i+6], A[i+7])
		foursB, twos = CSA(twos, twosA, twosB)
		eights, fours = CSA(fours, foursA, foursB)
		tot = tot + CountOnes(eights)
	}
	tot = (8 * tot) + (4 * CountOnes(fours)) + (2 * CountOnes(twos)) + CountOnes(ones)

	for ; i < len(A); i++ { // Simply add in the last
		tot += CountOnes(A[i]) // 0 to 7 elements.
	}

	return tot
}
