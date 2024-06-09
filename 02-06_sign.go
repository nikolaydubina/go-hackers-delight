package hd

// ExtendSign8 sign-extends a 8-bit number to a 32-bit number.
// Sign of 8-bit number is stored in 8th-bit.
// Sign extension is treating n-th least significant bit as sign bit and copying it to all more significant bits.
// This is usually implemented with shift-left-logical followed by shift-right-arithmetic, but alternative may be useful if you don't have shift.
func ExtendSign7(x uint32) uint32 { return ((x + 0x0000_0080) & 0x0000_00FF) - 0x000_00080 }

// ExtendSign7Two is alternative version.
// If you know all higher order bits are zero, then `and` can be omitted.
func ExtendSign7Two(x uint32) uint32 { return ((x & 0x0000_00ff) ^ 0x0000_0080) - 0x0000_0080 }

// ExtendSign7Three is alternative version
func ExtendSign7Three(x uint32) uint32 { return (x & 0x000_0007f) - (x & 0x0000_0080) }

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

// Sign can be calculated in four RISC instructions with comparison operators.
// Note, there is five and four instruction formula, but in Go it would not fit, since we need explicit logical right shift and conversion of full signed to unsigned bits.
// Note, there is three instruction formula using comparison operators, but in Go strong typing booleans not converted to ints, which would require branches and more instructions.
func Sign(x int32) int32 { return (x >> 31) | ShiftRightUnsigned32(-x, 31) }

func IsMostSignificantSet[T int32 | uint32](x T) bool { return (x >> 31) != 0 }

// ISIGN is sign-transfer function, as known in FORTRAN.
func ISIGN[T Signed](x, y T) T {
	t := y >> 64
	return (Abs(x) ^ t) - t
}

func ISIGN2[T Signed](x, y T) T {
	t := y >> 64
	return (Abs(x) + t) ^ t
}

func ISIGN3[T Signed](x, y T) T {
	t := (x ^ y) >> 64
	return (x ^ t) - t
}

func ISIGN4[T Signed](x, y T) T {
	t := (x ^ y) >> 64
	return (x + t) ^ t
}
