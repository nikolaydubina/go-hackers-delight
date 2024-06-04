/*
# Chapter 8: Multiplication

Any multiplication can be decomposed in a summation of left shifts.
For example x * 13 binary (1101) is:
- t1 := x << 2
- t2 := x << 3
- result := t1 + t2 + x

This is decomposed into result := 8x + 4x + x.

For every multiplication there are multiple paths possible that utilize registers and shifts.
These paths can have fewer or more instructions and registers to accomplish this.
In general, it is nontrivial to find minimum number instructions required and which these instructions are.
Theorem bellow gives upper bound.

Theorem. Multiplication of a variable x by an n-bit constant m, m >= 1,
can be accomplished with at most n instructions of the type add, subtract, and shift left by any given amount.
*/
package hd

// MulMultiWord (aka mulmns) multiplies two multiwords word-wise. w = u * v
// This does not overflow.
// We are using uint16 and uint32 to avoid overflow in word multiplication.
// Most important word can be negative when converted to int16.
// Refer to routines Int64To4x16b and Int64From4x16b for conversion.
func MulMultiWord(w, u, v []uint16) {
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

// MultiplyHighOrderSigned (aka mulhns) multiplies two signed integers and returns the high-order half of the product.
// This executes in 16 RISC instructions.
func MultiplyHighOrderSigned(u, v int32) int32 {
	var u0, v0, w0 uint32
	var u1, v1, w1, w2, t int32

	u0 = uint32(u & 0xFFFF)
	u1 = u >> 16
	v0 = uint32(v & 0xFFFF)
	v1 = v >> 16
	w0 = u0 * v0
	t = int32((uint32(u1) * v0) + (w0 >> 16))
	w1 = t & 0xFFFF
	w2 = t >> 16
	w1 = int32((u0 * uint32(v1))) + w1

	return (u1 * v1) + w2 + (w1 >> 16)
}

// TODO: MultiplyHighOrderUnsigned
// TODO: MultiplyHighOrderUnsigned from MultiplyHighOrderSigned and other way around
