// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package runebuilder implements a heavily copied version of
// strings.Builder from the Go standard library that has been
// refactored to support runes.
//
package runebuilder

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

// A Builder is used to efficiently build a utf-8 compliant string using Write methods.
// It minimizes memory copying. The zero value is ready to use.
// Do not copy a non-zero Builder.
//
// RuneBuilder implements a heavily copied version of
// strings.Builder from the Go standard library.
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf  []rune
}

// Link provides access to a pointer to the internal buffer.
func (b *Builder) Link() *[]rune {
    return &b.buf
}

func (b *Builder) Put(r rune, i int) {
	b.buf[i] = r
}

// ReSize creates a new buffer of size i and cap 2 * i
// Any data in the buffer is lost.
func (b *Builder) ReSize(i int) {
	b.buf = make([]rune, i, i+i)
}

// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input. noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (b *Builder) copyCheck() {
	if b.addr == nil {
		// This hack works around a failing of Go's escape analysis
		// that was causing b to escape and be heap allocated.
		// See issue 23382.
		// TODO: once issue 7921 is fixed, this should be reverted to
		// just "b.addr = b".
		b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
	} else if b.addr != b {
		panic("strings: illegal use of non-zero Builder copied by value")
	}
}

// String returns the accumulated string.
func (b *Builder) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}

// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
func (b *Builder) Len() int { return len(b.buf) }

// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
func (b *Builder) Cap() int { return cap(b.buf) }

// Reset resets the Builder to be empty.
func (b *Builder) Reset() {
	b.addr = nil
	b.buf = nil
}

// grow copies the buffer to a new, larger buffer so that there are at least n
// bytes of capacity beyond len(b.buf).
func (b *Builder) grow(n int) {
	b.buf = b.buf
	buf := make([]rune, len(b.buf), 2*cap(b.buf)+n)
	copy(buf, b.buf)
    b.buf = buf
    []byte.
}

// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, Grow panics.
func (b *Builder) Grow(n int) {
	b.copyCheck()
	if n < 0 {
		panic("runeBuilder.Grow: negative count")
	}
	// fmt.Printf("(rb.Grow(n) n: %d   cap: %d    len: %d", n, cap(b.buf), len(b.buf))
	if cap(b.buf)-len(b.buf) < n {
		b.grow(n)
	}
}

// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (b *Builder) Write(p []byte) (int, error) {
	b.copyCheck()
	return b.WriteString(string(p))

	//todo - is this the best way?
	// b.buf = append(b.buf, p...)
	// return len(p), nil
}

// WriteByte decodes the byte c to a valid utf8 rune
// and writes it to b's runeBuffer or returns any error.
func (b *Builder) WriteByte(c byte) error {
	b.copyCheck()
	r, size := utf8.DecodeRune([]byte{c})
	if r == utf8.RuneError {
		if size == 0 {
			return fmt.Errorf("invalid rune (byte empty): %v", c)
		}
		return fmt.Errorf("invalid rune encoding: %v", c)
	}
	_, err := b.WriteRune(r)
	return err
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
// It returns the length of r and a nil error.
func (b *Builder) WriteRune(r rune) (int, error) {
	b.copyCheck()
	// fast path for length 1 runes
	// todo - is this really needed? benchmark it
	// if r < utf8.RuneSelf {
	// 	b.buf = append(b.buf, r)
	// 	return 1, nil
	// }
	// todo - append takes care of this anyway ... ??
	// size := len(b.buf)
	// if cap(b.buf)-size < utf8.UTFMax {
	// 	b.grow(utf8.UTFMax + utf8.UTFMax)
	// 	//todo - is this enough? is it better in situations to grow x2
	// 	//todo - and make it half as often??
	// }
	b.buf = append(b.buf, r)
	return utf8.RuneLen(r), nil //todo do we need to do this? benchmark
}

// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
func (b *Builder) WriteString(s string) (int, error) {
	b.copyCheck()
	// b.buf = append(b.buf, s...) //todo - is this the best way?
	for _, r := range s {
		b.buf = append(b.buf, r)
	}
	return len(s), nil
}
