package hd

func RotateLeft(x uint32, n int) uint32 { return (x << n) | (x >> (32 - n)) }

func RotateRight(x uint32, n int) uint32 { return (x >> n) | (x << (32 - n)) }

// TODO: assembly version for double-length shifts
