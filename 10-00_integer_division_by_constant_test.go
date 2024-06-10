package hd_test

import (
	"math/rand/v2"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func BenchmarkDivMod(b *testing.B) {
	var out int32

	var vals []int32
	for i := 0; i < 1000; i++ {
		vals = append(vals, rand.Int32())
	}

	b.Run("DivMod", func(b *testing.B) {
		vs := []struct {
			name string
			f    func(int32) (int32, int32)
		}{
			{"3/basic", func(x int32) (int32, int32) { return x / 3, x % 3 }},
			{"3/DivMod3Signed", hd.DivMod3Signed},
			{"3/DivMod3Signed2", hd.DivMod3Signed2},
			{"7/basic", func(x int32) (int32, int32) { return x / 7, x % 7 }},
			{"7/DivMod7Signed", hd.DivMod7Signed},
		}
		for _, v := range vs {
			b.Run(v.name, func(b *testing.B) {
				for i := 0; i < b.N; i += len(vals) {
					for j := 0; j < len(vals); j++ {
						out, out = v.f(vals[j])
					}
				}
			})
		}
	})

	b.Run("Div", func(b *testing.B) {
		vs := []struct {
			name string
			f    func(int32) int32
		}{
			{"3/basic", func(x int32) int32 { return x / 3 }},
			{"3/Div3Signed", hd.Div3Signed},
			{"3/Div3ShiftSigned", hd.Div3ShiftSigned},
			{"7/basic", func(x int32) int32 { return x / 7 }},
			{"7/Div7Signed", hd.Div7Signed},
			{"7/Div7ShiftSigned", hd.Div7ShiftSigned},
		}
		for _, v := range vs {
			b.Run(v.name, func(b *testing.B) {
				for i := 0; i < b.N; i += len(vals) {
					for j := 0; j < len(vals); j++ {
						out = v.f(vals[j])
					}
				}
			})
		}
	})

	b.Run("Mod", func(b *testing.B) {
		vs := []struct {
			name string
			f    func(int32) int32
		}{
			{"3/basic", func(x int32) int32 { return x % 3 }},
			{"3/Mod3Signed", hd.Mod3Signed},
			{"3/Mod3Signed2", hd.Mod3Signed2},
			{"7/basic", func(x int32) int32 { return x % 7 }},
			{"7/Mod7Signed", hd.Mod7Signed},
			{"7/Mod7Signed2", hd.Mod7Signed2},
			{"10/basic", func(x int32) int32 { return x % 10 }},
			{"10/Mod10Signed", hd.Mod10Signed},
		}
		for _, v := range vs {
			b.Run(v.name, func(b *testing.B) {
				for i := 0; i < b.N; i += len(vals) {
					for j := 0; j < len(vals); j++ {
						out = v.f(vals[j])
					}
				}
			})
		}
	})

	b.Run("DivExact", func(b *testing.B) {
		var vals []int32
		for v := int32(0); len(vals) < 1000; v = rand.Int32() {
			if (v % 7) == 0 {
				vals = append(vals, v)
			}
		}

		vs := []struct {
			name string
			f    func(int32) int32
		}{
			{"7/basic", func(x int32) int32 { return x / 7 }},
			{"7/DivExact7", hd.DivExact7},
			{"7/Div7Signed", hd.Div7Signed},
			{"7/Div7ShiftSigned", hd.Div7ShiftSigned},
		}
		for _, v := range vs {
			b.Run(v.name, func(b *testing.B) {
				for i := 0; i < b.N; i += len(vals) {
					for j := 0; j < len(vals); j++ {
						out = v.f(vals[j])
					}
				}
			})
		}
	})

	if (out*2 - out - out) != 0 {
		b.Fatal("never")
	}
}
