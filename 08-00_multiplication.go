package hd

func Int64To4x16b(v int64) [4]uint16 {
	return [4]uint16{
		uint16((v & (0xFFFF << 00)) >> 00),
		uint16((v & (0xFFFF << 16)) >> 16),
		uint16((v & (0xFFFF << 32)) >> 32),
		uint16((uint64(v) & (0xFFFF << 48)) >> 48),
	}
}

func Int64From4x16b(v [4]uint16) int64 {
	return int64(v[0]) | int64(v[1])<<16 | int64(v[2])<<32 | int64(v[3])<<48
}

// MultiplyMultiWord (aka mulmns) multiplies two multiwords word-wise. w = u * v
// This does not overflow.
// We are using uint16 and uint32 to avoid overflow in word multiplication.
// Most important word can be negative when converted to int16.
// Refer to routines Int64To4x16b and Int64From4x16b for conversion.
func MultiplyMultiWord(w, u, v []uint16) {
	if len(w) != (len(u) + len(v)) {
		panic("len(w) != len(u) + len(v)")
	}

	var k, t, b uint32

	for j := range v {
		k = 0
		for i := range u {
			t = uint32(u[i])*uint32(v[j]) + uint32(w[i+j]) + k
			w[i+j] = uint16(t)
			k = t >> 16
		}
		w[j+len(u)] = uint16(k)
	}

	// Now w[] has the unsigned product. Correct by
	// subtracting v*2**16m, if u < 0, and
	// subtracting u*2**16n, if v < 0.
	if int16(u[len(u)-1]) < 0 {
		for j := range v {
			t = uint32(w[j+len(u)]) - uint32(v[j]) - b
			w[j+len(u)] = uint16(t)
			b = t >> 31
		}
	}
	if int16(v[len(v)-1]) < 0 {
		for i := range u {
			t = uint32(w[i+len(v)]) - uint32(u[i]) - b
			w[i+len(v)] = uint16(t)
			b = t >> 31
		}
	}
}
