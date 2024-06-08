package hd

// IsInRange illustrates how to perform signed integer bounds checking (both ends inclusive) in single comparison.
// This is useful for array bounds checking.
// Using uint conversion, to force unsigned comparison on signed integers, given Go lacks unsigned comparison operator on signed integers.
func IsInRange[T Integer](x, a, b T) bool { return uint(x-a) <= uint(b-a) }

// IsInRangeClosedOpen: a <= x < b
func IsInRangeClosedOpen[T Integer](x, a, b T) bool { return uint(x-a) < uint(b-a) }

// IsInRangeOpenClosed: a < x <= b
func IsInRangeOpenClosed[T Integer](x, a, b T) bool { return uint(b-x) < uint(b-a) }

// IsInRangeOpen: a < x < b
func IsInRangeOpen[T Integer](x, a, b T) bool { return uint(x-a-1) < uint(b-a-1) }

func IsInRangeOpen2[T Integer](x, a, b T) bool { return uint(b-x-1) < uint(b-a-1) }

// IsInRangePowerTwo: 0 <= x <= 2^n - 1
func IsInRangePowerTwo[T Unsigned](x T, n int) bool { return (x >> n) == 0 }

// IsInRangePowerTwoOffset: a <= x <= a + 2^n - 1
func IsInRangePowerTwoOffset[T Unsigned](x, a T, n int) bool { return ((x - a) >> n) == 0 }
