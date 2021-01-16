// Package format contains functions that format numeric values.
package format

import (
	"fmt"
	"strings"
)

func NumSpace(n float64) string {
	sb := strings.Builder{}
	mantissa := fmt.Sprintf("%g", n)
	exponent := ""

	parts := strings.Split(mantissa, "e")
	if len(parts) > 1 {
		mantissa = parts[0]
		exponent = parts[1]
	}

	intpart := mantissa
	decpart := ""

	parts = strings.Split(intpart, ".")
	if len(parts) > 1 {
		intpart = parts[0]
		decpart = parts[1]
	}

	for i := 0; i < len(intpart); i++ {
		if i%3 == 0 {
			sb.WriteString(" ")
			sb.WriteString(intpart[i])
		}

	}
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
