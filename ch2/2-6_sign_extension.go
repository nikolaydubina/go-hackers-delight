package ch2

// ExtendSign8 sign-extends a 8-bit number to a 32-bit number.
// Sign of 8-bit number is stored in 8th-bit.
// Sign extension is treating n-th least significant bit as sign bit and copying it to all more significant bits.
func ExtendSign7(x uint32) uint32 { return ((x + 0x00000080) & 0x000000FF) - 0x00000080 }

// ExtendSign7Two is alternative version.
// If you know all higher order bits are zero, then `and` can be omitted.
func ExtendSign7Two(x uint32) uint32 { return ((x & 0x000000ff) ^ 0x00000080) - 0x00000080 }

// ExtendSign7Three is alternative version
func ExtendSign7Three(x uint32) uint32 { return (x & 0x0000007f) - (x & 0x00000080) }
