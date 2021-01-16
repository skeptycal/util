// Package format contains functions that format numeric values.
package format

import (
	"fmt"
	"strings"
)

// NumSpace formats numeric values for readability by adding
// spaces every three digits.
//
// e.g.
//
//    12345678.87654321 e-42
// first, split off any exponent from the mantissa
//
//    12345678.87654321  and  e-42
// next, split off any decimal part from the integer part
//    12345678   and    .87654321
//
// next, add spaces between digits in the integer part
//
//    12 345 678  and    .876 543 21
//
// finally, add any exponent back to the mantissa
func NumSpace(n float64) string {
	sb := strings.Builder{}
	s := fmt.Sprintf("%g", n)
	mantissa := s
	exponent := ""

	eloc := strings.Index(mantissa, "e")
	if eloc > 0 {
		mantissa = s[:eloc]
		exponent = s[eloc:]
	}

	intpart := mantissa
	decpart := ""

	dloc := strings.Index(intpart, ".")
	if dloc > -1 {
		intpart = mantissa[:dloc]
		decpart = mantissa[dloc:]
	}

	rem := len(intpart) % 3

	sb.WriteString(intpart[:rem])
	sb.WriteString(" ")

	for i := rem; i < len(intpart); i++ {
		if i%3 == 0 {
			sb.WriteByte(intpart[i])
			sb.WriteString(" ")
		}

	}
	return sb.String()
}

/* python version
# Take a very large number and pretty print it in triplets of 3 digits, each triplet separated by a space.
def pnum_spc(n): print(' '.join([''.join(list(str(n))[::-1][i:i+3]) for i in range(0, len(str(n)), 3)][::-1]))
# >>> pnum_spc(32 ** 13)
# 36 893 488 147 419 103 232
*/

/* python version
# Print numbers as 32-bit binary numbers w/ spaces giving 4-bit words
def pbin_spc(n): print(' '.join([''.join(list(f'{n:032b}')[::-1][i:i+4][::-1]) for i in range(0, len(f'{n:032b}'), 4)][::-1]))
# >>> "{0:032b}".format(1234)
# '00000000000000000000010011010010'
# >>> pbin_spc(1234)
# 0000 0000 0000 0000 0000 0100 1101 0010
*/
