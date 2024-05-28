package hd

// Parity of number of 1-bits in bit-sequence.
// This executes in 10 instructions.
// This is example of "parallel prefix" operation, that very efficient for parallel computing.
func Parity(x uint32) int {
	y := x ^ (x >> 1)
	y = y ^ (y >> 2)
	y = y ^ (y >> 4)
	y = y ^ (y >> 8)
	y = y ^ (y >> 16)
	return int(y & 1)
}

// Parity2 executes in 9 instructions.
// It avoids computing higher-order parity bits that will not be used.
func Parity2(x uint32) int {
	x = x ^ (x >> 1)
	x = (x ^ (x >> 2)) & 0x1111_1111
	x = x * 0x1111_1111
	return int((x >> 28) & 1)
}

func parity3(x uint32) uint32 { return ((x * 0x1020_4081) & 0x8888_88FF) % 1920 }

// Parity3 computes parity of uint7 and sets 8th bit to 1/0 based on parity.
// Here we wrap into same signature as other Parity methods for convenience.
func Parity3(x uint32) int { return int((parity3(x) >> 7) & 1) }

func parity4(x uint32) uint32 { return ((x * 0x0020_4081) | 0x3DB6_DB00) % 1152 }

// Parity4 computes parity of uint7 and sets 8th bit to 1/0 based on parity, but flipped.
// Here we wrap into same signature as other Parity methods for convenience.
func Parity4(x uint32) int { return int(((parity4(x) >> 7) & 1) ^ 1) }
