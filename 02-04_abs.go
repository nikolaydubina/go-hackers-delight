package hd

// Abs can be computed in three or four branch-free instructions.
func Abs[T Signed](x T) T {
	// in Go shift has no upper limit, with this we can make function work for any signed type.
	// compiler will generate assembly for maximum amount of shift right for a given type.
	y := x >> 64
	return (x ^ y) - y
}

func Abs2[T Signed](x T) T {
	y := x >> 64
	return (x + y) ^ y
}

func Abs3[T Signed](x T) T {
	y := x >> 64
	return x - ((2 * x) & y)
}

func Abs4[T Signed](x T) T { return DifferenceOrZero(x, 0) + DifferenceOrZero(0, x) }

func NAbs[T Signed](x T) T {
	y := x >> 64
	return y - (x ^ y)
}

func NAbs2[T Signed](x T) T {
	y := x >> 64
	return (y - x) ^ y
}

func NAbs3[T Signed](x T) T {
	y := x >> 64
	return (2 * x & y) - x
}

// AbsFastMul when you have fast instruction for +1/-1 multiplication.
func AbsFastMul(x int32) int32 { return ((x >> 30) | 1) * x }

// AbsDiff is absolute difference that does not overflow.
func AbsDiff[T Integer](x, y T) T { return max(x, y) - min(x, y) }

func AbsDiff2[T Signed](x, y T) T { return DifferenceOrZero(x, y) + DifferenceOrZero(y, x) }

func AbsDiffUnsigned(x, y uint32) uint32 {
	return DifferenceOrZeroUnsigned(x, y) + DifferenceOrZeroUnsigned(y, x)
}
