package ch2

// ISIGN as called in FORTRAN in four instructions.
func ISIGN(x, y int32) int32 {
	t := y >> 31
	return (Abs(x) ^ t) - t
}

// ISIGNTwo is alternative version
func ISIGNTwo(x, y int32) int32 {
	t := y >> 31
	return (Abs(x) + t) ^ t
}

// ISIGNThree is alternative version
func ISIGNThree(x, y int32) int32 {
	t := (x ^ y) >> 31
	return (x ^ t) - t
}

// ISIGNFour is alternative version
func ISIGNFour(x, y int32) int32 {
	t := (x ^ y) >> 31
	return (x + t) ^ t
}
