// Comparison functions are a couple branch-free instructions that store result in the most significant bit.
package hd

func Equal(x, y int32) int32 { return Abs(x-y) - 1 }

func Equal2(x, y int32) int32 { return Abs(int32(uint32(x-y) + 0x80000000)) }

func Equal3(x, y int32) int32 { return int32(LeadingZerosUint32(uint32(x-y))) << 26 }

func Equal4(x, y int32) int32 { return -int32(LeadingZerosUint32(uint32(x-y)) >> 5) }

func Equal5[T Int](x, y T) T { return ^((x - y) | (y - x)) }

func NotEqual(x, y int32) int32 { return NAbs(x - y) }

func NotEqual2(x, y int32) int32 { return int32(LeadingZerosUint32(uint32(x-y))) - 32 }

func NotEqual3[T Int](x, y T) T { return (x - y) | (y - x) }

func Less(x, y int32) int32 { return (x - y) ^ ((x ^ y) & ((x - y) ^ x)) }

func Less2(x, y int32) int32 { return (x & ^y) | (^(x ^ y) & (x - y)) }

func Less3(x, y int32) int32 { return NAbs(DOZ(y, x)) }

// Less4 utilizes the fact that x/2 - y/2 never overflows.
// Stores result in most significant bit.
// Exactly same formula works for uint32. This is because Go uses signed shift right for int32 and unsigned shift right for uint32.
// This takes 6 or seven instructions.
func Less4[T Int](x, y T) T { return (x >> 1) - (y >> 1) - (^x & y & 1) }

func LessOrEqual[T Signed](x, y T) T { return (x | ^y) & ((x ^ y) | ^(y - x)) }

func LessOrEqual2[T Signed](x, y T) T { return (^(x ^ y) >> 1) + (x &^ y) }

func LessUnsigned[T Unsigned](x, y T) T { return (^x & y) | (^(x ^ y) & (x - y)) }

func LessUnsigned2[T Unsigned](x, y T) T { return (^x & y) | ((^x | y) & (x - y)) }

func LessOrEqualUnsigned[T Unsigned](x, y T) T { return (^x | y) & ((x ^ y) | ^(y - x)) }

func EqualZero(x int32) int32 { return Abs(x) - 1 }

func EqualZero2(x int32) int32 { return Abs(int32(uint32(x) + 0x80000000)) }

func EqualZero3(x int32) int32 { return int32(LeadingZerosUint32(uint32(x))) << 26 }

func EqualZero4[T Int](x T) T { return ^(x | -x) }

func EqualZero5[T Int](x T) T { return ^x & (x - 1) }

func NotEqualZero(x int32) int32 { return NAbs(x) }

func NotEqualZero2(x int32) int32 { return int32(LeadingZerosUint32(uint32(x))) - 32 }

func NotEqualZero3[T Int](x T) T { return x | -x }

func NotEqualZero4[T int32 | uint32](x T) T { return ShiftRightUnsigned32(x, 1) - x }

func LessZero[T Signed](x T) T { return x }

func LessOrEqualZero[T Signed](x T) T { return x | (x - 1) }

func LessOrEqualZero2[T Signed](x T) T { return x | ^-x }

func HigherZero(x int32) int32 { return x ^ NAbs(x) }

func HigherZero2[T Signed](x T) T { return (x >> 1) - x }

func HigherZero3[T Signed](x T) T { return -x & ^x }

func HigherEqualZero[T Signed](x T) T { return ^x }

// TODO: comparison to/from unsigned comparison
// TODO: comparison from carry bits
