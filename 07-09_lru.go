package hd

// LRUCache is eight-way set associative cache with least-recently-used replacement policy that uses reference matrix method.
// The whole structure fits into single 64bit register.
// Internally, least significant byte of uint64 holds row 0 of reference matrix.
type LRUCache struct{ m uint64 }

// Hit value i as most recently used.
// This is five or six instructions on 64bit RISC.
// Values of i should be in [0, 7].
func (c *LRUCache) Hit(i uint8) {
	c.m |= (0xFF << (8 * i))
	c.m &= ^(0x0101_0101_0101_0101 << i)
}

func (c *LRUCache) LeastRecentlyUsed() uint8 { return uint8(7 - ZByteL64(uint64(c.m))) }
