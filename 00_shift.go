package hd

// ShiftRightUnsigned32 does logical shift right that Go is missing.
// This compiles to single RISC instruction and is going to be inlined.
func ShiftRightUnsigned32(x int32, v int) int32 { return int32(uint32(x) >> v) }

// ShiftRightUnsigned64 does logical shift right that Go is missing.
// This compiles to single RISC instruction and is going to be inlined.
func ShiftRightUnsigned64(x int64, v int) int64 { return int64(uint64(x) >> v) }
