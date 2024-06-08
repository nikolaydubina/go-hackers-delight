package hd

// TurnOffRightMostBit this can be used to test if integer is power of 2
func TurnOffRightMostBit[T Signed](x T) T { return x & (x - 1) }

func TurnOnRightMostBit[T Signed](x T) T { return x | (x + 1) }

// TurnOffTrailingOnes this can be used to test if integer if of the form 2^n - 1
func TurnOffTrailingOnes[T Signed](x T) T { return x & (x + 1) }

func TurnOnTrailingZeros[T Signed](x T) T { return x | (x - 1) }

func SetBitLastZero[T Signed](x T) T { return ^x & (x + 1) }

func SetTrailingZeros[T Signed](x T) T { return ^x & (x - 1) }

func SetTrailingZeros2[T Signed](x T) T { return ^(x | -x) }

func SetTrailingZeros3[T Signed](x T) T { return (x & -x) - 1 }

func IsolateRightmostOneBit[T Signed](x T) T { return x & -x }

func SetTrailingZerosWithRightMostOne[T Signed](x T) T { return x ^ (x - 1) }

func SetTrailingOnesWithRightMostOne[T Signed](x T) T { return x ^ (x + 1) }

// TurnOffRightmostOnes this can be sued to determine if a nonnegative integer is of the for 2^j - 2^k for some j >= k >= 0
func TurnOffRightmostOnes[T Signed](x T) T { return ((x | (x - 1)) + 1) & x }

func TurnOffRightmostOnes2[T Signed](x T) T { return ((x & -x) + x) & x }
