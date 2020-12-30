// package IEEE754
/* `IEEE 754-2019`

The IEEE Standard for Floating-Point Arithmetic (IEEE 754) is a technical standard for floating-point arithmetic established in 1985 by the Institute of Electrical and Electronics Engineers (IEEE). The standard addressed many problems found in the diverse floating-point implementations that made them difficult to use reliably and portably. Many hardware floating-point units use the IEEE 754 standard.

------------------------------------------------------------------------------

The standard defines:

- arithmetic formats: sets of binary and decimal floating-point data, which consist of finite numbers (including signed zeros and subnormal numbers), infinities, and special "not a number" values (NaNs)

- interchange formats: encodings (bit strings) that may be used to exchange floating-point data in an efficient and compact form

- rounding rules: properties to be satisfied when rounding numbers during arithmetic and conversions

- operations: arithmetic and other operations (such as trigonometric functions) on arithmetic formats

- exception handling: indications of exceptional conditions (such as division by zero, overflow, etc.)

IEEE 754-2008, published in August 2008, includes nearly all of the original IEEE 754-1985 standard, plus the IEEE 854-1987 Standard for Radix-Independent Floating-Point Arithmetic. The current version, IEEE 754-2019, was published in July 2019. It is a minor revision of the previous version, incorporating mainly clarifications, defect fixes and new recommended operations.

Reference: https://en.wikipedia.org/wiki/IEEE_754#2019

*/
package IEEE754

// Standard development
/* `
The first standard for floating-point arithmetic, IEEE 754-1985, was published in 1985. It covered only binary floating-point arithmetic.

A new version, IEEE 754-2008, was published in August 2008, following a seven-year revision process, chaired by Dan Zuras and edited by Mike Cowlishaw. It replaced both IEEE 754-1985 (binary floating-point arithmetic) and IEEE 854-1987 Standard for Radix-Independent Floating-Point Arithmetic. The binary formats in the original standard are included in this new standard along with three new basic formats, one binary and two decimal.

    To conform to the current standard, an implementation must implement at least one of the basic formats as both an arithmetic format and an interchange format.

The international standard ISO/IEC/IEEE 60559:2011 (with content identical to IEEE 754-2008) has been approved for adoption through JTC1/SC 25 under the ISO/IEEE PSDO Agreement and published.

The current version, IEEE 754-2019 published in July 2019, is derived from and replaces IEEE 754-2008, following a revision process started in September 2015, chaired by David G. Hough and edited by Mike Cowlishaw. It incorporates mainly clarifications (e.g. totalOrder) and defect fixes (e.g. minNum), but also includes some new recommended operations (e.g. augmentedAddition).

The international standard ISO/IEC 60559:2020 (with content identical to IEEE 754-2019) has been approved for adoption through JTC1/SC 25 and published.
*/

// Formats
/*
An IEEE 754 format is a "set of representations of numerical values and symbols". A format may also include how the set is encoded.(To conform to the current standard, an implementation must implement at least one of the basic formats as both an arithmetic format and an interchange format.)

A floating-point format is specified by:

- a base (also called radix) b, which is either 2 (binary) or 10 (decimal) in IEEE 754;

- a precision p;

- an exponent range from emin to emax, with emin = 1 − emax for all IEEE 754 formats.

A format comprises:

Finite numbers, which can be described by three integers:

- s = a sign (zero or one)

- c = a significand (or coefficient) having no more than p digits when written in base b (i.e., an integer in the range through 0 to bp − 1)

- q = an exponent such that emin ≤ q + p − 1 ≤ emax. The numerical value of such a finite number is

    (−1)s × c × bq.

Moreover, there are two zero values, called signed zeros: the sign bit specifies whether a zero is +0 (positive zero) or −0 (negative zero).

- Two infinities: +∞ and −∞.

- Two kinds of NaN (not-a-number): a quiet NaN (qNaN) and a signaling NaN (sNaN).

For example, if b = 10, p = 7, and emax = 96, then emin = −95, the significand satisfies 0 ≤ c ≤ 9999999, and the exponent satisfies −101 ≤ q ≤ 90. Consequently, the smallest non-zero positive number that can be represented is 1×10−101, and the largest is 9999999×1090 (9.999999×1096), so the full range of numbers is −9.999999×1096 through 9.999999×1096. The numbers −b1−emax and b1−emax (here, −1×10−95 and 1×10−95) are the smallest (in magnitude) normal numbers; non-zero numbers between these smallest numbers are called subnormal numbers.
*/

type scinot struct {
	// The significand (also mantissa or coefficient, sometimes also argument, or ambiguously fraction or characteristic) is part of a number in scientific notation or a floating-point number, consisting of its significant digits.
	// Depending on the interpretation of the exponent, the significand may represent an integer or a fraction.
	// Reference: https://en.wikipedia.org/wiki/Significand

	mantissa  int
	exponent  int
	precision int8
}

// Finally, the value can be represented in the format given by the Language Independent Arithmetic standard and several programming language standards, including Ada, C, Fortran and Modula-2, as
//      123.45 = 0.12345 × 10+3
// Schmid called this representation with a significand ranging between 0.1 and 1.0 the true normalized form.
//
// Reference: https://en.wikipedia.org/wiki/ISO/IEC_10967
func (s scinot) TrueNorm() float64 {
	// todo - not implemented
	return 0.0
}

// Float Types Notes
/*
When to use pointer vs value semantics in Go

There are three types of data in Go
- Built-in types — strings, numerics, bool. For this data type always use value types, including fields in a struct.

- Reference types — use value semantics. There is one exception. A slice and a map can take a pointer only if you sharing down to the stack and if the function called “Decode” or “Unmarshal”.

- Struct types. If we can use value semantics, use it. Pointer semantics should only be an exception.
Value semantics are essential since they keep stack clean, which give us performance.

Reference: https://medium.com/@vladbezden/when-to-use-pointer-vs-value-semantics-in-go-3718d9288b92

*/

// percent represents a percentage between 0 and 100
type percent float64

func (p *percent) String() string {
	// return fmt.Sprintf("%d%%", *p)
	// return strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) fmt.Sprintf("%.2F%%", *p*100)
}

func (p *percent) Decimal() float64 {
	return float64(*p)
}

// func (p percent) String() string {
// 	return fmt.Sprintf("%2.2F", p)
// }

// func (p percent) Decimal(precision int) float32 {
// 	e := math.Pow10(precision)
// 	pFmtString := fmt.Sprintf("%%3.%dF", precision)
// 	fmt.Println(pFmtString)
// 	fmt.Printf(pFmtString, p)
// 	return float32(math.Round(p*e) / e)
// }
