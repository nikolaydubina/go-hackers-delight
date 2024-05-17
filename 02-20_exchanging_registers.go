package hd

// ExchangeRegisters illustrates very old trick on how to exchange two registers without using third one.
// This is swap operation. Also known as multiple assignment in Go.
func ExchangeRegisters(x, y int32) (int32, int32) {
	x = x ^ y
	y = y ^ x
	x = x ^ y
	return x, y
}
