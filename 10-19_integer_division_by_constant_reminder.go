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
