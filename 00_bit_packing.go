package hd

func IntToNx16b[T int | int32 | int64 | uint | uint32 | uint64](v T) []uint16 {
	if v == 0 {
		return []uint16{0}
	}
	var out []uint16
	for x := uint64(v); x != 0; x >>= 16 {
		out = append(out, uint16((x & 0xFFFF)))
	}
	return out
}

func Uint64FromNx16b(v []uint16) uint64 {
	var out uint64
	for i := len(v) - 1; i >= 0; i-- {
		out <<= 16
		out |= uint64(v[i])
	}
	return out
}

func Int64FromNx16b(v []uint16) int64 { return int64(Uint64FromNx16b(v)) }
