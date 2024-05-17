package hd

// CycleThreeValues is an ingenious and very efficient branch-free way of cycling in three constants.
// This requires heavy pre-compute that can be done at compile time.
// This executes in eight instructions.
// This relies on fact that among three different constants there are always two bit positions where they differ and each time one-odd-out.
// Precomputing these values required.
// Note, naive if/ternary code would take four to six instructions with branches for small values.
func CycleThreeValues(a, b, c, x int32) int32 {
	c1, c2, c3, n1, n2 := SetupCycleThreeValues(a, b, c) // WARNING: compute this at compile time
	return (((x << (31 - n1)) >> 31) & c1) + (((x << (31 - n2)) >> 31) & c2) + c3
}

// CycleThreeIdentifier constructs identifier byte from bits stored at positions n1 and n2.
// If all these are constants, this can be computed at compile time.
// This can be computed at compile time.
func CycleThreeIdentifier(a, b, c int32, n1, n2 int) [2][3]int32 {
	bn1 := [3]int32{(a >> n1) & 1, (b >> n1) & 1, (c >> n1) & 1}
	bn2 := [3]int32{(a >> n2) & 1, (b >> n2) & 1, (c >> n2) & 1}
	return [2][3]int32{bn1, bn2}

}

// FirstOneOffDifferentBits computes 1-based position of first least significant bit that is one-off among three values for corresponding value.
// This can be computed at compile time.
func FirstOneOffDifferentBits(a, b, c int32) (na, nb, nc int) {
	na, nb, nc = -1, -1, -1
	for i := range 31 {
		// TODO: negative values? right binary shift? uint32?
		ba := (a >> i) & 1
		bb := (b >> i) & 1
		bc := (c >> i) & 1
		if na == -1 && ba != bb && ba != bc {
			na = i
		}
		if nb == -1 && bb != ba && bb != bc {
			nb = i
		}
		if nc == -1 && bc != ba && bc != bb {
			nc = i
		}
	}
	return na, nb, nc
}

// SetupCycleThreeValuesN1N2 rearranges a,b,c in order required for cycling and defines n1 and n2 positions.
// This can be computed at compile time.
func SetupCycleThreeValuesN1N2(a, b, c int32) (na, nb, nc int32, n1, n2 int) {
	defer func() {
		if n1 < n2 {
			n1, n2 = n2, n1
			na, nb = nb, na
		}
	}()

	n1, n2, n3 := FirstOneOffDifferentBits(a, b, c)
	if n1 == -1 {
		return b, c, a, n2, n3
	}
	if n2 == -1 {
		return a, c, b, n1, n3
	}
	return a, b, c, n1, n2
}

// SetupCycleThreeValues is final routine that prepares constants for cycling.
// This can be computed at compile time.
func SetupCycleThreeValues(a, b, c int32) (c1, c2, c3 int32, n1, n2 int) {
	a, b, c, n1, n2 = SetupCycleThreeValuesN1N2(a, b, c)

	switch CycleThreeIdentifier(a, b, c, n1, n2) {
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
