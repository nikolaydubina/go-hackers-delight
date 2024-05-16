package hd_test

import "testing"

func FuzzLogicOperationsWithAddition(f *testing.F) {
	f.Fuzz(func(t *testing.T, x, y int32) {
		v := []bool{
			-x == ^x+1,                          //
			-x == ^(x - 1),                      //
			^x == -x-1,                          //
			-^x == x+1,                          // addition can be implemented with not and negation
			^-x == x-1,                          //
			x+y == (x - ^y - 1),                 //
			x+y == ((x ^ y) + (2 * (x & y))),    // this has carry probability of 0.25, when normally it is 0.5
			x+y == ((x | y) + (x & y)),          //
			x+y == ((2 * (x | y)) - (x ^ y)),    //
			x-y == x+^y+1,                       //
			x-y == ((x ^ y) - (2 * (^x & y))),   //
			x-y == ((x & (^y)) - (^x & y)),      //
			x-y == ((2 * (x &^ y)) - (x ^ y)),   //
			x^y == ((x | y) - (x & y)),          // shows how to do XOR in RISC with only three basic instructions
			x&^y == ((x | y) - y),               //
			x&^y == (x - (x & y)),               //
			^(x - y) == (y - x - 1),             //
			^(x - y) == ^x+y,                    //
			^(x ^ y) == ((x & y) - (x | y) - 1), // bitwise equality (===)
			^(x ^ y) == ((x & y) + ^(x | y)),    // bitwise equality (===)
			x|y == ((x &^ y) + y),               //
			x&y == ((^x | y) - ^x),              //
		}
		for i, q := range v {
			if !q {
				t.Error(i, x, y)
			}
		}
	})
}
