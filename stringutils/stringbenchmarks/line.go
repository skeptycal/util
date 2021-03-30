package stringbenchmarks

import (
	"bytes"
	"unsafe"
)

type line []byte // todo ?? use [255]byte or similar??

func (l *line) unsafeToStringPtr() string {
	return *(*string)(unsafe.Pointer(&l))
}

func (l *line) unsafeFromStringPtr(s string) []byte {
	l = (*line)(unsafe.Pointer(&s))
	return *l
}

type byteList struct {
	buf *[][]byte
}

func (l byteList) Len() int      { return len(*l.buf) }
func (l byteList) Cap() int      { return cap(*l.buf) }
func (l byteList) Reset(cap int) { *l.buf = make([][]byte, 0, cap+1) }
func (l byteList) Join() []byte  { return bytes.Join(*l.buf, []byte{0x10}) }
func (l byteList) Make(s string) [][]byte {
	l.Reset(len(s))
	*l.buf = bytes.SplitAfter([]byte(s), []byte{0x10})
	return *l.buf
}
func (l byteList) Contains(b []byte) bool {
	// for some reason, bytes.Index is several times faster than
	// bytes.Contains (see Index below for improvement)
	for _, s := range *l.buf {
		if bytes.Index(s, b) < 0 {
			return false
		}
	}
	return true
}
