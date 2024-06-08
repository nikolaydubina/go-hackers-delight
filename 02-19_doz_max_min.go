package hd

// DifferenceOrZero (doz) returns x - y if x > y else 0.
// It computes in 7-10 branch-free RISC instructions.
// Note, smaller DifferenceOrZero version with comparison is skipped as Go lacks comparison returning int.
func DifferenceOrZero(x, y int32) int32 {
	d := x - y
	return d & (^(d ^ ((x ^ y) & (d ^ x))) >> 31)
}

func Max(x, y int32) int32 { return y + DifferenceOrZero(x, y) }

func Min(x, y int32) int32 { return x - DifferenceOrZero(x, y) }

// DifferenceOrZeroUnsigned (dozu) computes in 7-10 branch-free RISC instructions.
// Note, smaller DOZ version with comparison is skipped as Go lacks comparison returning int.
func DifferenceOrZeroUnsigned(x, y uint32) uint32 {
	d := x - y
	return d & ^(ShiftRightSignedFromUnsigned(((^x & y) | (^(x ^ y) & d)), 31))
}

// DifferenceOrZeroRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
func DifferenceOrZeroRanges[T int32 | uint32](x, y T) T {
	return (x - y) & ^(ShiftRightSigned32((x - y), 31))
}

// MaxRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
func MaxRanges[T int32 | uint32](x, y T) T { return x - ((x - y) & (ShiftRightSigned32((x - y), 31))) }

// MinRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
func MinRanges[T int32 | uint32](x, y T) T { return y + ((x - y) & (ShiftRightSigned32((x - y), 31))) }

// TODO: expose assembly doz if detected on platform, for very high efficiency
// TODO: version with conditional move instructions
// TODO: version with carry
