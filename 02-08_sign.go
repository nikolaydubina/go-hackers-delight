package hd

// Sign can be calculated in four RISC instructions with comparison operators.
// We are converting to uint32 to get logical right shift, since there is no int32 logical right shift in Go. This can cause generating more instructions.
// Note, there is five and four instruction formula, but in Go it would not fit, since we need explicit logical right shift and conversion of full signed to unsigned bits.
// Note, there is three instruction formula using comparison operators, but in Go strong typing booleans not converted to ints, which would require branches and more instructions. (TODO: assembly to confirm this)
func Sign(x int32) int32 { return (x >> 31) | int32((uint32(-x) >> 31)) }

func IsMostSignificantSet[T int32 | uint32](x T) bool { return !(x>>31 == 0) }
