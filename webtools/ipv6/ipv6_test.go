package ipv6

import (
	"testing"
)

func Test_ipv6SubHeader_version(t *testing.T) {
	var (
		blankNew  ipv6SubHeader = 0
		firstBit  ipv6SubHeader = 0b10001111111111111111111111111111
		secondBit ipv6SubHeader = 0b01001111111111111111111111111111
		thirdBit  ipv6SubHeader = 0b00101111111111111111111111111111
	)

	tests := []struct {
		name string
		h    ipv6SubHeader
		want uint32
	}{
		// TODO: Add test cases.
		{"blank new", blankNew, 0},
		{"first bit", firstBit, 0b10000000000000000000000000000000},
		{"2nd bit", secondBit, 0b01000000000000000000000000000000},
		{"blank new", thirdBit, 0b00100000000000000000000000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.version(); got != tt.want {
				t.Errorf("ipv6SubHeader.version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ipv6SubHeader_trafficClass(t *testing.T) {
	var (
		blankNew  ipv6SubHeader = 0
		firstBit  ipv6SubHeader = 0b11111000000011111111111111111111
		secondBit ipv6SubHeader = 0b11110100000011111111111111111111
		thirdBit  ipv6SubHeader = 0b11110010000011111111111111111111
	)
	tests := []struct {
		name string
		h    ipv6SubHeader
		want uint32
	}{
		// TODO: Add test cases.
		{"blank new", blankNew, 0},
		{"first bit", firstBit, 0b00001000000000000000000000000000},
		{"2nd bit", secondBit, 0b00000100000000000000000000000000},
		{"blank new", thirdBit, 0b00000010000000000000000000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.trafficClass(); got != tt.want {
				t.Errorf("ipv6SubHeader.trafficClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ipv6SubHeader_flowLabel(t *testing.T) {
	var (
		blankNew  ipv6SubHeader = 0
		firstBit  ipv6SubHeader = 0b11111111111100000000000000000001
		secondBit ipv6SubHeader = 0b11111111111100000000000000000010
		thirdBit  ipv6SubHeader = 0b11111111111100000000000000000100
	)
	tests := []struct {
		name string
		h    ipv6SubHeader
		want uint32
	}{
		// TODO: Add test cases.
		{"blank new", blankNew, 0},
		{"first bit", firstBit, 0b00000000000000000000000000000001},
		{"2nd bit", secondBit, 0b00000000000000000000000000000010},
		{"blank new", thirdBit, 0b00000000000000000000000000000100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.flowLabel(); got != tt.want {
				t.Errorf("ipv6SubHeader.flowLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}
