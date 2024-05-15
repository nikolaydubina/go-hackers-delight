package ch2

// ISIGN is sign-transfer function, as known in FORTRAN.
func ISIGN(x, y int32) int32 {
	t := y >> 31
	return (Abs(x) ^ t) - t
}

func ISIGN2(x, y int32) int32 {
	t := y >> 31
	return (Abs(x) + t) ^ t
}

func ISIGN3(x, y int32) int32 {
	t := (x ^ y) >> 31
	return (x ^ t) - t
}

func ISIGN4(x, y int32) int32 {
	t := (x ^ y) >> 31
	return (x + t) ^ t
}
