package ch2

// AvgUnsigned does not cause overflow. Rounded down.
func AvgUnsigned(x, y uint32) uint32 { return (x & y) + ((x ^ y) >> 1) }

// AvgUnsignedCeil does not cause overflow. Rounded up.
func AvgUnsignedCeil(x, y uint32) uint32 { return (x | y) - ((x ^ y) >> 1) }

// Avg does not cause overflow. Rounded down.
// This is same as unsigned version, but uses signed-shift right.
// In Go, when left operand of shift operator is signed int, then arithmetic shift is performed.
// Arithmetic shift is same thing as signed shift.
func Avg(x, y int32) int32 { return (x & y) + ((x ^ y) >> 1) }

// AvgUnsignedCeil does not cause overflow. Rounded up.
func AvgCeil(x, y int32) int32 { return (x | y) - ((x ^ y) >> 1) }
