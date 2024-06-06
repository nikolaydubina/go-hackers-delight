package hd_test

import (
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func TestLRUCache(t *testing.T) {
	var cache hd.LRUCache

	cache.Hit(0)
	cache.Hit(1)
	cache.Hit(2)
	cache.Hit(3)
	cache.Hit(4)
	cache.Hit(5)
	cache.Hit(6)
	cache.Hit(7)

	if cache.LeastRecentlyUsed() != 0 {
		t.Errorf("expected %d, got %d: cache=%032b", 0, cache.LeastRecentlyUsed(), cache)
	}

	cache.Hit(0)
	if cache.LeastRecentlyUsed() != 1 {
		t.Errorf("expected %d, got %d: cache=%032b", 1, cache.LeastRecentlyUsed(), cache)
	}
}

func FuzzLRUCache(f *testing.F) {
	var cache hd.LRUCache
	cacheBasic := NewLRUCacheBasic()

	count := 0
	f.Fuzz(func(t *testing.T, x uint8) {
		x = x % 8
		count++

		cache.Hit(x)
		cacheBasic.Hit(x)

		// fill out the cache
		if count < 8 {
			return
		}

		if cache.LeastRecentlyUsed() != cacheBasic.LeastRecentlyUsed() {
			t.Errorf("x=%v expected %d, got %d: cache=%016x cache_basic=%v", x, cacheBasic.LeastRecentlyUsed(), cache.LeastRecentlyUsed(), cache, cacheBasic)
		}
	})
}

// LRUCacheBasic is reference implementation that uses array and simple movement of all elements.
type LRUCacheBasic struct {
	vals [8]uint8
}

func NewLRUCacheBasic() *LRUCacheBasic { return &LRUCacheBasic{vals: [8]uint8{0, 1, 2, 3, 4, 5, 6, 7}} }

func (m *LRUCacheBasic) Hit(i uint8) {
	i = i % 8
	for j, v := range m.vals {
		if v == i {
			copy(m.vals[1:j+1], m.vals[:j])
			m.vals[0] = i
			return
		}
	}
}

func (m *LRUCacheBasic) LeastRecentlyUsed() uint8 { return m.vals[len(m.vals)-1] }
