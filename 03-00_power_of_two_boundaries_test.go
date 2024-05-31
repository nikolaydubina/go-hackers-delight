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
	fmt.Println(hd.IsPowerOfTwoBoundaryCrossed(100, 300, 256))
	// Output: true
}

func ExampleIsPowerOfTwoBoundaryCrossed_not_crossed() {
	fmt.Println(hd.IsPowerOfTwoBoundaryCrossed(100, 100, 256))
	// Output: false
}

func FuzzRoundBlockPowerOfTwo(f *testing.F) {
	for p := range 32 {
		for _, x := range fuzzUint32 {
			f.Add(x, p)
		}
	}
	f.Fuzz(func(t *testing.T, x uint32, p int) {
		if p <= 0 || p >= 32 {
			t.Skip()
		}

		// simple algo to get boundaries
		// it is actually naive version of code that we are testing, but is easier to follow
		l := x - (x % (1 << p))
		h := l + (1 << p)
		if (x % (1 << p)) == 0 {
			h = x
		}

		vs := []struct {
			exp uint32
			got uint32
		}{
			{l, hd.RoundDownBlockPowerOfTwo(x, p)},
			{l, hd.RoundDownBlockPowerOfTwo2(x, p)},
			{h, hd.RoundUpBlockPowerOfTwo(x, p)},
			{h, hd.RoundUpBlockPowerOfTwo2(x, p)},
		}
		for i, v := range vs {
			if v.got != v.exp {
				t.Error(i, x, p, "exp", v.exp, "got", v.got)
			}
		}
	})
}

func FuzzRoundToPowerOfTwo(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		// definition
		var l, h uint32
		if x > 0 {
			for _, p := range hd.PowerOfTwo[:32] {
				p := uint32(p)
				if p <= x {
					l = p
				}
				if h == 0 && x <= p {
					h = p
				}
			}
		}

		vs := []struct {
			exp uint32
			got uint32
		}{
			{l, hd.FLPTwo(x)},
			{l, hd.FLPTwo2(x)},
			{l, hd.FLPTwo3(x)},
			{l, hd.FLPTwo4(x)},
			{l, hd.FLPTwo5(x)},
			{h, hd.CLPTwo(x)},
			{h, hd.CLPTwo2(x)},
			{h, hd.CLPTwo3(x)},
		}
		for i, v := range vs {
			if v.got != v.exp {
				t.Error(i, "x", x, "exp", v.exp, "got", v.got)
			}
		}
	})
}

func FuzzIsPowerOfTwoBoundaryCrossed(f *testing.F) {
	for _, a := range fuzzUint32 {
		for _, l := range fuzzUint32 {
			for i := range hd.PowerOfTwo[:32] {
				f.Add(a, l, uint8(i))
			}
		}
	}

	f.Fuzz(func(t *testing.T, a, l uint32, ib uint8) {
		if int(ib) >= 31 {
			t.Skip()
		}
		if l < 3 {
			t.Skip()
		}
		b := uint32(hd.PowerOfTwo[ib])

		// naive approach, relying on uint64 to protect overflows
		isCrossed := (uint64(a) / uint64(b)) != ((uint64(a) + uint64(l) - 1) / uint64(b))

		vs := []struct {
			exp bool
			got bool
		}{
			{isCrossed, hd.IsPowerOfTwoBoundaryCrossed(a, l, b)},
			{isCrossed, hd.IsPowerOfTwoBoundaryCrossed2(a, l, b)},
			{isCrossed, hd.IsPowerOfTwoBoundaryCrossed3(a, l, b)},
			{isCrossed, hd.IsPowerOfTwoBoundaryCrossed4(a, l, b)},
		}
		for i, v := range vs {
			if v.got != v.exp {
				t.Error(i, a, l, b, "exp", v.exp, "got", v.got)
			}
		}
	})
}
