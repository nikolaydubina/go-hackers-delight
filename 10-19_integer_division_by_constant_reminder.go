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
