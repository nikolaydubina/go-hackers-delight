package hd

const u = 99

var nlz_goryavsky = [...]uint{
	32, 20, 19, u, u, 18, u, 7, 10, 17, u, u, 14, u, 6, u,
	u, 9, u, 16, u, u, 1, 26, u, 13, u, u, 24, 5, u, u,
	u, 21, u, 8, 11, u, 15, u, u, u, u, 2, 27, 0, 25, u,
	22, u, 12, u, u, 3, 28, u, 23, u, 4, 29, u, u, 30, 31,
}

// NLZ is Number of Leading Zeros.
// This is algorithm from Robert Harley.
// It consists of 14 instructions branch-free.
// It uses Julius Goryavsky variation for smaller lookup table size.
// NLZ has direct relationship of log2 as well, and can be used to compute it directly.
// Some instruction sets, such as ARM M1 chips, include single assembly instruction for this operation.
func NLZ(x uint32) uint {
	x = x | (x >> 1) // Propagate leftmost
	x = x | (x >> 2) // 1-bit to the right
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x & ^(x >> 16) // Goryavsky
	x = x * 0xFD7049FF // Multiplier is 7 * 255 ** 3, Gorvsky
	return nlz_goryavsky[(x >> 26)]
}

// NLZBasic is basic algorithm.
func NLZBasic(x uint32) int {
	n := 0
	for i := range 32 {
		if (x & (1 << (31 - i))) != 0 {
			n = i
			break
		}
	}
	return n
}

// NLZ2 uses binary search.
func NLZ2(x uint32) uint {
	var y uint32 = 0
	n := 32

	y = x >> 16
	if y != 0 {
		n = n - 16
		x = y
	}
	y = x >> 8
	if y != 0 {
		n = n - 8
		x = y
	}
	y = x >> 4
	if y != 0 {
		n = n - 4
		x = y
	}
	y = x >> 2
	if y != 0 {
		n = n - 2
		x = y
	}
	y = x >> 1
	if y != 0 {
		return uint(n - 2)
	}
	return uint(n - int(x))
}

// NLZ16Basic is basic algorithm for 16bit numbers.
func NLZ16Basic(x uint16) int {
	for i := range 16 {
		if (x & (1 << (15 - i))) != 0 {
			return i
		}
	}
	return 0
}

// NLZ64Basic is basic algorithm for 64bit numbers.
func NLZ64Basic(x uint64) int {
	for i := range 64 {
		if (x & (1 << (63 - i))) != 0 {
			return i
		}
	}
	return 0
}

func NLZEq(x, y uint32) bool { return (x ^ y) <= (x & y) }

func NLZLess(x, y uint32) bool { return (x & ^y) > y }

func NLZLessEq(x, y uint32) bool { return (y & ^x) <= x }

// BitSize returns minimum number of bits requires to represent number in two's complement signed number.
// This function uses NLZ.
func BitSize(x int32) int { return int(32 - NLZ(uint32(x)^(uint32(x)<<1))) }

var ntz_reiser = [...]int{
	32, 0, 1, 26, 2, 23, 27,
	u, 3, 16, 24, 30, 28, 11, u, 13, 4,
	7, 17, u, 25, 22, 31, 15, 29, 10, 12,
	6, u, 21, 14, 9, 5, 20, 8, 19, 18,
}

// NTZ is Number of Trailing Zeroes.
// This implementation uses John Reiser variant of David Seal method.
// Among applications of NTZ is R.W.Gosper Loop Detection Algorithm.
func NTZ(x uint32) int { return ntz_reiser[((x & -x) % 37)] }

// LoopDetectionGosper uses R.W.Gosper algorithm to detect start index of a loop and it's period.
// loop is defined on sequence: X_n+1 = f(X_n); X_0, X_1, ..., X_μ-1, X_μ, ... X_μ+λ
// This is [HAK #132].
func LoopDetectionGosper(f func(int) int, x0 int) (μLower, μUpper, λ int) {
	T := [33]int{}

	T[0] = x0
	Xn := x0
	for n := 1; ; n++ {
		Xn = f(Xn)
		kmax := 31 - NLZ(uint32(n)) // Floor (log2 n)
		for k := 0; k <= int(kmax); k++ {
			if Xn == T[k] {
				// Compute m = max({i | i < n and ntz(i+1) = k})
				m := ((((n >> k) - 1) | 1) << k) - 1
				λ = n - m
				lgl := 31 - NLZ(uint32(λ-1)) // Ceil(log2 lambda) - 1
				μUpper = m
				μLower = m - max(1, 1<<lgl) + 1
				return μLower, μUpper, λ
			}
		}
		T[NTZ(uint32(n)+1)] = Xn // No match
	}
}
