package hd

const u = 99

var nlz_goryavsky = [...]uint8{
	32, 20, 19, u, u, 18, u, 7, 10, 17, u, u, 14, u, 6, u,
	u, 9, u, 16, u, u, 1, 26, u, 13, u, u, 24, 5, u, u,
	u, 21, u, 8, 11, u, 15, u, u, u, u, 2, 27, 0, 25, u,
	22, u, 12, u, u, 3, 28, u, 23, u, 4, 29, u, u, 30, 31,
}

// LeadingZerosUint32 (aka nlz) is algorithm from Robert Harley.
// It consists of 14 instructions branch-free.
// It uses Julius Goryavsky variation for smaller lookup table size.
// LeadingZerosUint32 has direct relationship of log2 as well, and can be used to compute it directly.
// Some instruction sets, such as ARM M1 chips, include single assembly instruction for this operation.
func LeadingZerosUint32(x uint32) uint8 {
	x |= x >> 1 // Propagate leftmost
	x |= x >> 2 // 1-bit to the right
	x |= x >> 4
	x |= x >> 8
	x &= ^(x >> 16)  // Goryavsky
	x *= 0xFD70_49FF // Multiplier is 7 * 255 ** 3, Gorvsky
	return nlz_goryavsky[(x >> 26)]
}

func LeadingZerosUint64(x uint64) uint8 {
	n := LeadingZerosUint32(uint32(x >> 32))
	if (x >> 32) == 0 {
		return n + LeadingZerosUint32(uint32(x))
	}
	return n
}

func LeadingZerosUint32BinarySearch(x uint32) uint8 {
	var y uint32 = 0
	n := 32
	y = x >> 16
	if y != 0 {
		n -= 16
		x = y
	}
	y = x >> 8
	if y != 0 {
		n -= 8
		x = y
	}
	y = x >> 4
	if y != 0 {
		n -= 4
		x = y
	}
	y = x >> 2
	if y != 0 {
		n -= 2
		x = y
	}
	y = x >> 1
	if y != 0 {
		return uint8(n - 2)
	}
	return uint8(n - int(x))
}

func LeadingZerosEqual(x, y uint32) bool { return (x ^ y) <= (x & y) }

func LeadingZerosLess(x, y uint32) bool { return (x & ^y) > y }

func LeadingZerosLessOrEqual(x, y uint32) bool { return (y & ^x) <= x }

// BitSize returns minimum number of bits requires to represent number in two's complement signed number.
func BitSize(x int32) uint8 { return 32 - LeadingZerosUint32((uint32(x ^ (x << 1)))) }

var ntz_reiser = [...]uint8{
	32, 0, 1, 26, 2, 23, 27,
	u, 3, 16, 24, 30, 28, 11, u, 13, 4,
	7, 17, u, 25, 22, 31, 15, 29, 10, 12,
	6, u, 21, 14, 9, 5, 20, 8, 19, 18,
}

// TrailingZerosUint32 (aka ntz) uses John Reiser variant of David Seal method.
// Among applications of TrailingZerosUint32 is R.W.Gosper Loop Detection Algorithm.
func TrailingZerosUint32(x uint32) uint8 { return ntz_reiser[((x & -x) % 37)] }

// LoopDetectionGosper uses R.W.Gosper algorithm to detect start index of a loop and it's period.
// loop is defined on sequence: X_n+1 = f(X_n); X_0, X_1, ..., X_μ-1, X_μ, ... X_μ+λ. [HAK #132].
func LoopDetectionGosper(f func(int) int, x0 int) (μLower, μUpper, λ int) {
	var T [33]int
	T[0] = x0
	Xn := x0
	for n := 1; ; n++ {
		Xn = f(Xn)
		kmax := 31 - LeadingZerosUint32(uint32(n)) // Floor (log2 n)
		for k := uint8(0); k <= kmax; k++ {
			if Xn == T[k] {
				// Compute m = max({i | i < n and ntz(i+1) = k})
				m := ((((n >> k) - 1) | 1) << k) - 1
				λ = n - m
				lgl := 31 - LeadingZerosUint32(uint32(λ-1)) // Ceil(log2 lambda) - 1
				return m - max(1, 1<<lgl) + 1, m, λ
			}
		}
		T[TrailingZerosUint32(uint32(n)+1)] = Xn // No match
	}
}
