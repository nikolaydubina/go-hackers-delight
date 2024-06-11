package hd

// Cbrt is cube root, it starts of with shift based Newton algorithm for Sqrt and modifies it to cuber root
func Cbrt(x uint32) uint32 {
	var y uint32
	for s := int32(30); s >= 0; s -= 3 {
		y *= 2
		if b := ((3 * y * (y + 1)) + 1) << s; x >= b {
			x -= b
			y++
		}
	}
	return y
}
