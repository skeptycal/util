package ipv6

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
)

const ( // 4 + 8 + 20 = 32
	versionBitMask      = 0xF0000000 // 0b11110000000000000000000000000000
	trafficClassBitMask = 0xFF00000  // 0b00001111111100000000000000000000
	flowLabelBitMask    = 0xFFFFF    // 0b00000000000011111111111111111111
)

const (
	Version   = 6  // protocol version
	HeaderLen = 40 // header length
)

type ipv6SubHeader uint32

func (h ipv6SubHeader) version() uint32 {
	return uint32(h & versionBitMask)
}

func (h ipv6SubHeader) trafficClass() uint32 {
	return uint32(h & trafficClassBitMask)
}

func (h ipv6SubHeader) flowLabel() uint32 {
	return uint32(h & flowLabelBitMask)
}

// ipv6Header defines the IPv6 Header Format from
// Reference: https://tools.ietf.org/html/rfc2460#section-3
// Apple uses https://tools.ietf.org/html/rfc3542
//      version               should be 4 bit
//      trafficClass       uint8
//      flowLabel           should be 20 bit
//      payloadLength   uint16
// Go cannot easily access anything less than a byte, but since
// 4 + 8 + 20 = 32 bits, the (version, traffic class, and flow label)
// are wrapped in an ipv6SubHeader of type uint32 is used with
// bitmap getters and setters to keep the size correct.
type ipv6Header struct {
	subHeader     ipv6SubHeader
	payloadLength uint16
	next          uint8
	hopLimit      uint8
	source        uint // 128 bit
	destination   uint // 128 bit

}

// A Header represents an IPv6 base header.
type Header struct {
	Version      int    // protocol version
	TrafficClass int    // traffic class
	FlowLabel    int    // flow label
	PayloadLen   int    // payload length
	NextHeader   int    // next header
	HopLimit     int    // hop limit
	Src          net.IP // source address
	Dst          net.IP // destination address
}

func (h *Header) String() string {
	if h == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ver=%d tclass=%#x flowlbl=%#x payloadlen=%d nxthdr=%d hoplim=%d src=%v dst=%v", h.Version, h.TrafficClass, h.FlowLabel, h.PayloadLen, h.NextHeader, h.HopLimit, h.Src, h.Dst)
}

// ParseHeader parses b as an IPv6 base header.
func ParseHeader(b []byte) (*Header, error) {
	if len(b) < HeaderLen {
		return nil, http.MethodHeaderrHeaderTooShort
	}
	h := &Header{
		Version:      int(b[0]) >> 4,
		TrafficClass: int(b[0]&0x0f)<<4 | int(b[1])>>4,
		FlowLabel:    int(b[1]&0x0f)<<16 | int(b[2])<<8 | int(b[3]),
		PayloadLen:   int(binary.BigEndian.Uint16(b[4:6])),
		NextHeader:   int(b[6]),
		HopLimit:     int(b[7]),
	}
	h.Src = make(net.IP, net.IPv6len)
	copy(h.Src, b[8:24])
	h.Dst = make(net.IP, net.IPv6len)
	copy(h.Dst, b[24:40])
	return h, nil
}

const IPv6HeaderStandardText = `IPv6 Header Format

   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |Version| Traffic Class |           Flow Label                  |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |         Payload Length        |  Next Header  |   Hop Limit   |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                                                               |
   +                                                               +
   |                                                               |
   +                         Source Address                        +
   |                                                               |
   +                                                               +
   |                                                               |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
   |                                                               |
   +                                                               +
   |                                                               |
   +                      Destination Address                      +
   |                                                               |
   +                                                               +
   |                                                               |
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

   Version              4-bit Internet Protocol version number = 6.

   Traffic Class        8-bit traffic class field.  See section 7.

   Flow Label           20-bit flow label.  See section 6.

   Payload Length       16-bit unsigned integer.  Length of the IPv6
                        payload, i.e., the rest of the packet following
                        this IPv6 header, in octets.  (Note that any`
