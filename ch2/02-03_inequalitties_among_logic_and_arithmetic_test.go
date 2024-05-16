package ch2

import "testing"

func willSumOverflow(x, y uint64) bool { return (x+y < x) || (x+y < y) }

func absDiff(x, y uint64) uint64 { return max(x, y) - min(x, y) }

func Fuzz_InequalitiesAmongLogicAndArithmetic(f *testing.F) {
	f.Fuzz(func(t *testing.T, x, y uint64) {
		v := []bool{
			(x ^ y) <= (x | y),
			(x | y) >= max(x, y),
			(x & y) <= min(x, y),
			((x | y) <= (x + y)) && !willSumOverflow(x, y),
			//((x | y) > (x + y)) && willSumOverflow(x, y), // TODO: find why does not work
			absDiff(x, y) <= (x ^ y),
		}
		for i, q := range v {
			if !q {
				t.Error(i, x, y)
			}
		}
	})
}
