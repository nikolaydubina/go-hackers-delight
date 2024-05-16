package hd

func ShiftLeftDoubleLength(x [2]uint32, n int) [2]uint32 {
	// TODO: shift for n > 32, in Go negative shift panics, so need to find workaround
	return [2]uint32{x[0] << n, (x[1] << n) | (x[0] >> (32 - n))}
}

func ShiftRightUnsignedDoubleLength(x [2]uint32, n int) [2]uint32 {
	// TODO: shift for n > 32, in Go negative shift panics, so need to find workaround
	return [2]uint32{(x[0] >> n) | (x[1] << (32 - n)), x[1] >> n}
}

func ShiftRightSignedDoubleLength(x [2]uint32, n int) [2]uint32 {
	// TODO: shift for n > 32, in Go negative shift panics, so need to find workaround
	var y [2]uint32
	y[0] = ((x[0] >> n) | (x[1] << (32 - n)))
	y[1] = ShiftRightSignedFromUnsigned(x[1], n)
	return y
}

// TODO: shift right signed double length for conditional move instructions
