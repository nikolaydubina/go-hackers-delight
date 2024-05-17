package hd

// DOZ is difference-or-zero operation.
// This computes in severn to ten RISC instructions.
// Note, smaller DOZ version with comparison is skipped as Go lacks comparison returning int.
func DOZ(x, y int32) int32 {
	d := x - y
	return d & (^(d ^ ((x ^ y) & (d ^ x))) >> 31)
}

func Max(x, y int32) int32 { return y + DOZ(x, y) }

func Min(x, y int32) int32 { return x - DOZ(x, y) }

// DOZU is difference-or-zero unsigned operation.
// This computes in severn to ten RISC instructions.
// Note, smaller DOZ version with comparison is skipped as Go lacks comparison returning int.
func DOZU(x, y uint32) uint32 {
	d := x - y
	return d & ^(ShiftRightSignedFromUnsigned(((^x & y) | (^(x ^ y) & d)), 31))
}

// DozRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
// TODO: unsigned version (same code, but requires signed shift right)
func DOZRanges[T int32](x, y T) T { return (x - y) & ^((x - y) >> 31) }

// MaxRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
// TODO: unsigned version (same code, but requires signed shift right)
func MaxRanges[T int32](x, y T) T { return x - ((x - y) & ((x - y) >> 31)) }

// MinRanges requires
// signed x and y be in range [-2^30, 2^30-1]
// and unsigned x and y be in range [0, 2^31-1]
// TODO: unsigned version (same code, but requires signed shift right)
func MinRanges[T int32](x, y T) T { return y + ((x - y) & ((x - y) >> 31)) }

// TODO: expose assembly doz if detected on platform, for very high efficiency
// TODO: version with conditional move instructions
// TODO: version with carry
