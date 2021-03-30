package stringbenchmarks

import (
	"bytes"
	"unicode/utf8"

	"golang.org/x/sys/cpu"
)

const MaxBruteForce = 64 // x86 values

var MaxLen int = 0

func init() {
	if cpu.X86.HasAVX2 { // x86 values
		MaxLen = 63
	} else {
		MaxLen = 31
	}
}

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func Cutover(n int) int {
	// 1 error per 8 characters, plus a few slop to start.
	return (n + 16) / 8
}

func bytealgCount(b []byte, c byte) int {
	n := 0
	for _, x := range b {
		if x == c {
			n++
		}
	}
	return n
}

// Count counts the number of non-overlapping instances of sep in s.
// If sep is an empty slice, Count returns 1 + the number of UTF-8-encoded code points in s.
func bCount(s, sep []byte) int {
	// special case
	if len(sep) == 0 {
		return utf8.RuneCount(s) + 1
	}
	if len(sep) == 1 {
		return bytealgCount(s, sep[0])
	}
	n := 0
	for {
		i := bIndex(s, sep)
		if i == -1 {
			return n
		}
		n++
		s = s[i+len(sep):]
	}
}

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func bIndex(s, sep []byte) int {
	n := len(sep)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return bytes.IndexByte(s, sep[0])
	case n == len(s):
		if Equal(sep, s) {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	case n <= MaxLen:
		// Use brute force when s and sep both are small
		if len(s) <= MaxBruteForce {
			return bytes.Index(s, sep)
		}
		c0 := sep[0]
		c1 := sep[1]
		i := 0
		t := len(s) - n + 1
		fails := 0
		for i < t {
			if s[i] != c0 {
				// IndexByte is faster than bytealg.Index, so use it as long as
				// we're not getting lots of false positives.
				o := bytes.IndexByte(s[i+1:t], c0)
				if o < 0 {
					return -1
				}
				i += o + 1
			}
			if s[i+1] == c1 && Equal(s[i:i+n], sep) {
				return i
			}
			fails++
			i++
			// Switch to bytealg.Index when IndexByte produces too many false positives.
			if fails > Cutover(i) {
				r := bIndex(s[i:], sep)
				if r >= 0 {
					return r + i
				}
				return -1
			}
		}
		return -1
	}
	c0 := sep[0]
	c1 := sep[1]
	i := 0
	fails := 0
	t := len(s) - n + 1
	for i < t {
		if s[i] != c0 {
			o := bytes.IndexByte(s[i+1:t], c0)
			if o < 0 {
				break
			}
			i += o + 1
		}
		if s[i+1] == c1 && Equal(s[i:i+n], sep) {
			return i
		}
		i++
		fails++
		if fails >= 4+i>>4 && i < t {
			// Give up on IndexByte, it isn't skipping ahead
			// far enough to be better than Rabin-Karp.
			// Experiments (using IndexPeriodic) suggest
			// the cutover is about 16 byte skips.
			// TODO: if large prefixes of sep are matching
			// we should cutover at even larger average skips,
			// because Equal becomes that much more expensive.
			// This code does not take that effect into account.
			j := IndexRabinKarpBytes(s[i:], sep)
			if j < 0 {
				return -1
			}
			return i + j
		}
	}
	return -1
}

// Equal reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func bEqual(a, b []byte) int {
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	if string(a) == string(b) {
		return 1
	}
	return 0
}

// Equal reports whether a and b
// are the same length and contain the same bytes.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []byte) bool {
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	return string(a) == string(b)
}

// Contains reports whether subslice is within b.
func bContains(b, subslice []byte) int {
	if bIndex(b, subslice) != -1 {
		return 1
	}
	return 0
}

// IndexRabinKarpBytes uses the Rabin-Karp search algorithm to return the index of the
// first occurrence of substr in s, or -1 if not present.
func IndexRabinKarpBytes(s, sep []byte) int {
	// Rabin-Karp search
	hashsep, pow := HashStrBytes(sep)
	n := len(sep)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*PrimeRK + uint32(s[i])
	}
	if h == hashsep && Equal(s[:n], sep) {
		return 0
	}
	for i := n; i < len(s); {
		h *= PrimeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashsep && Equal(s[i-n:i], sep) {
			return i - n
		}
	}
	return -1
}

// PrimeRK is the prime base used in Rabin-Karp algorithm.
const PrimeRK = 16777619

// HashStrBytes returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func HashStrBytes(sep []byte) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*PrimeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
