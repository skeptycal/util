package gofile

import "bytes"

// This file contains the portions of code that are largely modified variations
//  or unexported parts of the original standard library code (from Go 1.15.5)

const (
	minReadBufferSize        = 16
	smallBufferSize          = 64
	defaultBufSize           = 4096
	maxConsecutiveEmptyReads = 100
	maxInt                   = int(^uint(0) >> 1)
	chunk                    = bytes.MinRead
)

// The readOp constants describe the last action performed on
// the buffer, so that UnreadRune and UnreadByte can check for
// invalid usage. opReadRuneX constants are chosen such that
// converted to int they correspond to the rune size that was read.
// (from Go 1.15.5 bytes/buffer.go)
type readOp int8

// Don't use iota for these, as the values need to correspond with the
// names and comments, which is easier to see when being explicit.
// (from Go 1.15.5 bytes/buffer.go)
const (
	opRead      readOp = -1 // Any other read operation.
	opInvalid   readOp = 0  // Non-read operation.
	opReadRune1 readOp = 1  // Read rune of size 1.
	opReadRune2 readOp = 2  // Read rune of size 2.
	opReadRune3 readOp = 3  // Read rune of size 3.
	opReadRune4 readOp = 4  // Read rune of size 4.
)
