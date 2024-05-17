package hd

type integers interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

// ExchangeRegisters illustrates very old trick on how to exchange two registers without using third one.
// This is swap operation. Also known as multiple assignment in Go.
func ExchangeRegisters[T integers](x, y T) (T, T) {
	x = x ^ y
	y = y ^ x
	x = x ^ y
	return x, y
}

// ExchangeRegistersMasked illustrates how to exchange only masked bits between two registers.
// This can be done in 3 cycles, given parallelism and and-not instructions.
func ExchangeRegistersMasked[T integers](x, y, m T) (T, T) {
	t := (x & ^m) | (y & m)
	y = (y & ^m) | (x & m)
	x = t
	return x, y
}

func ExchangeRegistersMasked2[T integers](x, y, m T) (T, T) {
	x = x ^ y
	y = y ^ (x & m)
	x = x ^ y
	return x, y
}

// ExchangeRegistersMasked3 is heavily using  bitwise equality, but in Go there is no such operator, so expanding that notation.
func ExchangeRegistersMasked3[T integers](x, y, m T) (T, T) {
	x = ^(x ^ y)
	y = ^(y ^ (x | ^m))
	x = ^(x ^ y)
	return x, y
}

// ExchangeRegistersMasked4 also executes in three cycles, given sufficient instruction parallelism in the CPU.
func ExchangeRegistersMasked4[T integers](x, y, m T) (T, T) {
	t := (x ^ y) & m
	x = x ^ t
	y = y ^ t
	return x, y
}

// ExchangeBitsInRegister swaps two regions of bits within single register without affecting other bits.
// For example, swapping B and D regions. B and D should be same size, but A C E and B/D can be different sizes.
// k is number of bits of C + D
// md is mask for field D
// mo is mask for fields A C and E (all fields to not swap)
// [aaaa bbbb cccc dddd eeee] -> [aaaa dddd cccc bbbb eeee]
// This is straight forward approach and requires eleven instructions and six cycles.
func ExchangeBitsInRegister(x, md, mo uint32, k int) uint32 {
	t1 := (x & md) << k
	t2 := (x >> k) & md
	return (x & mo) | t1 | t2
}

// ExchangeBitsInRegisterFast requires eight instructions and five cycles.
func ExchangeBitsInRegisterFast(x, md, mo uint32, k int) uint32 {
	t1 := (x ^ (x >> k)) & md
	t2 := t1 << k
	return x ^ t1 ^ t2
}
