package hd

// LRUCache is eight-way set associative cache with LRU replacement policy that uses reference matrix method.
// This can count up to 8 lines, indexed [0, 7].
// Least significant byte holds row 0.
// The whole structure fits into single 64bit register.
// Line is is references by i-th byte of the register.
type LRUCache struct{ m uint64 }

// Add line i to the cache.
// This is five or six instructions on 64bit RISC.
// Values of i should be in [0, 7].
func (c *LRUCache) Add(i uint8) {
	c.m |= (0xFF << (8 * i))
	c.m &= ^(0x0101_0101_0101_0101 << i)
}

func (c *LRUCache) LeastRecentlyUsed() uint8 { return uint8(7 - ZByteL64(uint64(c.m))) }
