package hd

// AvgFloor does not cause overflow. Rounded down.
// This is same as unsigned version, but uses signed-shift right.
// In Go, when left operand of shift operator is signed int, then arithmetic shift is performed.
// Since in Go right shift is unsigned, then formula for int32 is the same as for uint32.
func AvgFloor[T int32 | uint32](x, y T) T { return (x & y) + ((x ^ y) >> 1) }

// AvgCeil does not cause overflow. Rounded up.
// In Go, when left operand of shift operator is signed int, then arithmetic shift is performed.
// Arithmetic shift is same thing as signed shift.
// Since in Go right shift is unsigned, then formula for int32 is the same as for uint32.
func AvgCeil[T int32 | uint32](x, y T) T { return (x | y) - ((x ^ y) >> 1) }
