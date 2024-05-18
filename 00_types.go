package hd

type unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type signed interface {
	int | int8 | int16 | int32 | int64
}

type integers interface {
	signed | unsigned
}
