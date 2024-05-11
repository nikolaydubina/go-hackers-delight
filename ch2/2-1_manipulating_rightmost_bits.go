package ch2

// TurnOffRightMostBit this can be used to test if integer is power of 2
func TurnOffRightMostBit(x int64) int64 { return x & (x - 1) }

func TurnOnRightMostBit(x int64) int64 { return x | (x + 1) }

// TurnOffTrailingOnes this can be used to test if integer if of the form 2^n - 1
func TurnOffTrailingOnes(x int64) int64 { return x & (x + 1) }

func TurnOnTrailingZeros(x int64) int64 { return x | (x - 1) }

func SetBitLastZero(x int64) int64 { return ^x & (x + 1) }

//func SetZeroBitLastOne(x int64) int64 { return ^x | (x - 1) }

func SetTrailingZeroes(x int64) int64 { return ^x & (x - 1) }

func SetTrailingZeroes2(x int64) int64 { return ^(x | -x) }

func SetTrailingZeroes3(x int64) int64 { return (x & -x) - 1 }

//func SetTrailingOnes(x int64) int64 { return ^x | (x + 1) }

func IsolateRightmostOneBit(x int64) int64 { return x & -x }

func SetTrailingZeroesWithRightMostOne(x int64) int64 { return x ^ (x - 1) }

func SetTrailingOnesWithRightMostOne(x int64) int64 { return x ^ (x + 1) }

// TurnOffRightmostOnes this can be sued to determine if a nonnegative integer is of the for 2^j - 2^k for some j >= k >= 0
func TurnOffRightmostOnes(x int64) int64 { return (((x | (x - 1)) + 1) & x) }

func TurnOffRightmostOnes2(x int64) int64 { return ((x & -x) + x) & x }
