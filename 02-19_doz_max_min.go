package hd

// DifferenceOrZero (doz) computes in 7-10 branch-free RISC instructions.
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

// DozRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
// TODO: unsigned version (same code, but requires signed shift right)
func DifferenceOrZeroRanges(x, y int32) int32 { return (x - y) & ^((x - y) >> 31) }

// MaxRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
// TODO: unsigned version (same code, but requires signed shift right)
func MaxRanges(x, y int32) int32 { return x - ((x - y) & ((x - y) >> 31)) }

// MinRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
// TODO: unsigned version (same code, but requires signed shift right)
func MinRanges(x, y int32) int32 { return y + ((x - y) & ((x - y) >> 31)) }

// TODO: expose assembly doz if detected on platform, for very high efficiency
// TODO: version with conditional move instructions
// TODO: version with carry
