package ch2

import "testing"

func Fuzz_InequalitiesAmongLogicAndArithmetic(f *testing.F) {
	f.Fuzz(func(t *testing.T, x, y uint64) {
		if !((x ^ y) <= (x | y)) {
			t.Error(x, y)
		}
	})
}
