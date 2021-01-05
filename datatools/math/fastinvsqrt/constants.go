package fastinvsqrt

const (
	zeroBitMask     uint32 = 0b00000000000000000000000000000000 // 0
	signBitMask     uint32 = 0b10000000000000000000000000000000 // 0x80000000
	expBitMask      uint32 = 0b01111111100000000000000000000000 // 0x7F800000
	mantissaBitMask uint32 = 0b00000000011111111111111111111111 // 0x7FFFFF
	all32BitMask    uint32 = 0b11111111111111111111111111111111 // 0xFFFFFFFF (2^32)
)

// signBit represents the sign bit: 0 or 1
// a value of -1 indicates an error
type signBit int8

const (
	signError signBit = iota - 1
	positive          // >= 0
	negative          // < 0
)
