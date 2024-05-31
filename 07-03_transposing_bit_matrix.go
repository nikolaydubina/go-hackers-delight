package hd

// TransposeMatrix8bx8b (aka transpose8) uses divide and conquer method.
// Each matrix i,j value is encoded in i-th byte and j-th bit of a byte.
// This is 2x fewer calculating RISC instructions than straightforward masking method.
// This is 21 calculating RISC instructions.
// A: m rows * n columns is 1-based coordinates of source submatrix.
// B: n rows * m columns is 1-based coordinates of destination submatrix.
func TransposeMatrix8bx8b(A, B []byte, m, n int) {
	var x uint64
	for i := 0; i <= 7; i++ {
		x = (x << 8) | uint64(A[m*i])
	}

	x = (x & 0xAA55_AA55_AA55_AA55) | ((x & 0x00AA_00AA_00AA_00AA) << 07) | ((x >> 07) & 0x00AA_00AA_00AA_00AA)
	x = (x & 0xCCCC_3333_CCCC_3333) | ((x & 0x0000_CCCC_0000_CCCC) << 14) | ((x >> 14) & 0x0000_CCCC_0000_CCCC)
	x = (x & 0xF0F0_F0F0_0F0F_0F0F) | ((x & 0x0000_0000_F0F0_F0F0) << 28) | ((x >> 28) & 0x0000_0000_F0F0_F0F0)

	for i := 7; i >= 0; i-- {
		B[n*i] = byte(x & 0xFF)
		x = x >> 8
	}
}

// TODO: TransposeMatrix32bx32b is using 64bit registers and runs in 3886 RISC instructions,
// compared to 64 executions of 8x8 transpose that runs in 5760 RISC instructions.
