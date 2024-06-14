package hd

import "unsafe"

// SqrtFloat32ErrorRate is error rate (abs(true_value - result) / result) that RSqrtFloat32 produces.
const SqrtFloat32ErrorRate float32 = 0.035

// RSqrtFloat32 (rsqrt) algorithm is fast floating point approximation of reciprocal square root (1 / sqrt).
// It does not work +/- Inf, NaN, denormals, negative numbers.
// This algorithm caused some buzz in early 2000s.
// It produces result with 3.5% error rate.
func RSqrtFloat32(x float32) float32 {
	// C uses union here, in Go we get to access this memory through unsafe.Pointer
	p := unsafe.Pointer(&x)
	xi, xf := (*int32)(p), (*float32)(p)

	xhalf := 0.5 * x
	*xi = 0x5F37_5A82 - (*xi >> 1)         // Initial guess.
	*xf *= 1.5008908 - (xhalf * *xf * *xf) // Newton step.
	return *xf
}
