package hd

// Abs can be computed in three or four branch-free instructions.
func Abs(x int32) int32 {
	y := x >> 31
	return (x ^ y) - y
}

func Abs2(x int32) int32 {
	y := x >> 31
	return (x + y) ^ y
}

func Abs3(x int32) int32 {
	y := x >> 31
	return x - ((2 * x) & y)
}

func Abs4(x int32) int32 { return DifferenceOrZero(x, 0) + DifferenceOrZero(0, x) }

func NAbs(x int32) int32 {
	y := x >> 31
	return y - (x ^ y)
}

func NAbs2(x int32) int32 {
	y := x >> 31
	return (y - x) ^ y
}

func NAbs3(x int32) int32 {
	y := x >> 31
	return (2 * x & y) - x
}

// AbsFastMul when you have fast instruction for +1/-1 multiplication.
func AbsFastMul(x int32) int32 { return ((x >> 30) | 1) * x }

// AbsDiff is absolute difference that does not overflow.
func AbsDiff[T Integer](x, y T) T { return max(x, y) - min(x, y) }

func AbsDiff2(x, y int32) int32 { return DifferenceOrZero(x, y) + DifferenceOrZero(y, x) }

func AbsDiffUnsigned(x, y uint32) uint32 {
	return DifferenceOrZeroUnsigned(x, y) + DifferenceOrZeroUnsigned(y, x)
}
