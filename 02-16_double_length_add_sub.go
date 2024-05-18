package hd

// DoubleLengthInt32FromUint64 unpacks two uint32 from single uint64.
// This is known as Double-Length arithmetics.
// This can be implemented in five instructions by using only 31 bits and storing carry in most significant bit.
// Here is only simplified versions.
func DoubleLengthInt32FromUint64(x uint64) [2]uint32 { return [2]uint32{uint32(x), uint32(x >> 32)} }

// AddDoubleLength of uint64 numbers encoded in two uint32.
// This takes nine instructions.
func AddDoubleLength(x [2]uint32, y [2]uint32) [2]uint32 {
	var z [2]uint32
	z[0] = x[0] + y[0]
	c := ((x[0] & y[0]) | ((x[0] | y[0]) & ^z[0])) >> 31
	z[1] = x[1] + y[1] + c
	return z
}

// SubDoubleLength of uint64 numbers encoded in two uint32.
// This takes eight instructions.
func SubDoubleLength(x [2]uint32, y [2]uint32) [2]uint32 {
	var z [2]uint32
	z[0] = x[0] - y[0]
	b := ((^x[0] & y[0]) | (^(x[0] ^ y[0]) & z[0])) >> 31
	z[1] = x[1] - y[1] - b
	return z
}
