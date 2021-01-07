// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// +build ppc64 s390x mips mips64

package fastinvsqrt

const isBigEndian = true

// Bytes casts b to a []byte
// Ref: func (bigEndian) PutUint32(b []byte, v uint32)
func (b Bits) Bytes() []byte {
	buf := make([]byte, 4)
	_ = buf[3] // early bounds check to guarantee safety of writes below
	buf[0] = byte(b >> 24)
	buf[1] = byte(b >> 16)
	buf[2] = byte(b >> 8)
	buf[3] = byte(b)
	return buf
}
