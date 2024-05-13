package ch2

// Abs can be computed in three or four branch-free instructions.
func Abs(x int32) int32 {
	y := x >> 31
	return ((x ^ y) - y)
}

func Abs2(x int32) int32 {
	y := x >> 31
	return (x + y) ^ y
}

func Abs3(x int32) int32 {
	y := x >> 31
	return x - ((2 * x) & y)
}

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

// AbsFastMult when you have fast instruction for +1/-1 multiplication.
func AbsFastMult(x int32) int32 { return (x>>30 | 1) * x }
