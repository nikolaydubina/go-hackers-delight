package hd

// CycleThreeValues is an ingenious and very efficient (8 branch-free instructions) way of cycling in three different constants.
// It requires pre-computing constants that can be done at compile time.
// It relies on the fact that among three different constants there are always two bit positions where they differ and each time one-odd-out.
func CycleThreeValues(a, b, c, x int32) int32 {
	c1, c2, c3, n1, n2 := SetupCycleThreeValues(a, b, c)
	return (((x << (31 - n1)) >> 31) & c1) + (((x << (31 - n2)) >> 31) & c2) + c3
}

// CycleThreeIdentifier constructs identifier byte from bits stored at positions n1 and n2.
func CycleThreeIdentifier(vs [3]int32, n1, n2 uint8) [2][3]int32 {
	return [2][3]int32{iBitAmong(vs, n1), iBitAmong(vs, n2)}
}

// iBitAmong returns i-th bit among three values moved to the least significant bit.
func iBitAmong[T Integer](x [3]T, i uint8) [3]T {
	return [3]T{
		(x[0] >> i) & 1,
		(x[1] >> i) & 1,
		(x[2] >> i) & 1,
	}
}

// FirstOneOffDifferentBits computes 1-based position of first least significant bit
// that is one-off among three values for corresponding value or -1 if there is no such bit.
func FirstOneOffDifferentBits[T int32 | uint32](vs [3]T) [3]int8 {
	n := [3]int8{-1, -1, -1}
	for i := range 31 {
		b := iBitAmong(vs, uint8(i))
		if n[0] == -1 && b[0] != b[1] && b[0] != b[2] {
			n[0] = int8(i)
		}
		if n[1] == -1 && b[1] != b[0] && b[1] != b[2] {
			n[1] = int8(i)
		}
		if n[2] == -1 && b[2] != b[0] && b[2] != b[1] {
			n[2] = int8(i)
		}
	}
	return n
}

// SetupCycleThreeValuesN1N2 rearranges a,b,c in order required for cycling and defines n1 and n2 positions.
func SetupCycleThreeValuesN1N2(a, b, c int32) (na, nb, nc int32, n1, n2 uint8) {
	defer func() {
		if n1 < n2 {
			n1, n2 = n2, n1
			na, nb = nb, na
		}
	}()

	switch n := FirstOneOffDifferentBits([3]int32{a, b, c}); {
	case n[0] == -1 && n[1] >= 0 && n[2] >= 0:
		return b, c, a, uint8(n[1]), uint8(n[2])
	case n[1] == -1 && n[0] >= 0 && n[2] >= 0:
		return a, c, b, uint8(n[0]), uint8(n[2])
	case n[2] == -1 && n[0] >= 0 && n[1] >= 0:
		return a, b, c, uint8(n[0]), uint8(n[1])
	default:
		return a, b, c, uint8(n[0]), uint8(n[1])
	}
}

func cardinality[T comparable](vs []T) int {
	m := make(map[T]bool)
	for _, v := range vs {
		m[v] = true
	}
	return len(m)
}

// SetupCycleThreeValues is final routine that prepares constants for cycling.
// This can be computed at compile time.
func SetupCycleThreeValues(a, b, c int32) (c1, c2, c3 int32, n1, n2 uint8) {
	if cardinality([]int32{a, b, c}) != 3 {
		panic("three values must be different")
	}

	a, b, c, n1, n2 = SetupCycleThreeValuesN1N2(a, b, c)

	switch CycleThreeIdentifier([3]int32{a, b, c}, n1, n2) {
	case [2][3]int32{{0, 1, 1}, {0, 1, 0}}:
		return a - b, c - a, b, n1, n2
	case [2][3]int32{{0, 1, 1}, {1, 0, 1}}:
		return a - b, a - c, b + c - a, n1, n2
	case [2][3]int32{{1, 0, 0}, {0, 1, 0}}:
		return b - a, c - a, a, n1, n2
	case [2][3]int32{{1, 0, 0}, {1, 0, 1}}:
		return b - a, a - c, c, n1, n2
	default:
		panic("unexpected bits positions at n1 and n2")
	}
}
