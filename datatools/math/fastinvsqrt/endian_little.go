// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// +build 386 amd64 arm arm64 ppc64le mips64le mipsle riscv64 wasm

package fastinvsqrt

const isBigEndian = false

// Bytes casts b to a []byte
// Ref: func (littleEndian) PutUint32(b []byte, v uint32)
func (b Bits) Bytes() []byte {
	buf := make([]byte, 4)
	_ = buf[3] // early bounds check to guarantee safety of writes below
	buf[0] = byte(b)
	buf[1] = byte(b >> 8)
	buf[2] = byte(b >> 16)
	buf[3] = byte(b >> 24)
	return buf
}
