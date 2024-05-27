package hd

// MinOR is min(x|y) for x in [minx, maxx] and y in [miny, maxy].
// This gives tighter bound than max(minx, miny).
func MinOR(minx, maxx, miny, maxy uint32) uint32 {
	var temp uint32
	for m := uint32(0x8000_0000); m != 0; m >>= 1 {
		if (^minx & miny & m) != 0 {
			if temp = (minx | m) & -m; temp <= maxx {
				minx = temp
				break
			}
		} else {
			if (minx & ^miny & m) != 0 {
				if temp = (miny | m) & -m; temp <= maxy {
					miny = temp
					break
				}
			}
		}
	}
	return minx | miny
}

// MaxOR is max(x|y) for x in [minx, maxx] and y in [miny, maxy].
// This gives tighter bound than maxx + maxy.
func MaxOR(minx, maxx, miny, maxy uint32) uint32 {
	var temp uint32
	for m := uint32(0x8000_0000); m != 0; m >>= 1 {
		if (maxx & maxy & m) != 0 {
			if temp = (maxx - m) | (m - 1); temp >= minx {
				maxx = temp
				break
			}
			if temp = (maxy - m) | (m - 1); temp >= miny {
				maxy = temp
				break
			}
		}
	}
	return maxx | maxy
}

// MinAND is min(x&y) for x in [minx, maxx] and y in [miny, maxy].
// This gives tighter bound than 0.
func MinAND(minx, maxx, miny, maxy uint32) uint32 {
	var temp uint32
	for m := uint32(0x8000_0000); m != 0; m >>= 1 {
		if (^minx & ^miny & m) != 0 {
			if temp = (minx | m) & -m; temp <= maxx {
				minx = temp
				break
			}
			if temp = (miny | m) & -m; temp <= maxy {
				miny = temp
				break
			}
		}
	}
	return minx & miny
}

// MaxAND is max(x|y) for x in [minx, maxx] and y in [miny, maxy].
// This gives tighter bound than min(maxx, maxy).
func MaxAND(minx, maxx, miny, maxy uint32) uint32 {
	var temp uint32
	for m := uint32(0x8000_0000); m != 0; m >>= 1 {
		if (maxx & ^maxy & m) != 0 {
			if temp = (maxx & ^m) | (m - 1); temp >= minx {
				maxx = temp
				break
			}
		} else {
			if (^maxx & maxy & m) != 0 {
				if temp = (maxy & ^m) | (m - 1); temp >= miny {
					maxy = temp
					break
				}
			}
		}
	}
	return maxx & maxy
}
