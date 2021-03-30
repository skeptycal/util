package stringbenchmarks

import (
	"math/rand"
	"time"
)

const (
	TAB   = 0x09 // '\t'
	LF    = 0x0A // '\n'
	VT    = 0x0B // '\v'
	FF    = 0x0C // '\f'
	CR    = 0x0D // '\r'
	SPACE = ' '
	NBSP  = 0x00A0
	NEL   = 0x0085

	defaultSamples = 1<<8 - 1
	maxSamples     = 1<<32 - 1

	numSamples = 1 << 4
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SmallRuneSamples() []rune {
	return []rune{
		'A', '0', '5', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8, 0x1680, 0x2028, 0x3000, 0x1680, 0x200C, 0x2123, 0x3333, 0xFFDF, 0xFFEE,
	}
}

func SmallByteSamples() []byte {
	buf := make([]byte, 0, 256)
	for i := 0; i < 256; i++ {
		if i != 95 { // underscore is not tested here because some "alphanumeric type functions" include it
			buf = append(buf, byte(i))
		}
	}
	return buf
	// return []byte{
	// 	'A', '=', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8, 0xFF,
	// }
}

func SmallByteStringSamples() (list []string) {
	for _, c := range SmallByteSamples() {
		list = append(list, string(c))
	}
	return
}

func SmallRuneStringSamples() (list []string) {
	for _, r := range SmallRuneSamples() {
		list = append(list, string(r))
	}
	return
}

func ByteSamples() []byte {
	n := numSamples
	if n < 2 || n > maxSamples {
		n = defaultSamples
	}
	retval := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		retval = append(retval, byte(rand.Intn(126)))
	}
	retval = append(retval, 0xFF)
	return retval
}

func RuneSamples() []rune {
	n := numSamples
	if n < 2 || n > maxSamples {
		n = defaultSamples
	}
	retval := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		retval = append(retval, rune(rand.Intn(0x3000)))
	}
	return retval
}

func byteStringSamples() (list []string) {
	for _, r := range RuneSamples() {
		list = append(list, string(r))
	}
	return
}

func runeStringSamples() (list []string) {
	for _, r := range RuneSamples() {
		list = append(list, string(r))
	}
	return
}
