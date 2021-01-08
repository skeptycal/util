package bits

// Count returns the number of nonzero bits in w.
func Count(w uint64) int {
	// “Software Optimization Guide for AMD64 Processors”, Section 8.6.
	const maxw = 1<<64 - 1
	const bpw = 64

	// Compute the count for each 2-bit group.
	// Example using 16-bit word w = 00,01,10,11,00,01,10,11
	// w - (w>>1) & 01,01,01,01,01,01,01,01 = 00,01,01,10,00,01,01,10
	w -= (w >> 1) & (maxw / 3)

	// Add the count of adjacent 2-bit groups and store in 4-bit groups:
	// w & 0011,0011,0011,0011 + w>>2 & 0011,0011,0011,0011 = 0001,0011,0001,0011
	w = w&(maxw/15*3) + (w>>2)&(maxw/15*3)

	// Add the count of adjacent 4-bit groups and store in 8-bit groups:
	// (w + w>>4) & 00001111,00001111 = 00000100,00000100
	w += w >> 4
	w &= maxw / 255 * 15

	// Add all 8-bit counts with a multiplication and a shift:
	// (w * 00000001,00000001) >> 8 = 00001000
	w *= maxw / 255
	w >>= (bpw/8 - 1) * 8
	return int(w)
}
