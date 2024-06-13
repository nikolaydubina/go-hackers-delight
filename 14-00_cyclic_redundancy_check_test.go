package hd_test

import (
	"hash/crc32"
	"math/rand/v2"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func TestCRC32GeneratorConst(t *testing.T) {
	r := hd.ReverseBits(hd.CRC32Generator)
	if r != hd.CRC32GeneratorReversed {
		t.Errorf("expected %032b, got %032b", hd.CRC32GeneratorReversed, r)
	}
}

func FuzzCRC32(f *testing.F) {
	f.Fuzz(func(t *testing.T, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 byte) {
		data := []byte{x1, x2, x3, x4, x5, x6, x7, x8, x9, x10}

		exp := crc32.ChecksumIEEE(data)
		got := []uint32{
			hd.CRC32Basic(data),
			hd.CRC32TableLookup(data),
		}

		for i, got := range got {
			if exp != got {
				t.Error(i, "exp", exp, "got", got)
			}
		}
	})
}

func BenchmarkCRC32(b *testing.B) {
	var out uint32

	var vals [][100]byte
	for i := 0; i < 1000; i++ {
		var v [100]byte
		for j := range 100 {
			v[j] = byte(rand.Int32())
		}
		vals = append(vals, v)
	}

	vs := []struct {
		name string
		f    func(data []byte) uint32
	}{
		{"basic", crc32.ChecksumIEEE},
		{"CRC32Basic", hd.CRC32Basic},
		{"CRC32TableLookup", hd.CRC32TableLookup},
	}
	for _, v := range vs {
		b.Run(v.name, func(b *testing.B) {
			for i := 0; i < b.N; i += len(vals) {
				for j := 0; j < len(vals)-1; j++ {
					out = v.f(vals[j][:])
				}
			}
		})
	}

	if (out*2 - out - out) != 0 {
		b.Fatal("never")
	}
}
