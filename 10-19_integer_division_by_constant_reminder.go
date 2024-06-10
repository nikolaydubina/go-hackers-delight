package hd

// Mod3Unsigned (remu3) is using CountOnes (pop).
func Mod3Unsigned(n uint32) uint32 {
	n = CountOnes((n ^ 0xAAAA_AAAA)) + 23            // Now 23 <= n <= 55
	n = CountOnes((n ^ 0x2A)) - 3                    // Now -3 <= n <= 2
	return uint32(int32(n) + ((int32(n) >> 31) & 3)) // (Signed shift).
}

var mod3_unsigned_2 = [...]uint8{
	2,
	0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2,
	0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2,
	0, 1,
}

// Mod3Unsigned2 (remu3) is using CountOnes (pop) and table lookup.
func Mod3Unsigned2(n uint32) uint32 { return uint32(mod3_unsigned_2[CountOnes((n ^ 0xAAAA_AAAA))]) }

// Mod3Unsigned3 (remu3) is using digit summing and in-register lookup.
func Mod3Unsigned3(n uint32) uint32 {
	n = (n >> 16) + (n & 0xFFFF) // Max 0x1FFFE
	n = (n >> 8) + (n & 0x00FF)  // Max 0x2FD
	n = (n >> 4) + (n & 0x000F)  // Max 0x3D
	n = (n >> 2) + (n & 0x0003)  // Max 0x11
	n = (n >> 2) + (n & 0x0003)  // Max 0x6
	return (0x0924 >> (n << 1)) & 3
}

var mod3_unsigned_4 = [...]uint8{
	0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2,
	0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2,
	0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2,
	0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2,
	0, 1,
}

// Mod3Unsigned3 (remu3) is using digit summing and in-memory lookup.
func Mod3Unsigned4(n uint32) uint32 {
	n = (n >> 16) + (n & 0xFFFF) // Max 0x1FFFE
	n = (n >> 8) + (n & 0x00FF)  // Max 0x2FD
	n = (n >> 4) + (n & 0x000F)  // Max 0x3D
	return uint32(mod3_unsigned_4[n])
}

// Mod5Unsigned (remu5) is using CountOnes (pop) and in-memory lookup.
func Mod5Unsigned(n uint32) uint32 {
	n = (n >> 16) + (n & 0xFFFF)              // Max 0x1FFFE
	n = (n >> 8) + (n & 0x00FF)               // Max 0x2FD
	n = (n >> 4) + (n & 0x000F)               // Max 0x3D
	n = (n >> 4) - ((n >> 2) & 3) + (n & 3)   // -3 to 6
	return (01043210432 >> (3 * (n + 3))) & 7 // Octal const
}

var mod7_unsigned = [...]uint8{
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4,
}

// Mod7Unsigned (remu7) is using summing method and in-memory lookup.
func Mod7Unsigned(n uint32) uint32 {
	n = (n >> 15) + (n & 0x7FFF) // Max 0x27FFE
	n = (n >> 9) + (n & 0x001FF) // Max 0x33D
	n = (n >> 6) + (n & 0x0003F) // Max 0x4A
	return uint32(mod7_unsigned[n])
}

var mod9_unsigned = [...]uint8{
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2,
}

// Mod9Unsigned (remu9) is using summing method and in-memory lookup.
func Mod9Unsigned(n uint32) uint32 {
	r := int32((n & 0x7FFF) - (n >> 15)) // FFFE0001 to 7FFF
	r = (r & 0x01FF) - (r >> 9)          // FFFFFC1 to 2FF
	r = (r & 0x003F) + (r >> 6)          // 0 to 4A
	return uint32(mod9_unsigned[r])
}

// Mod3Signed (rems3) is similar to unsigned version, but remaps output of it differently.
func Mod3Signed(n int32) int32 {
	r := uint32(n)
	r = (r >> 16) + (r & 0xFFFF) // Max 0x1FFFE
	r = (r >> 8) + (r & 0x00FF)  // Max 0x2FD
	r = (r >> 4) + (r & 0x000F)  // Max 0x3D
	r = uint32(mod3_unsigned_4[r])
	return int32(r) - (int32((uint32(n) >> 31)) << (r & 2))
}

var mod5_signed = [...]uint8{
	2, 3, 4,
	0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
	0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4,
	0, 1, 2, 3, 4,
	0, 1, 2, 3,
}

// Mod5Signed (rems5) is similar to unsigned version, but remaps output of it differently.
func Mod5Signed(n int32) int32 {
	r := n
	r = (r >> 16) + (r & 0xFFFF) // FFFF8000 to 17FFE
	r = (r >> 8) + (r & 0x00FF)  // FFFFFF80 to 27D
	r = (r >> 4) + (r & 0x000F)  // -8 to 53 (decimal)
	r = int32(mod5_signed[(r + 8)])
	return r - (((n & -r) >> 31) & 5)
}

var mod7_signed = [...]uint8{
	5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4, 5, 6,
	0, 1, 2,
}

// Mod7Signed (rems7) is similar to unsigned version, but remaps output of it differently.
func Mod7Signed(n int32) int32 {
	r := n
	r = (r >> 15) + (n & 0x7FFF) // FFFF0000 to 17FFE
	r = (r >> 9) + (r & 0x001FF) // FFFFFF80 to 2BD
	r = (r >> 6) + (r & 0x0003F) // -2 to 72 (decimal)
	r = int32(mod7_signed[(r + 2)])
	return r - (((n & -r) >> 31) & 7)
}

var mod9_signed = [...]uint8{
	7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0, 1, 2, 3, 4, 5, 6, 7, 8,
	0,
}

// Mod9Signed (rems7) is similar to unsigned version, but remaps output of it differently.
func Mod9Signed(n int32) int32 {
	r := n
	r = (r & 0x7FFF) - (r >> 15) // FFFF7001 to 17FFF
	r = (r & 0x01FF) - (r >> 9)  // FFFFFF41 to 0x27F
	r = (r & 0x003F) + (r >> 6)  // -2 to 72 (decimal)
	r = int32(mod9_signed[(r + 2)])
	return r - (((n & -r) >> 31) & 9)
}
