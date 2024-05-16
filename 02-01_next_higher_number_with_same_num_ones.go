package hd

// NextHigherNumberWithSameNumberOfOnes is useful in bitsets stored in integers to find next set of same cardinality.
// This function utilizes the theorem of Right-to-Left Computability.
func NextHigherNumberWithSameNumberOfOnes(x uint64) uint64 {
	var smallest, ripple, ones uint64 // x = xxx0 11111 00000
	smallest = x & -x                 //     0000 00001 00000
	ripple = x + smallest             //     xxx1 00000 00000
	ones = x ^ ripple                 //     0001 11111 00000
	ones = (ones >> 2) / smallest     //     0000 00000 01111
	return ripple | ones              //     xxx1 00000 01111
}
