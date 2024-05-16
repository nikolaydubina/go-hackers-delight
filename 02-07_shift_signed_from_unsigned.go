package hd

// ShiftRightSignedFromUnsigned computes signed shift from unsigned shift.
// This compiles to five or six RISC instructions.
// If n is fixed, then this compiles to three or four RISC instructions.
func ShiftRightSignedFromUnsigned(x, n uint32) uint32 {
	return ((x + 0x80000000) >> n) - (0x80000000 >> n)
}

// ShiftRightSignedFromUnsigned2 is alternative version.
// If n is fixed, then this compiles to three or four RISC instructions.
func ShiftRightSignedFromUnsigned2(x, n uint32) uint32 {
	var t uint32 = 0x80000000 >> n
	return ((x >> n) ^ t) - t
}

func ShiftRightSignedFromUnsigned3(x, n uint32) uint32 {
	var t uint32 = (x & 0x80000000) >> n
	return (x >> n) - (t + t)
}

func ShiftRightSignedFromUnsigned4(x, n uint32) uint32 { return (x >> n) | (-(x >> 31) << (31 - n)) }

func ShiftRightSignedFromUnsigned5(x, n uint32) uint32 {
	var t uint32 = -(x >> 31)
	return ((x ^ t) >> n) ^ t
}
