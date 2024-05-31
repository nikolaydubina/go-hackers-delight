package hd_test

import "math"

var fuzzUint32 = [...]uint32{
	0,
	1,
	math.MaxInt32,
	math.MaxInt32 / 2,
	math.MaxInt32 - 1,
	math.MaxUint32,
	math.MaxUint32 / 2,
	math.MaxUint32 - 1,
}

var fuzzUint64 = [...]uint64{
	0,
	1,
	math.MaxInt32,
	math.MaxInt32 / 2,
	math.MaxInt32 - 1,
	math.MaxUint32,
	math.MaxUint32 / 2,
	math.MaxUint32 - 1,
	math.MaxUint64,
	math.MaxUint64 / 2,
	math.MaxUint64 - 1,
}

var fuzzInt32 = [...]int32{
	0,
	1,
	-1,
	math.MinInt32,
	math.MinInt32 / 2,
	math.MinInt32 + 1,
	math.MaxInt32,
	math.MaxInt32 / 2,
	math.MaxInt32 - 1,
}
