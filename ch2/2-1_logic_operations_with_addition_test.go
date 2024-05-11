package ch2

import "testing"

func Fuzz_LogicOperationsWithAddition(f *testing.F) {
	f.Fuzz(func(t *testing.T, x, y int32) {
		v := []bool{
			-x == ^x+1,
			-x == ^(x - 1),
			^x == -x-1,
			-^x == x+1,
			^-x == x-1,
			x+y == x-^y-1,
			x+y == (x^y)+2*(x&y),
			x+y == (x|y)+(x&y),
			x+y == 2*(x|y)-(x^y),
			x-y == x+^y+1,
			x-y == (x^y)-2*(^x&y),
			x-y == (x&(^y))-(^x&y),
			x-y == 2*(x&^y)-(x^y),
			x^y == (x|y)-(x&y),
			x&^y == (x|y)-y,
			x&^y == x-(x&y),
			^(x - y) == y-x-1,
			^(x - y) == ^x+y,
			// x === y == (x&y) - (x|y) - 1
			//         == (x&y) + ^(x|y)
			x|y == (x&^y)+y,
			x&y == (^x|y)-^x,
		}
		for i, q := range v {
			if !q {
				t.Error(i, x, y)
			}
		}
	})
}
