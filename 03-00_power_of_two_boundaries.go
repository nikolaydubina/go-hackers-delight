package hd

func RoundDownBlockPowerOfTwo[T Unsigned](x T, k uint8) T { return x & (-T(1) << k) }

func RoundDownBlockPowerOfTwo2[T Unsigned](x T, k uint8) T { return (x >> k) << k }

func RoundUpBlockPowerOfTwo[T Unsigned](x T, k uint8) T {
	var t T = (1 << k) - 1 // if k is const, this is const too
	return (x + t) & ^t
}

func RoundUpBlockPowerOfTwo2[T Unsigned](x T, k uint8) T {
	var t T = -T(1) << k // if k is const, this is const too
	return (x - t - 1) & t
}

// FLPTwo is Floor to nearest Power of Two.
// Values >= 2^31 will be zero.
// This formula works for x > 0.
func FLPTwo(x uint32) uint32 { return 1 << (31 - LeadingZerosUint32(x)) }

// FLPTwo2 is alternative version that works for x == 0 as well.
func FLPTwo2(x uint32) uint32 { return 1 << (LeadingZerosUint32(x) ^ 31) }

func FLPTwo3(x uint32) uint32 { return 0x80000000 >> LeadingZerosUint32(x) }

// FLPTwo4 is alternative branch-free version when NLZ is not available.
func FLPTwo4(x uint32) uint32 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x - (x >> 1)
}

// FLPTwo5 uses simple loop that executes in 4 * NLZ(x) + 3 instructions.
func FLPTwo5(x uint32) uint32 {
	var y uint32 = 0x80000000
	for y > x {
		y >>= 1
	}
	return y
}

// CLPTwo is Ceil to nearest Power of Two.
// Values >= 2^31 will be zero.
func CLPTwo(x uint32) uint32 { return 1 << (32 - LeadingZerosUint32(x-1)) }

func CLPTwo2(x uint32) uint32 { return 0x80000000 >> (LeadingZerosUint32(x-1) - 1) }

// CLPTwo3 is alternative branch-free version when NLZ is not available.
func CLPTwo3(x uint32) uint32 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x + 1
}

// IsPowerOfTwoBoundaryCrossed checks if adding l to a crosses power of block of size b.
// b size has to be power of two.
// This and versions bellow are five or six RISC instructions.
// b has to be power of two and likely to be a constant.
func IsPowerOfTwoBoundaryCrossed[T Unsigned](a, l, b T) bool { return -(a | -b) < l }

func IsPowerOfTwoBoundaryCrossed2[T Unsigned](a, l, b T) bool { return (^(a | -b) + 1) < l }

func IsPowerOfTwoBoundaryCrossed3[T Unsigned](a, l, b T) bool { return ((^a & (b - 1)) + 1) < l }

func IsPowerOfTwoBoundaryCrossed4[T Unsigned](a, l, b T) bool { return (b - (a & (b - 1))) < l }

// TODO: round power of two signed
