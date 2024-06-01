package hd_test

import (
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func TestLRUCache(t *testing.T) {
	var cache hd.LRUCache

	cache.Add(0)
	cache.Add(1)
	cache.Add(2)
	cache.Add(3)
	cache.Add(4)
	cache.Add(5)
	cache.Add(6)
	cache.Add(7)

	if cache.LeastRecentlyUsed() != 0 {
		t.Errorf("expected %d, got %d: cache=%032b", 0, cache.LeastRecentlyUsed(), cache)
	}

	cache.Add(0)
	if cache.LeastRecentlyUsed() != 1 {
		t.Errorf("expected %d, got %d: cache=%032b", 1, cache.LeastRecentlyUsed(), cache)
	}
}
