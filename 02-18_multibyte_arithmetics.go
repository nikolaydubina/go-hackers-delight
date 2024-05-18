package hd

// FourUint8FromUint32 unpacks 4 uint8 from single uint32.
func FourUint8FromUint32(x uint32) [4]uint8 {
	return [4]uint8{uint8(x), uint8(x >> 8), uint8(x >> 16), uint8(x >> 24)}
}

// FourUint8ToUint32  packs 4 uint4 into single uint32.
func FourUint8ToUint32(x [4]uint8) uint32 {
	return uint32(x[0]) | (uint32(x[1]) << 8) | (uint32(x[2]) << 16) | (uint32(x[3]) << 24)
}

// AddFourUint8 executes in eight branch-free instructions
func AddFourUint8(x, y uint32) uint32 {
	s := (x & 0x7F7F7F7F) + (y & 0x7F7F7F7F) // mask and add, no carries
	return ((x ^ y) & 0x80808080) ^ s        // fix higher order bits and carry into that bit
}

// SubFourUint8 executes in eight branch-free instructions
func SubFourUint8(x, y uint32) uint32 {
	d := (x | 0x80808080) - (y & 0x7F7F7F7F) // mask and subtract, no borrows
	return ^(((x ^ y) | 0x7F7F7F7F) ^ d)     // fix higher order bits and borrow into that bit
}

// TwoUint16FromUint32 unpacks two uint16 from single uint32.
func TwoUint16FromUint32(x uint32) [2]uint16 { return [2]uint16{uint16(x), uint16(x >> 16)} }

// TwoUint16ToUint32 packs two uint16 into single uint32.
func TwoUint16ToUint32(x [2]uint16) uint32 { return uint32(x[0]) | (uint32(x[1]) << 16) }

// AddFourUint8 executes in seven branch-free instructions
func AddTwoUint16(x, y uint32) uint32 {
	s := x + y
	c := (s ^ x ^ y) & 0x00010000
	return s - c
}

// SubFourUint8 executes in seven branch-free instructions
func SubTwoUint16(x, y uint32) uint32 {
	d := x - y
	b := (d ^ x ^ y) & 0x000010000
	return d + b
}

// TODO: container for multiple signed ints, preliminary work shows it does not work
// TODO: abs function
