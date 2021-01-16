// Package format contains functions that format numeric values.
package format

import (
	"strings"
)

type stringWriter struct {
	strings.Builder
	intpart  string
	decpart  string
	exponent string
}

func (w *stringWriter) space() { w.WriteString(" ") }
func (w *stringWriter) dot()   { w.WriteString(".") }
func (w *stringWriter) exp()   { w.WriteString(" e" + w.exponent) }
func (w *stringWriter) parse(value string) {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)
	mantissa := value

	eloc := strings.Index(value, "e")
	if eloc != -1 {
		mantissa = value[:eloc]
		w.exponent = value[eloc+1:]
	}

	w.intpart = mantissa
	w.decpart = ""

	dloc := strings.Index(mantissa, ".")
	if dloc > 0 {
		w.intpart = mantissa[:dloc]
		w.decpart = mantissa[dloc+1:]
	}

	w.loadString()
}
func (w *stringWriter) loadString() {

	if w.intpart[0] == '=' {
		w.WriteByte(w.intpart[0])
		w.intpart = w.intpart[1:]
	}

	rem := len(w.intpart)%3 + 2

	t := ""
	j := len(w.intpart) - 1
	for i := j; i >= 0; i-- {
		t += string(w.intpart[i])
		// w.WriteByte(w.intpart[i])
		if (i+rem)%3 == 0 && i < len(w.intpart) {
			t += " "
			// w.space()
		}
	}

	// w.WriteByte(w.intpart[len(w.intpart)-1])

	if w.decpart != "" {
		w.dot()

		for i := 0; i < len(w.decpart); i++ {
			w.WriteByte(w.decpart[i])
			if (i+1)%3 == 0 {
				w.space()
			}

		}
	}
	if w.exponent != "" {
		w.exp()
	}
}

func Reverse(s string) string {
	var ret []byte
	b := []byte(s)
	for i := len(b); i > 0; i-- {
		ret = append(ret, b[i-1])
	}
	return string(ret)
}

// Reference: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Reference: https://stackoverflow.com/a/1754209
func Reverse3(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

// Reverse4 combines the best from SO answers
// (Reverse2 and Reverse3)
func Reverse4(s string) string {
	runes := []rune(s)
	n := len(runes)
	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)

}

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
func NumSpace(s string) string {

	sb := &stringWriter{strings.Builder{}, "", "", ""}
	sb.parse(s)

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
