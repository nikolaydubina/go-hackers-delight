package hd

// ShiftRightSignedFromUnsigned computes signed shift from unsigned shift.
// This compiles to five or six RISC instructions.
// If n is fixed, then this compiles to three or four RISC instructions.
func ShiftRightSignedFromUnsigned(x uint32, n int) uint32 {
	return ((x + 0x8000_0000) >> n) - (0x8000_0000 >> n)
}

// ShiftRightSignedFromUnsigned2 is alternative version.
// If n is fixed, then this compiles to three or four RISC instructions.
func ShiftRightSignedFromUnsigned2(x uint32, n int) uint32 {
	var t uint32 = 0x8000_0000 >> n
	return ((x >> n) ^ t) - t
}

func ShiftRightSignedFromUnsigned3(x uint32, n int) uint32 {
	var t uint32 = (x & 0x8000_0000) >> n
	return (x >> n) - (t + t)
}

func ShiftRightSignedFromUnsigned4(x uint32, n int) uint32 {
	return (x >> n) | (-(x >> 31) << (31 - n))
}

func ShiftRightSignedFromUnsigned5(x uint32, n int) uint32 {
	var t uint32 = -(x >> 31)
	return ((x ^ t) >> n) ^ t
}
