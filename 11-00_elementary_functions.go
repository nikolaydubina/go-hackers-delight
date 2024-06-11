package hd

// Cbrt is cube root, it starts of with shift based Newton algorithm for Sqrt and modifies it to cuber root
func Cbrt(x uint32) uint32 {
	var y uint32
	for s := int32(30); s >= 0; s -= 3 {
		y *= 2
		if b := ((3 * y * (y + 1)) + 1) << s; x >= b {
			x -= b
			y++
		}
	}
	return y
}

// Pow (iexp) takes x to the power of n.
// This utilizes binary expression of exponent.
// For example: x^13, and binary of 13 = 0b1101 = 0b1000 + 0b0100 + 0b0001 = 8 + 4 + 1.
// Thus, we can compute exponent in log number of multiplications.
// However, this is not always optimal code (e.g. n=27 there is better decomposition ((x^3)^3)^3)).
// Note, this function is not resistant to overflows.
func Pow[T Integer](x T, n uint) T {
	var y T = 1
	for p := x; n != 0; p *= p {
		if (n & 1) != 0 {
			y *= p
		}
		n >>= 1
	}
	return y
}

// TODO: pow2(n) in FORTRAN

// Log2 (ilog2)
func Log2(x uint32) uint32 { return 31 - uint32(LeadingZerosUint32(x)) }

var log10_lookup_32 = [...]uint32{0, 9, 99, 999, 9999, 99999, 999999, 9999999, 99999999, 999999999, 0xFFFF_FFFF}

// Log10x32 (ilog10) utilizes multiple techniques, but at its core it is number of leading zeroes with adjustments.
// It is 11 branch free RISC instructions.
func Log10x32(x uint32) uint32 {
	var y int32 = (19 * (31 - int32(LeadingZerosUint32(x)))) >> 6
	y += int32((log10_lookup_32[y+1] - x) >> 31)
	return uint32(y)
}

// Log2x64 (ilog2).
// Interestingly, it is more accurate than float64 based version, since float64 overflows, but this code does not.
func Log2x64(x uint64) uint64 { return 63 - uint64(LeadingZerosUint64(x)) }

var log10_lookup_64 = [...]uint64{
	0, 9, 99, 999, 9999, 99999, 999999, 9999999, 99999999, 999999999,
	9999999999, 99999999999, 999999999999, 9999999999999, 99999999999999,
	999999999999999, 9999999999999999, 99999999999999999, 999999999999999999,
	9999999999999999999,
}

// Log10x64 (ilog10) utilizes multiple techniques, but at its core it is number of leading zeroes with adjustments.
// It is 11 branch free RISC instructions.
func Log10x64(x uint64) uint64 {
	var y int64 = (19 * (63 - int64(LeadingZerosUint64(x)))) >> 6
	y += int64((log10_lookup_64[y+1] - x) >> 63)
	return uint64(y)
}
