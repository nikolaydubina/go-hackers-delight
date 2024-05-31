package hd

// ShuffleOuter performs perfect outer shuffle of bits.
// This function is used in cryptography.
// This is 32 RISC instructions.
func ShuffleOuter(x uint32) uint32 {
	x = ((x & 0x0000_FF00) << 8) | ((x >> 8) & 0x0000_FF00) | (x & 0xFF00_00FF)
	x = ((x & 0x00F0_00F0) << 4) | ((x >> 4) & 0x00F0_00F0) | (x & 0xF00F_F00F)
	x = ((x & 0x0C0C_0C0C) << 2) | ((x >> 2) & 0x0C0C_0C0C) | (x & 0xC3C3_C3C3)
	x = ((x & 0x2222_2222) << 1) | ((x >> 1) & 0x2222_2222) | (x & 0x9999_9999)
	return x
}

// UnShuffleOuter is reverse of ShuffleOuter.
func UnShuffleOuter(x uint32) uint32 {
	x = ((x & 0x2222_2222) << 1) | ((x >> 1) & 0x2222_2222) | (x & 0x9999_9999)
	x = ((x & 0x0C0C_0C0C) << 2) | ((x >> 2) & 0x0C0C_0C0C) | (x & 0xC3C3_C3C3)
	x = ((x & 0x00F0_00F0) << 4) | ((x >> 4) & 0x00F0_00F0) | (x & 0xF00F_F00F)
	x = ((x & 0x0000_FF00) << 8) | ((x >> 8) & 0x0000_FF00) | (x & 0xFF00_00FF)
	return x
}

// ShuffleInner performs perfect inner shuffle of bits.
func ShuffleInner(x uint32) uint32 { return ShuffleOuter((x >> 16) | (x << 16)) }

// UnShuffleInner is reverse of ShuffleInner.
func UnShuffleInner(x uint32) uint32 {
	u := UnShuffleOuter(x)
	return (u >> 16) | (u << 16)
}

// TODO: half-shuffle
