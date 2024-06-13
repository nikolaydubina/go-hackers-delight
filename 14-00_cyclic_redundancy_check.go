package hd

// CRC32Generator is Cycle Redundancy Check generator polynomial.
// Each bit i corresponds to the 1/0 of coefficient at x^i in polynomial.
// Among others, it is used in following standards:
// ISO 3309 (HDLC), ANSI X3.66 (ADCCP), ISO/IEC/IEEE 802-3 (Ethernet), MPEG-2, PKZIP, Gzip, Bzip2, POSIX cksum, PNG.
var CRC32Generator uint32 = 0x4C11DB7

var CRC32GeneratorReversed uint32 = 0xEDB88320

// CRC32Basic computes CRC 32bit checksum one bit a time.
func CRC32Basic(data []byte) uint32 {
	crc := uint32(0xFFFF_FFFF)
	for _, b := range data {
		b := ReverseBits(uint32(b))
		for j := 0; j < 8; j++ {
			if int32((crc ^ b)) < 0 {
				crc = (crc << 1) ^ CRC32Generator
			} else {
				crc <<= 1
			}
			b <<= 1
		}
	}
	return ReverseBits(^crc)
}

var crc32TableLookup [256]uint32

func init() {
	crc32TableLookup = NewCRC32TableLookupTable()
}

func NewCRC32TableLookupTable() [256]uint32 {
	table := [256]uint32{}
	for i := range table {
		crc := uint32(i)
		for j := 7; j >= 0; j-- {
			mask := -(crc & 1)
			crc = (crc >> 1) ^ (CRC32GeneratorReversed & mask)
		}
		table[i] = crc
	}
	return table
}

// CRC32TableLookup uses precomputed values.
func CRC32TableLookup(data []byte) uint32 {
	crc := uint32(0xFFFF_FFFF)
	for _, b := range data {
		crc = (crc >> 8) ^ crc32TableLookup[byte(crc)^b]
	}
	return ^crc
}
