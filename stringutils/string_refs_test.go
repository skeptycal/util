package stringutils

import (
	"bytes"
	"errors"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var strLen int = 100
var numTrials int = 100

// testString is the string that is repeated during benchmarks.
var testString string

func BenchmarkRandomOrder(b *testing.B) {
	tests := map[int]struct {
		name string
		f    func(b *testing.B)
	}{
		1: {"concat +", BenchmarkConcatString},
		2: {"strings.Builder", BenchmarkConcatBuilder},
		3: {"bytes.Buffer", BenchmarkConcatBuffer},
		4: {"streamlined Buffer", BenchmarkConcatBufferImplementation},
	}
	testString = "x"
	rand.Seed(time.Now().UnixNano())

	var a = make([]int, 0, len(tests))

	for i := 1; i < len(tests); i++ {
		a = append(a, i)
	}
	for j := 0; j < numTrials; j++ {

		// Randomize order of tests for each run.
		// Reference: https://yourbasic.org/golang/shuffle-slice-array/
		for i := len(a) - 1; i > 0; i-- { // Fisher–Yates shuffle
			j := rand.Intn(i + 1)
			a[i], a[j] = a[j], a[i]
		}

		for _, n := range a {
			tt := tests[n]
			b.ResetTimer()
			b.Run(tt.name, func(b *testing.B) {

				for n := 0; n < b.N; n++ {
					tt.f(b)
				}

			})
		}
	}
}

///-----> Streamlined implementation of bytes.Buffer for speed benchmarking.
///-----> Only WriteString and supporting functionality is implemented.

// Benchmarks:
/*
Initial:
BenchmarkConcatBufferImplementation-8   	124117700	         9.78 ns/op	       2 B/op	       0 allocs/op
BenchmarkConcatString-8                 	 7810017	       152 ns/op	     530 B/op	       0 allocs/op
BenchmarkConcatBuffer-8                 	125407339	         9.60 ns/op	       2 B/op	       0 allocs/op
BenchmarkConcatBuilder-8                	411729237	         2.86 ns/op	       2 B/op	       0 allocs/op
*/

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type bytesBuffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}

// makeSlice allocates a slice of size n. If the allocation fails, it panics
// with ErrTooLarge.
func makeSlice(n int) []byte {
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {
			panic(ErrTooLarge)
		}
	}()
	return make([]byte, n)
}

// smallBufferSize is an initial allocation minimal capacity.
const smallBufferSize = 64

// ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
var ErrTooLarge = errors.New("bytes.Buffer: too large")
var errNegativeRead = errors.New("bytes.Buffer: reader returned negative count from Read")

const maxInt = int(^uint(0) >> 1)

// WriteString appends the contents of s to the buffer, growing the buffer as
// needed. The return value n is the length of s; err is always nil. If the
// buffer becomes too large, WriteString will panic with ErrTooLarge.
func (b *bytesBuffer) WriteString(s string) (n int, err error) {
	/// removed for performance implementation
	b.lastRead = opInvalid
	/// removed for performance implementation ... the buffer is
	/// pre-sized and will never grow ...
	m, ok := b.tryGrowByReslice(len(s))
	if !ok {
		m = b.grow(len(s))
	}
	return copy(b.buf[m:], s), nil
}

// Len returns the number of bytes of the unread portion of the buffer;
// b.Len() == len(b.Bytes()).
func (b *bytesBuffer) Len() int { return len(b.buf) - b.off }

// tryGrowByReslice is a inlineable version of grow for the fast-case where the
// internal buffer only needs to be resliced.
// It returns the index where bytes should be written and whether it succeeded.
func (b *bytesBuffer) tryGrowByReslice(n int) (int, bool) {
	if l := len(b.buf); n <= cap(b.buf)-l {
		b.buf = b.buf[:l+n]
		return l, true
	}
	return 0, false
}

// Reset resets the buffer to be empty,
// but it retains the underlying storage for use by future writes.
// Reset is the same as Truncate(0).
func (b *bytesBuffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
	b.lastRead = opInvalid
}

// grow grows the buffer to guarantee space for n more bytes.
// It returns the index where bytes should be written.
// If the buffer can't grow it will panic with ErrTooLarge.
func (b *bytesBuffer) grow(n int) int {
	m := b.Len()
	// If buffer is empty, reset to recover space.
	if m == 0 && b.off != 0 {
		b.Reset()
	}
	// Try to grow by means of a reslice.
	if i, ok := b.tryGrowByReslice(n); ok {
		return i
	}
	if b.buf == nil && n <= smallBufferSize {
		b.buf = make([]byte, n, smallBufferSize)
		return 0
	}
	c := cap(b.buf)
	if n <= c/2-m {
		// We can slide things down instead of allocating a new
		// slice. We only need m+n <= c to slide, but
		// we instead let capacity get twice as large so we
		// don't spend all our time copying.
		copy(b.buf, b.buf[b.off:])
	} else if c > maxInt-c-n {
		panic(ErrTooLarge)
	} else {
		// Not enough space anywhere, we need to allocate.
		buf := makeSlice(2*c + n)
		copy(buf, b.buf[b.off:])
		b.buf = buf
	}
	// Restore b.off and len(b.buf).
	b.off = 0
	b.buf = b.buf[:m+n]
	return m
}

func BenchmarkConcatBufferImplementation(b *testing.B) {
	var buffer bytesBuffer

	for i := 0; i < strLen; i++ {
		buffer.WriteString(testString)
	}

	buffer = bytesBuffer{}
}

///-----> end of bytesBuffer

// String Concatenation Reference Benchmarks
// Reference: https://www.instana.com/blog/practical-golang-benchmarks/

func BenchmarkConcatString(b *testing.B) {
	var str string

	for i := 0; i < strLen; i++ {
		str += testString
	}

	str = ""
}

func BenchmarkConcatBuffer(b *testing.B) {
	var buffer bytes.Buffer

	for i := 0; i < strLen; i++ {
		buffer.WriteString(testString)
	}

	buffer = bytes.Buffer{}
}

func BenchmarkConcatBuilder(b *testing.B) {
	var builder strings.Builder

	for i := 0; i < strLen; i++ {
		builder.WriteString(testString)
	}

	builder = strings.Builder{}
}

// The readOp constants describe the last action performed on
// the buffer, so that UnreadRune and UnreadByte can check for
// invalid usage. opReadRuneX constants are chosen such that
// converted to int they correspond to the rune size that was read.
type readOp int8

// Don't use iota for these, as the values need to correspond with the
// names and comments, which is easier to see when being explicit.
const (
	opRead      readOp = -1 // Any other read operation.
	opInvalid   readOp = 0  // Non-read operation.
	opReadRune1 readOp = 1  // Read rune of size 1.
	opReadRune2 readOp = 2  // Read rune of size 2.
	opReadRune3 readOp = 3  // Read rune of size 3.
	opReadRune4 readOp = 4  // Read rune of size 4.
)
