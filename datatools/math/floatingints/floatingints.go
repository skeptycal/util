// Reference: (youtube) Floating point bit hacks every programmer should know
package floatingints

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/zhuangsirui/binpacker"
)

type Pack struct {
	buf      *bytes.Buffer
	packer   *binpacker.Packer
	unpacker *binpacker.Unpacker
}

func (p *Pack) Write(b []byte) (c int, err error) {

	for _, w := range b {
		p.packer.PushByte(w)
		c++
	}
	return c, nil
}

func Bpit() error {
	buffer := new(bytes.Buffer)
	packer := binpacker.NewPacker(binary.BigEndian, buffer)

	packer.PushByte(0x01)
	packer.PushBytes([]byte{0x02, 0x03})
	packer.PushUint16(math.MaxUint16)
	return packer.Error()
}

func BpPush() error {
	// You can push data like this
	buffer := new(bytes.Buffer)
	packer := binpacker.NewPacker(binary.BigEndian, buffer)
	packer.PushByte(0x01).PushBytes([]byte{0x02, 0x03}).PushUint16(math.MaxUint16)
	return packer.Error()
}

func BpExample() error {
	buffer := new(bytes.Buffer)
	packer := binpacker.NewPacker(binary.BigEndian, buffer)
	// unpacker := binpacker.NewUnpacker(binary.BigEndian, buffer)

	packer.PushByte(0x01)
	packer.PushUint16(math.MaxUint16)
	return packer.Error()
}

func NewPack() *Pack {
	buffer := new(bytes.Buffer)
	packer := binpacker.NewPacker(binary.BigEndian, buffer)
	unpacker := binpacker.NewUnpacker(binary.BigEndian, buffer)

	return &Pack{
		buf:      buffer,
		packer:   packer,
		unpacker: unpacker,
	}
}
