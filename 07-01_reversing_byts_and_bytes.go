package hd

// ReverseBits (aka rev) reverses bits in x using divide and conquer.
func ReverseBits(x uint32) uint32 {
	x = ((x & 0x5555_5555) << 1) | ((x & 0xAAAA_AAAA) >> 1)
	x = ((x & 0x3333_3333) << 2) | ((x & 0xCCCC_CCCC) >> 2)
	x = ((x & 0x0F0F_0F0F) << 4) | ((x & 0xF0F0_F0F0) >> 4)
	x = ((x & 0x00FF_00FF) << 8) | ((x & 0xFF00_FF00) >> 8)
	x = ((x & 0x0000_FFFF) << 16) | ((x & 0xFFFF_0000) >> 16)
	return x
}

// shlr15 rotates left 15 bits
func shlr15(x uint32) uint32 { return (x << 15) | (x >> 17) }

// ReverseBits2 uses Knuth algorithm. It is using 21 RISC instructions.
func ReverseBits2(x uint32) uint32 {
	var t uint32
	x = shlr15(x)
	t = (x ^ (x >> 10)) & 0x003F_801F
	x = (t | (t << 10)) ^ x
	t = (x ^ (x >> 4)) & 0x0E03_8421
	x = (t | (t << 4)) ^ x
	t = (x ^ (x >> 2)) & 0x2248_8842
	x = (t | (t << 2)) ^ x
	return x
}

func ReverseBits64Knuth(x uint64) uint64 {
	var t uint64
	x = (x << 31) | (x >> 33) // I.e., shlr(x, 31)
	t = (x ^ (x >> 20)) & 0x0000_0FFF_8000_07FF
	x = (t | (t << 20)) ^ x
	t = (x ^ (x >> 8)) & 0x00F8_000F_8070_0807
	x = (t | (t << 8)) ^ x
	t = (x ^ (x >> 4)) & 0x0808_7090_8080_7008
	x = (t | (t << 4)) ^ x
	t = (x ^ (x >> 2)) & 0x1111_1111_1111_1111
	x = (t | (t << 2)) ^ x
	return x
}

// ReverseBits64Knuth2 .. does not work. TODO: why?
/*
func ReverseBits64Knuth2(x uint64) uint64 {
	var t uint64
	x = (x << 32) | (x >> 32)                                                     // Swap register halves.
	x = ((x & 0x0001_FFFF_0001_FFFF) << 15) | ((x & 0xFFFE_0000_FFFE_0000) >> 17) // Rotate Left 15.
	t = (x ^ (x >> 10)) & 0x003F_801F_003F_801F
	x = (t | (t << 10)) ^ x
	t = (x ^ (x >> 4)) & 0x0E03_8421_0E03_8421
	x = (t | (t << 4)) ^ x
	t = (t ^ (x >> 2)) & 0x2248_8842_2248_8842
	x = (t | (t << 2)) ^ x
	return x
}
*/

// TODO: Generalized bit reversal from [GLS1].

func Reverse8Bit(x uint32) uint32 {
	u := x * 0x0002_0202
	var m uint32 = 0x0104_4010
	s := u & m
	t := (u << 2) & (m << 1)
	return (0x0100_1001) * (s + t) >> 24
}

// IncrReversed value is used in Fast Furier Transform (FFT)
// when computing reversed and incrementing is used in the loop.
// However, computing reversed takes 29 instructions, and lookup table for large values of is not practical.
// Thus storing previous reversed value and incrementing reversed is useful.
// This executes in 5 RISC instructions.
func IncrReversed(x uint32) uint32 {
	s := NLZ(^x)
	return ((x << s) + 0x8000_0000) >> s
}
