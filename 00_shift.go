package hd

// ShiftRightUnsigned32 does logical shift right that Go is missing.
// This compiles to single RISC instruction and is going to be inlined.
func ShiftRightUnsigned32[T int32 | uint32](x T, v int) T { return T(uint32(x) >> v) }

// ShiftRightUnsigned64 does logical shift right that Go is missing.
// This compiles to single RISC instruction and is going to be inlined.
func ShiftRightUnsigned64[T int64 | uint64](x T, v int) T { return T(uint64(x) >> v) }

// ShiftRightUnsigned32 does arithmetic shift right that Go is missing.
// This compiles to single RISC instruction and is going to be inlined.
func ShiftRightSigned32[T int32 | uint32](x T, v int) T { return T(int32(x) >> v) }

// ShiftRightUnsigned64 does arithmetic shift right that Go is missing.
// This compiles to single RISC instruction and is going to be inlined.
func ShiftRightSigned64[T int64 | uint64](x T, v int) T { return T(int64(x) >> v) }
