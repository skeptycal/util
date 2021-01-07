package fastinvsqrt

import "math/big"

// BitMask32 constants are used to mask bit operations in 32 bit numbers.
/*
   MantissaBitMask  BitMask32 = 1<<23 - 1
   ExpBitMask       BitMask32 = (0xFF) << 23
   SignBitMask      BitMask32 = 1 << 31
   All32BitMask     BitMask32 = 1<<32 - 1

Notes:

   ZeroBitMask      BitMask32 = 0b00000000000000000000000000000000
   MantissaBitMask  BitMask32 = 0b00000000011111111111111111111111
   ExpBitMask       BitMask32 = 0b01111111100000000000000000000000
   SignBitMask      BitMask32 = 0b10000000000000000000000000000000
   All32BitMask     BitMask32 = 0b11111111111111111111111111111111

   MantissaBitMask  0x00 7F FF FF     -or-         8 388 607
   ExpBitMask       0x7F 80 00 00     -or-     2 139 095 040
   SignBitMask      0x80 00 00 00     -or-     2 147 483 648
   All32BitMask     0xFF FF FF FF     -or-     4 294 967 295

*/
type BitMask32 = uint32

const (
	MantissaBitMask BitMask32 = 1<<23 - 1
	ExpBitMask      BitMask32 = (0xFF) << 23
	SignBitMask     BitMask32 = 1 << 31
	All32BitMask    BitMask32 = 1<<32 - 1
)

type BitMask64 = uint64

const (
	MantissaBitMask64 BitMask64 = 1<<23 - 1
	ExpBitMask64      BitMask64 = (0xFF) << 23
	SignBitMask64     BitMask64 = 1 << 63
	All32BitMask64    BitMask64 = 1<<64 - 1
)

type biggie big.Int

// SignBit represents the sign bit of a floating point number.
//
// Normal values are 0 or 1;
// a value of -1 indicates an error, infinity, or NaN.
type SignBit = int8

const (
	SignError SignBit = iota - 1 // error, infinity, or NaN
	Positive                     // sign bit off; number >= 0
	Negative                     // sign bit on: number < 0
)

const (
	// arbitrary 1 ppt tolerance // TODO - move to config
	DefaultTolerance = OnePPT
)

// ToleranceType constants describe allowable
// tolerances for estimates
type ToleranceType = float64

// func (t *tolerance) String() string {
// 	return fmt.Sprintf("%.3F", t)
// }

const (
	FivePercent ToleranceType = 5e-2
	TwoPercent  ToleranceType = 2e-2
	OnePercent  ToleranceType = 1e-2
	OnePPT      ToleranceType = 1e-3
	OnePPM      ToleranceType = 1e-6
	OnePPB      ToleranceType = 1e-9
)
