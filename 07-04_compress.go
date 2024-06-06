package hd

// Compress (aka Generalized Extract) selects bits from the x where m is 1 and puts them in order to least significant bits.
// This uses divide and conquer.
// This executes in 169 instructions on 64bit RISC.
// There is also hardware focused version that does not work well in general RISC.
func Compress(x, m uint32) uint32 {
	x &= m                   // Clear irrelevant bits.
	mk := ^m << 1            // We will count 0's to right.
	for i := 0; i < 5; i++ { //
		mp := mk ^ (mk << 1)            // Parallel suffix.
		mp ^= mp << 2                   //
		mp ^= mp << 4                   //
		mp ^= mp << 8                   //
		mp ^= mp << 16                  //
		mv := mp & m                    // Bits to move.
		m = (m ^ mv) | (mv >> (1 << i)) // Compress m.
		t := x & mv                     //
		x = x ^ t | (t >> (1 << i))     // Compress x.
		mk &= ^mp                       // Clear suffix.
	}
	return x
}

// Compress2 is direct version of Compress. This is 260 RISC instructions in loop brach-free.
func Compress2(x, m uint32) uint32 {
	var r, s, b uint32
	for m != 0 {
		b = m & 1
		r = r | ((x & b) << s)
		s += b
		x >>= 1
		m >>= 1
	}
	return r
}

// Expand (aka Generalized Insert, unpack, scatter, deposit) selects bits from the x where m is 1 and puts them in order to least significant bits.
// This is reverse of Compress steps.
// This is 168 RISC instructions, or 200 instructions on 64bit RISC.
// There is also hardware focused version that does not work well in general RISC.
func Expand(x, m uint32) uint32 {
	var array [5]uint32

	m0 := m       // Save original mask.
	mk := ^m << 1 // We will count 0's to right.

	for i := 0; i < 5; i++ {
		mp := mk ^ (mk << 1)            // Parallel suffix.
		mp ^= mp << 2                   //
		mp ^= mp << 4                   //
		mp ^= mp << 8                   //
		mp ^= mp << 16                  //
		mv := mp & m                    // Bits to move.
		array[i] = mv                   //
		m = (m ^ mv) | (mv >> (1 << i)) // Compress m.
		mk &= ^mp
	}

	for i := 4; i >= 0; i-- {
		mv := array[i]
		t := x << (1 << i)
		x = (x & ^mv) | (t & mv)
	}

	return x & m0 // Clear out extraneous bits.
}

// TODO: compress left, GeneralizedUnshuffle (aka SheepAndGoats, SAG), bitgather, GRP
