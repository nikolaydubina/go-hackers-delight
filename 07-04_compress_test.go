package hd_test

import (
	"fmt"
	"math/rand"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleCompress() {
	//          x:   abcd_efgh_ijkl_mnop_qrst_uvwx_yzAB_CDEF
	var x uint32 = 0b1010_1011_0101_1110_0001_0101_0000_1000
	var m uint32 = 0b0000_1111_0011_0011_1010_1010_0101_0101
	// 		  exp:   0000_0000_0000_0000_efgh_klop_qsuw_zBDF
	fmt.Printf("%032b", hd.Compress(x, m))
	// Output: 00000000000000001011011000000000
}

func FuzzCompressEquality(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, m := range fuzzUint32 {
			f.Add(x, m)
		}
	}
	f.Fuzz(func(t *testing.T, x, m uint32) {
		exp := hd.Compress(x, m)

		vs := []uint32{
			hd.Compress2(x, m),
		}
		for i, v := range vs {
			if v != exp {
				t.Errorf("%d: %d != %d", i, v, exp)
			}
		}
	})
}

func FuzzCompressIdentity(f *testing.F) {
	for _, x := range fuzzUint32 {
		for _, m := range fuzzUint32 {
			f.Add(x, m)
		}
	}
	f.Fuzz(func(t *testing.T, x, m uint32) {
		if (x & m) != hd.Expand(hd.Compress(x, m), m) {
			t.Error(x, m)
		}
	})
}

func BenchmarkCompress(b *testing.B) {
	var vs []uint32
	for i := 0; i < 100; i++ {
		vs = append(vs, rand.Uint32())
	}

	b.Run("Compress", func(b *testing.B) {
		b.ResetTimer()
		var v uint32
		for n := 0; n < b.N; n++ {
			v = hd.Compress(vs[n%len(vs)], vs[(n+1)%len(vs)])
			if (v + 1) == 0 {
				b.Error(v)
			}
		}
	})

	b.Run("Compress2", func(b *testing.B) {
		b.ResetTimer()
		var v uint32
		for n := 0; n < b.N; n++ {
			v = hd.Compress2(vs[n%len(vs)], vs[(n+1)%len(vs)])
			if (v + 1) == 0 {
				b.Error(v)
			}
		}
	})
}
