package hd

// Compress (aka Generalized Extract) selects bits from the x where m is 1 and puts them in order to least significant bits.
// This executes in 169 instructions on 64bit RISC.
func Compress(x, m uint32) uint32 {
	var mp, mv, t uint32

	x = x & m                // Clear irrelevant bits.
	mk := ^m << 1            // We will count 0's to right.
	for i := 0; i < 5; i++ { //
		mp = mk ^ (mk << 1)             // Parallel suffix.
		mp = mp ^ (mp << 2)             //
		mp = mp ^ (mp << 4)             //
		mp = mp ^ (mp << 8)             //
		mp = mp ^ (mp << 16)            //
		mv = mp & m                     // Bits to move.
		m = (m ^ mv) | (mv >> (1 << i)) // Compress m.
		t = x & mv                      //
		x = x ^ t | (t >> (1 << i))     // Compress x.
		mk = mk & ^mp                   // Clear suffix.
	}

	return x
}

// Compress2 is direct version of Compress. This is 260 RISC instructions in loop brach-free.
func Compress2(x, m uint32) uint32 {
	var r, s, b uint32

	for m != 0 {
		b = m & 1
		r = r | ((x & b) << s)
		s = s + b
		x = x >> 1
		m = m >> 1
	}

	return r
}
