package hd

// GrayCodeFromUint converts a binary number to Gray code.
// Gray Code is sequence of binary numbers where each successive differs only by one bit.
// It is possible to iterate through all 2ⁿ numbers in single Gray code.
// There are many versions of Gray Codes, here we use most common one, "reflected binary Gray code".
// Gray codes named after Frank Gray, physicist at Bell Labs, who invented it in 1930s for TV.
func GrayCodeFromUint[T Unsigned](n T) T { return n ^ (n >> 1) }

func GrayCodeToUint16(g uint16) (b uint16) {
	for i := uint16(0); i < 16; i++ {
		b ^= g >> i
	}
	return b
}

func GrayCodeToUint32(g uint32) (b uint32) {
	for i := uint32(0); i < 32; i++ {
		b ^= g >> i
	}
	return b
}

func GrayCodeToUint64(g uint64) (b uint64) {
	for i := uint64(0); i < 64; i++ {
		b ^= g >> i
	}
	return b
}
