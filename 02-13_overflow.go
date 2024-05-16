// Overflow means result of operation result is too large or too small to fit into the variable.
// Many hardware supply bits for overflow (e.g. MIPS does not have), but high-lever languages may not have access to it.
// Unless stated opposite, formulas bellow assume carry to be 0 (version when carry is one does not work; and cannot get carry in Go).
package hd

// IsAddOverflow sets most significant bit if overflow occurs.
// This version does not use carry bit and is efficient.
func IsAddOverflow(x, y int32) int32 { return ^(x ^ y) & ((x + y) ^ x) }

func IsAddOverflow2(x, y int32) int32 { return ^(x ^ y) & ((x + y) ^ x) }

func IsAddOverflow3(x, y int32) int32 { return ((x + y) ^ x) & ((x + y) ^ y) }

// IsSubOverflow sets most significant bit if overflow occurs.
// This version does not use carry bit and is efficient.
func IsSubOverflow(x, y int32) int32 { return (x ^ y) & ((x - y) ^ x) }

func IsSubOverflow2(x, y int32) int32 { return (x ^ y) & ((x - y) ^ x) }

func IsSubOverflow3(x, y int32) int32 { return ((x - y) ^ x) & (^((x - y) ^ y)) }

// TODO: with overflow interrupts
// TODO: carry with 2^31 version

func IsAddOverflowUnsigned(x, y uint32) uint32 { return (x & y) | ((x | y) & ^(x + y)) }

func IsAddOverflowUnsigned2(x, y uint32) uint32 { return (x >> 1) + (y >> 1) + ((x & y) & 1) }

func IsAddOverflowUnsigned3(x, y uint32) bool { return ^x < y }

func IsAddOverflowUnsigned4(x, y, sum uint32) bool { return sum < x }

func IsSubOverflowUnsigned(x, y uint32) uint32 { return (^x & y) | (^(x &^ y) & (x - y)) }

func IsSubOverflowUnsigned2(x, y uint32) uint32 { return (x >> 1) - (y >> 1) - ((^x & y) & 1) }

func IsSubOverflowUnsigned3(x, y uint32) bool { return x < y }

func IsSubOverflowUnsigned4(x, y, sub uint32) bool { return sub > x }
