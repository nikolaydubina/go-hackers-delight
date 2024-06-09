package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleRoundDownBlockPowerOfTwo() {
	fmt.Println(hd.RoundDownBlockPowerOfTwo[uint](257, 8))
	// Output: 256
}

func ExampleRoundUpBlockPowerOfTwo() {
	fmt.Println(hd.RoundUpBlockPowerOfTwo[uint](210, 8))
	// Output: 256
}

func ExampleFLPTwo() {
	fmt.Println(hd.FLPTwo(310))
	// Output: 256
}

func ExampleCLPTwo() {
	fmt.Println(hd.CLPTwo(200))
	// Output: 256
}

func ExampleIsPowerOfTwoBoundaryCrossed_crossed() {
	fmt.Println(hd.IsPowerOfTwoBoundaryCrossed[uint32](100, 300, 256))
	// Output: true
}

func ExampleIsPowerOfTwoBoundaryCrossed_not_crossed() {
	fmt.Println(hd.IsPowerOfTwoBoundaryCrossed[uint32](100, 100, 256))
	// Output: false
}

func powTwoBoundaries[T hd.Unsigned](x T, p uint8) (l, h T) {
	// simple algo to get boundaries
	// it is actually naive version of code that we are testing, but is easier to follow
	l = x - (x % (1 << p))
	h = l + (1 << p)
	if (x % (1 << p)) == 0 {
		h = x
	}
	return l, h
}

func fuzzRoundBlockPowerOfTwo[T hd.Unsigned](t *testing.T, x T, p uint8) {
	l, h := powTwoBoundaries(x, p)

	got := [...]struct {
		exp T
		got T
	}{
		{l, hd.RoundDownBlockPowerOfTwo(x, p)},
		{l, hd.RoundDownBlockPowerOfTwo2(x, p)},
		{h, hd.RoundUpBlockPowerOfTwo(x, p)},
		{h, hd.RoundUpBlockPowerOfTwo2(x, p)},
	}
	for i, v := range got {
		if v.got != v.exp {
			t.Error(i, x, p, "exp", v.exp, "got", v.got)
		}
	}
}

func FuzzRoundBlockPowerOfTwo_uint32(f *testing.F) {
	for p := range 32 {
		for _, x := range fuzzUint32 {
			f.Add(x, uint8(p))
		}
	}
	f.Fuzz(func(t *testing.T, x uint32, p uint8) { fuzzRoundBlockPowerOfTwo(t, x, (p % 32)) })
}

func FuzzRoundBlockPowerOfTwo_uint16(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint32, p uint8) { fuzzRoundBlockPowerOfTwo(t, x, (p % 16)) })
}

func FuzzRoundBlockPowerOfTwo_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint32, p uint8) { fuzzRoundBlockPowerOfTwo(t, x, (p % 64)) })
}

func roundPowerTwo32(x uint32) (l, h uint32) {
	if x == 0 {
		return 0, 0
	}
	for _, p := range hd.PowerOfTwo[:32] {
		p := uint32(p)
		if p <= x {
			l = p
		}
		if h == 0 && x <= p {
			h = p
		}
	}
	return l, h
}

func FuzzRoundToPowerOfTwo(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		l, h := roundPowerTwo32(x)

		got := []struct {
			exp uint32
			got uint32
		}{
			{l, hd.FLPTwo2(x)},
			{l, hd.FLPTwo3(x)},
			{l, hd.FLPTwo4(x)},
			{l, hd.FLPTwo5(x)},
			{h, hd.CLPTwo2(x)},
			{h, hd.CLPTwo3(x)},
		}
		if x > 0 {
			got = append(got, struct {
				exp uint32
				got uint32
			}{l, hd.FLPTwo(x)})
		}
		for i, v := range got {
			if v.got != v.exp {
				t.Error(i, "x", x, "exp", v.exp, "got", v.got)
			}
		}
	})
}

func isPowTwoBoundaryCrossed[T hd.Unsigned](a, l, b T) bool {
	return (uint64(a) / uint64(b)) != ((uint64(a) + uint64(l) - 1) / uint64(b))
}

func fuzzIsPowerOfTwoBoundaryCrossed[T hd.Unsigned](t *testing.T, a, l, b T) {
	if l < 3 {
		t.Skip()
	}
	exp := isPowTwoBoundaryCrossed(a, l, b)
	got := [...]bool{
		hd.IsPowerOfTwoBoundaryCrossed(a, l, b),
		hd.IsPowerOfTwoBoundaryCrossed2(a, l, b),
		hd.IsPowerOfTwoBoundaryCrossed3(a, l, b),
		hd.IsPowerOfTwoBoundaryCrossed4(a, l, b),
	}
	for i, q := range got {
		if q != exp {
			t.Error(i, a, l, b, "exp", exp, "got", q)
		}
	}
}

func FuzzIsPowerOfTwoBoundaryCrossed_uint32(f *testing.F) {
	for _, a := range fuzzUint32 {
		for _, l := range fuzzUint32 {
			for i := range hd.PowerOfTwo[:32] {
				f.Add(a, l, uint8(i))
			}
		}
	}
	f.Fuzz(func(t *testing.T, a, l uint32, ib uint8) {
		b := uint32(hd.PowerOfTwo[(ib % 32)])
		fuzzIsPowerOfTwoBoundaryCrossed(t, a, l, b)
	})
}

func FuzzIsPowerOfTwoBoundaryCrossed_uint16(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, l uint16, ib uint8) {
		b := uint16(hd.PowerOfTwo[(ib % 16)])
		fuzzIsPowerOfTwoBoundaryCrossed(t, a, l, b)
	})
}
