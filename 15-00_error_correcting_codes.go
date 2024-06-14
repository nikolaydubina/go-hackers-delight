package hd

// CheckBits produces 6bits of parity checks.
// It is simple Hamming Code based error detection and correction scheme.
func CheckBits(u uint32) uint32 {
	p := [6]uint32{}

	p[0] = u
	p[0] ^= p[0] >> 2
	p[0] ^= p[0] >> 4
	p[0] ^= p[0] >> 8
	p[0] ^= p[0] >> 16

	t := u ^ (u >> 1)
	p[1] = t
	p[1] ^= p[1] >> 4
	p[1] ^= p[1] >> 8
	p[1] ^= p[1] >> 16

	t ^= t >> 2
	p[2] = t
	p[2] ^= p[2] >> 8
	p[2] ^= p[2] >> 16

	t ^= t >> 4
	p[3] = t
	p[3] ^= p[3] >> 16

	t ^= t >> 8
	p[4] = t

	p[5] = p[4]
	p[5] ^= p[5] >> 16

	var pp uint32
	pp |= (p[0] >> 1) & 1
	pp |= (p[1] >> 1) & 2
	pp |= (p[2] >> 2) & 4
	pp |= (p[3] >> 5) & 8
	pp |= (p[4] >> 12) & 16
	pp |= (p[5] & 1) << 5

	pp ^= (-(u & 1) & 0x3F)
	return pp
}

// Syndrome in Hamming based error checking code
func Syndrome(pr, u uint32) uint32 {
	p := CheckBits(u)
	syn := p ^ (pr & 0x3F)
	return syn
}

// Correct bits in pr that contain 32 "information" bits stored in ur and 7 "check" bits stored in pr.
// Check bits are extracted and error is corrected in returned result.
// This can correct up to 1 wrong bit.
// This is based on Hamming error correcting code.
// It returns corrected bits and number of error bits detected (0,1,2).
// TODO: 2 bits error detection is not working.
func Correct(pr, u uint32) (cr uint32, errb uint8) {
	// TODO: why parity for detection of 2 errors is not working?
	syn := Syndrome(pr, u)
	if syn == 0 {
		return u, 0
	}

	// One error ocurred, but error is in check bits, so no correction of "information" bits required.
	if ((syn - 1) & syn) == 0 {
		return u, 1
	}

	b := syn - 31 - (syn >> 5) // map syn to range [0, 31]
	return u ^ (1 << b), 1
}
