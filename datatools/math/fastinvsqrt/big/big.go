// Package big is an implementation of IEEE 754 floating point numbers.
//
// Some of the code in this package is borrowed from the Go standard
// library Package big, specifically math/big/float.go
//
// The Go Team's code:
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE_GOAUTHORS file.

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// This file implements multi-precision floating-point numbers.
// Like in the GNU MPFR library (https://www.mpfr.org/), operands
// can be of mixed precision. Unlike MPFR, the rounding mode is
// not specified with each operation, but with each operand. The
// rounding mode of the result operand determines the rounding
// mode of an operation. This is a from-scratch implementation.

package big

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/bits"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

const debugFloat = true // enable for debugging

// A nonzero finite Float represents a multi-precision floating point number
//
//   sign × mantissa × 2**exponent
//
// with 0.5 <= mantissa < 1.0, and MinExp <= exponent <= MaxExp.
// A Float may also be zero (+0, -0) or infinite (+Inf, -Inf).
// All Floats are ordered, and the ordering of two Floats x and y
// is defined by x.Cmp(y).
//
// Each Float value also has a precision, rounding mode, and accuracy.
// The precision is the maximum number of mantissa bits available to
// represent the value. The rounding mode specifies how a result should
// be rounded to fit into the mantissa bits, and accuracy describes the
// rounding error with respect to the exact result.
//
// Unless specified otherwise, all operations (including setters) that
// specify a *Float variable for the result (usually via the receiver
// with the exception of MantExp), round the numeric result according
// to the precision and rounding mode of the result variable.
//
// If the provided result precision is 0 (see below), it is set to the
// precision of the argument with the largest precision value before any
// rounding takes place, and the rounding mode remains unchanged. Thus,
// uninitialized Floats provided as result arguments will have their
// precision set to a reasonable value determined by the operands, and
// their mode is the zero value for RoundingMode (ToNearestEven).
//
// By setting the desired precision to 24 or 53 and using matching rounding
// mode (typically ToNearestEven), Float operations produce the same results
// as the corresponding float32 or float64 IEEE-754 arithmetic for operands
// that correspond to normal (i.e., not denormal) float32 or float64 numbers.
// Exponent underflow and overflow lead to a 0 or an Infinity for different
// values than IEEE-754 because Float exponents have a much larger range.
//
// The zero (uninitialized) value for a Float is ready to use and represents
// the number +0.0 exactly, with precision 0 and rounding mode ToNearestEven.
//
// Operations always take pointer arguments (*Float) rather
// than Float values, and each unique Float value requires
// its own unique *Float pointer. To "copy" a Float value,
// an existing (or newly allocated) Float must be set to
// a new value using the Float.Set method; shallow copies
// of Floats are not supported and may lead to errors.
type Float struct {
	prec uint32
	mode RoundingMode
	acc  Accuracy
	form form
	neg  bool
	mant nat
	exp  int32
}

// An ErrNaN panic is raised by a Float operation that would lead to
// a NaN under IEEE-754 rules. An ErrNaN implements the error interface.
type ErrNaN struct {
	msg string
}

func (err ErrNaN) Error() string {
	return err.msg
}

// NewFloat allocates and returns a new Float set to x,
// with precision 53 and rounding mode ToNearestEven.
// NewFloat panics with ErrNaN if x is a NaN.
func NewFloat(x float64) *Float {
	if math.IsNaN(x) {
		panic(ErrNaN{"NewFloat(NaN)"})
	}
	return new(Float).SetFloat64(x)
}

// Exponent and precision limits.
const (
	MaxExp  = math.MaxInt32  // largest supported exponent
	MinExp  = math.MinInt32  // smallest supported exponent
	MaxPrec = math.MaxUint32 // largest (theoretically) supported precision; likely memory-limited
)

// Internal representation: The mantissa bits x.mant of a nonzero finite
// Float x are stored in a nat slice long enough to hold up to x.prec bits;
// the slice may (but doesn't have to) be shorter if the mantissa contains
// trailing 0 bits. x.mant is normalized if the msb of x.mant == 1 (i.e.,
// the msb is shifted all the way "to the left"). Thus, if the mantissa has
// trailing 0 bits or x.prec is not a multiple of the Word size _W,
// x.mant[0] has trailing zero bits. The msb of the mantissa corresponds
// to the value 0.5; the exponent x.exp shifts the binary point as needed.
//
// A zero or non-finite Float x ignores x.mant and x.exp.
//
// x                 form      neg      mant         exp
// ----------------------------------------------------------
// ±0                zero      sign     -            -
// 0 < |x| < +Inf    finite    sign     mantissa     exponent
// ±Inf              inf       sign     -            -

// A form value describes the internal representation.
type form byte

// The form value order is relevant - do not change!
const (
	zero form = iota
	finite
	inf
)

// RoundingMode determines how a Float value is rounded to the
// desired precision. Rounding may change the Float value; the
// rounding error is described by the Float's Accuracy.
type RoundingMode byte

// These constants define supported rounding modes.
const (
	ToNearestEven RoundingMode = iota // == IEEE 754-2008 roundTiesToEven
	ToNearestAway                     // == IEEE 754-2008 roundTiesToAway
	ToZero                            // == IEEE 754-2008 roundTowardZero
	AwayFromZero                      // no IEEE 754-2008 equivalent
	ToNegativeInf                     // == IEEE 754-2008 roundTowardNegative
	ToPositiveInf                     // == IEEE 754-2008 roundTowardPositive
)

//go:generate stringer -type=RoundingMode

// Accuracy describes the rounding error produced by the most recent
// operation that generated a Float value, relative to the exact value.
type Accuracy int8

// Constants describing the Accuracy of a Float.
const (
	Below Accuracy = -1
	Exact Accuracy = 0
	Above Accuracy = +1
)

//go:generate stringer -type=Accuracy

// SetPrec sets z's precision to prec and returns the (possibly) rounded
// value of z. Rounding occurs according to z's rounding mode if the mantissa
// cannot be represented in prec bits without loss of precision.
// SetPrec(0) maps all finite values to ±0; infinite values remain unchanged.
// If prec > MaxPrec, it is set to MaxPrec.
func (z *Float) SetPrec(prec uint) *Float {
	z.acc = Exact // optimistically assume no rounding is needed

	// special case
	if prec == 0 {
		z.prec = 0
		if z.form == finite {
			// truncate z to 0
			z.acc = makeAcc(z.neg)
			z.form = zero
		}
		return z
	}

	// general case
	if prec > MaxPrec {
		prec = MaxPrec
	}
	old := z.prec
	z.prec = uint32(prec)
	if z.prec < old {
		z.round(0)
	}
	return z
}

func makeAcc(above bool) Accuracy {
	if above {
		return Above
	}
	return Below
}

// SetMode sets z's rounding mode to mode and returns an exact z.
// z remains unchanged otherwise.
// z.SetMode(z.Mode()) is a cheap way to set z's accuracy to Exact.
func (z *Float) SetMode(mode RoundingMode) *Float {
	z.mode = mode
	z.acc = Exact
	return z
}

// Prec returns the mantissa precision of x in bits.
// The result may be 0 for |x| == 0 and |x| == Inf.
func (x *Float) Prec() uint {
	return uint(x.prec)
}

// MinPrec returns the minimum precision required to represent x exactly
// (i.e., the smallest prec before x.SetPrec(prec) would start rounding x).
// The result is 0 for |x| == 0 and |x| == Inf.
func (x *Float) MinPrec() uint {
	if x.form != finite {
		return 0
	}
	return uint(len(x.mant))*_W - x.mant.trailingZeroBits()
}

// Mode returns the rounding mode of x.
func (x *Float) Mode() RoundingMode {
	return x.mode
}

// Acc returns the accuracy of x produced by the most recent
// operation, unless explicitly documented otherwise by that
// operation.
func (x *Float) Acc() Accuracy {
	return x.acc
}

// Sign returns:
//
//	-1 if x <   0
//	 0 if x is ±0
//	+1 if x >   0
//
func (x *Float) Sign() int {
	if debugFloat {
		x.validate()
	}
	if x.form == zero {
		return 0
	}
	if x.neg {
		return -1
	}
	return 1
}

// MantExp breaks x into its mantissa and exponent components
// and returns the exponent. If a non-nil mant argument is
// provided its value is set to the mantissa of x, with the
// same precision and rounding mode as x. The components
// satisfy x == mant × 2**exp, with 0.5 <= |mant| < 1.0.
// Calling MantExp with a nil argument is an efficient way to
// get the exponent of the receiver.
//
// Special cases are:
//
//	(  ±0).MantExp(mant) = 0, with mant set to   ±0
//	(±Inf).MantExp(mant) = 0, with mant set to ±Inf
//
// x and mant may be the same in which case x is set to its
// mantissa value.
func (x *Float) MantExp(mant *Float) (exp int) {
	if debugFloat {
		x.validate()
	}
	if x.form == finite {
		exp = int(x.exp)
	}
	if mant != nil {
		mant.Copy(x)
		if mant.form == finite {
			mant.exp = 0
		}
	}
	return
}

func (z *Float) setExpAndRound(exp int64, sbit uint) {
	if exp < MinExp {
		// underflow
		z.acc = makeAcc(z.neg)
		z.form = zero
		return
	}

	if exp > MaxExp {
		// overflow
		z.acc = makeAcc(!z.neg)
		z.form = inf
		return
	}

	z.form = finite
	z.exp = int32(exp)
	z.round(sbit)
}

// SetMantExp sets z to mant × 2**exp and returns z.
// The result z has the same precision and rounding mode
// as mant. SetMantExp is an inverse of MantExp but does
// not require 0.5 <= |mant| < 1.0. Specifically:
//
//	mant := new(Float)
//	new(Float).SetMantExp(mant, x.MantExp(mant)).Cmp(x) == 0
//
// Special cases are:
//
//	z.SetMantExp(  ±0, exp) =   ±0
//	z.SetMantExp(±Inf, exp) = ±Inf
//
// z and mant may be the same in which case z's exponent
// is set to exp.
func (z *Float) SetMantExp(mant *Float, exp int) *Float {
	if debugFloat {
		z.validate()
		mant.validate()
	}
	z.Copy(mant)
	if z.form != finite {
		return z
	}
	z.setExpAndRound(int64(z.exp)+int64(exp), 0)
	return z
}

// Signbit reports whether x is negative or negative zero.
func (x *Float) Signbit() bool {
	return x.neg
}

// IsInf reports whether x is +Inf or -Inf.
func (x *Float) IsInf() bool {
	return x.form == inf
}

// IsInt reports whether x is an integer.
// ±Inf values are not integers.
func (x *Float) IsInt() bool {
	if debugFloat {
		x.validate()
	}
	// special cases
	if x.form != finite {
		return x.form == zero
	}
	// x.form == finite
	if x.exp <= 0 {
		return false
	}
	// x.exp > 0
	return x.prec <= uint32(x.exp) || x.MinPrec() <= uint(x.exp) // not enough bits for fractional mantissa
}

// debugging support
func (x *Float) validate() {
	if !debugFloat {
		// avoid performance bugs
		panic("validate called but debugFloat is not set")
	}
	if x.form != finite {
		return
	}
	m := len(x.mant)
	if m == 0 {
		panic("nonzero finite number with empty mantissa")
	}
	const msb = 1 << (_W - 1)
	if x.mant[m-1]&msb == 0 {
		panic(fmt.Sprintf("msb not set in last word %#x of %s", x.mant[m-1], x.Text('p', 0)))
	}
	if x.prec == 0 {
		panic("zero precision finite number")
	}
}

// round rounds z according to z.mode to z.prec bits and sets z.acc accordingly.
// sbit must be 0 or 1 and summarizes any "sticky bit" information one might
// have before calling round. z's mantissa must be normalized (with the msb set)
// or empty.
//
// CAUTION: The rounding modes ToNegativeInf, ToPositiveInf are affected by the
// sign of z. For correct rounding, the sign of z must be set correctly before
// calling round.
func (z *Float) round(sbit uint) {
	if debugFloat {
		z.validate()
	}

	z.acc = Exact
	if z.form != finite {
		// ±0 or ±Inf => nothing left to do
		return
	}
	// z.form == finite && len(z.mant) > 0
	// m > 0 implies z.prec > 0 (checked by validate)

	m := uint32(len(z.mant)) // present mantissa length in words
	bits := m * _W           // present mantissa bits; bits > 0
	if bits <= z.prec {
		// mantissa fits => nothing to do
		return
	}
	// bits > z.prec

	// Rounding is based on two bits: the rounding bit (rbit) and the
	// sticky bit (sbit). The rbit is the bit immediately before the
	// z.prec leading mantissa bits (the "0.5"). The sbit is set if any
	// of the bits before the rbit are set (the "0.25", "0.125", etc.):
	//
	//   rbit  sbit  => "fractional part"
	//
	//   0     0        == 0
	//   0     1        >  0  , < 0.5
	//   1     0        == 0.5
	//   1     1        >  0.5, < 1.0

	// bits > z.prec: mantissa too large => round
	r := uint(bits - z.prec - 1) // rounding bit position; r >= 0
	rbit := z.mant.bit(r) & 1    // rounding bit; be safe and ensure it's a single bit
	// The sticky bit is only needed for rounding ToNearestEven
	// or when the rounding bit is zero. Avoid computation otherwise.
	if sbit == 0 && (rbit == 0 || z.mode == ToNearestEven) {
		sbit = z.mant.sticky(r)
	}
	sbit &= 1 // be safe and ensure it's a single bit

	// cut off extra words
	n := (z.prec + (_W - 1)) / _W // mantissa length in words for desired precision
	if m > n {
		copy(z.mant, z.mant[m-n:]) // move n last words to front
		z.mant = z.mant[:n]
	}

	// determine number of trailing zero bits (ntz) and compute lsb mask of mantissa's least-significant word
	ntz := n*_W - z.prec // 0 <= ntz < _W
	lsb := Word(1) << ntz

	// round if result is inexact
	if rbit|sbit != 0 {
		// Make rounding decision: The result mantissa is truncated ("rounded down")
		// by default. Decide if we need to increment, or "round up", the (unsigned)
		// mantissa.
		inc := false
		switch z.mode {
		case ToNegativeInf:
			inc = z.neg
		case ToZero:
			// nothing to do
		case ToNearestEven:
			inc = rbit != 0 && (sbit != 0 || z.mant[0]&lsb != 0)
		case ToNearestAway:
			inc = rbit != 0
		case AwayFromZero:
			inc = true
		case ToPositiveInf:
			inc = !z.neg
		default:
			panic("unreachable")
		}

		// A positive result (!z.neg) is Above the exact result if we increment,
		// and it's Below if we truncate (Exact results require no rounding).
		// For a negative result (z.neg) it is exactly the opposite.
		z.acc = makeAcc(inc != z.neg)

		if inc {
			// add 1 to mantissa
			if addVW(z.mant, z.mant, lsb) != 0 {
				// mantissa overflow => adjust exponent
				if z.exp >= MaxExp {
					// exponent overflow
					z.form = inf
					return
				}
				z.exp++
				// adjust mantissa: divide by 2 to compensate for exponent adjustment
				shrVU(z.mant, z.mant, 1)
				// set msb == carry == 1 from the mantissa overflow above
				const msb = 1 << (_W - 1)
				z.mant[n-1] |= msb
			}
		}
	}

	// zero out trailing bits in least-significant word
	z.mant[0] &^= lsb - 1

	if debugFloat {
		z.validate()
	}
}

func (z *Float) setBits64(neg bool, x uint64) *Float {
	if z.prec == 0 {
		z.prec = 64
	}
	z.acc = Exact
	z.neg = neg
	if x == 0 {
		z.form = zero
		return z
	}
	// x != 0
	z.form = finite
	s := bits.LeadingZeros64(x)
	z.mant = z.mant.setUint64(x << uint(s))
	z.exp = int32(64 - s) // always fits
	if z.prec < 64 {
		z.round(0)
	}
	return z
}

// SetUint64 sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to 64 (and rounding will have
// no effect).
func (z *Float) SetUint64(x uint64) *Float {
	return z.setBits64(false, x)
}

// SetInt64 sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to 64 (and rounding will have
// no effect).
func (z *Float) SetInt64(x int64) *Float {
	u := x
	if u < 0 {
		u = -u
	}
	// We cannot simply call z.SetUint64(uint64(u)) and change
	// the sign afterwards because the sign affects rounding.
	return z.setBits64(x < 0, uint64(u))
}

// SetFloat64 sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to 53 (and rounding will have
// no effect). SetFloat64 panics with ErrNaN if x is a NaN.
func (z *Float) SetFloat64(x float64) *Float {
	if z.prec == 0 {
		z.prec = 53
	}
	if math.IsNaN(x) {
		panic(ErrNaN{"Float.SetFloat64(NaN)"})
	}
	z.acc = Exact
	z.neg = math.Signbit(x) // handle -0, -Inf correctly
	if x == 0 {
		z.form = zero
		return z
	}
	if math.IsInf(x, 0) {
		z.form = inf
		return z
	}
	// normalized x != 0
	z.form = finite
	fmant, exp := math.Frexp(x) // get normalized mantissa
	z.mant = z.mant.setUint64(1<<63 | math.Float64bits(fmant)<<11)
	z.exp = int32(exp) // always fits
	if z.prec < 53 {
		z.round(0)
	}
	return z
}

// fnorm normalizes mantissa m by shifting it to the left
// such that the msb of the most-significant word (msw) is 1.
// It returns the shift amount. It assumes that len(m) != 0.
func fnorm(m nat) int64 {
	if debugFloat && (len(m) == 0 || m[len(m)-1] == 0) {
		panic("msw of mantissa is 0")
	}
	s := nlz(m[len(m)-1])
	if s > 0 {
		c := shlVU(m, m, s)
		if debugFloat && c != 0 {
			panic("nlz or shlVU incorrect")
		}
	}
	return int64(s)
}

// SetInt sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to the larger of x.BitLen()
// or 64 (and rounding will have no effect).
func (z *Float) SetInt(x *Int) *Float {
	// TODO(gri) can be more efficient if z.prec > 0
	// but small compared to the size of x, or if there
	// are many trailing 0's.
	bits := uint32(x.BitLen())
	if z.prec == 0 {
		z.prec = umax32(bits, 64)
	}
	z.acc = Exact
	z.neg = x.neg
	if len(x.abs) == 0 {
		z.form = zero
		return z
	}
	// x != 0
	z.mant = z.mant.set(x.abs)
	fnorm(z.mant)
	z.setExpAndRound(int64(bits), 0)
	return z
}

// SetRat sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to the largest of a.BitLen(),
// b.BitLen(), or 64; with x = a/b.
func (z *Float) SetRat(x *Rat) *Float {
	if x.IsInt() {
		return z.SetInt(x.Num())
	}
	var a, b Float
	a.SetInt(x.Num())
	b.SetInt(x.Denom())
	if z.prec == 0 {
		z.prec = umax32(a.prec, b.prec)
	}
	return z.Quo(&a, &b)
}

// SetInf sets z to the infinite Float -Inf if signbit is
// set, or +Inf if signbit is not set, and returns z. The
// precision of z is unchanged and the result is always
// Exact.
func (z *Float) SetInf(signbit bool) *Float {
	z.acc = Exact
	z.form = inf
	z.neg = signbit
	return z
}

// Set sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to the precision of x
// before setting z (and rounding will have no effect).
// Rounding is performed according to z's precision and rounding
// mode; and z's accuracy reports the result error relative to the
// exact (not rounded) result.
func (z *Float) Set(x *Float) *Float {
	if debugFloat {
		x.validate()
	}
	z.acc = Exact
	if z != x {
		z.form = x.form
		z.neg = x.neg
		if x.form == finite {
			z.exp = x.exp
			z.mant = z.mant.set(x.mant)
		}
		if z.prec == 0 {
			z.prec = x.prec
		} else if z.prec < x.prec {
			z.round(0)
		}
	}
	return z
}

// Copy sets z to x, with the same precision, rounding mode, and
// accuracy as x, and returns z. x is not changed even if z and
// x are the same.
func (z *Float) Copy(x *Float) *Float {
	if debugFloat {
		x.validate()
	}
	if z != x {
		z.prec = x.prec
		z.mode = x.mode
		z.acc = x.acc
		z.form = x.form
		z.neg = x.neg
		if z.form == finite {
			z.mant = z.mant.set(x.mant)
			z.exp = x.exp
		}
	}
	return z
}

// msb32 returns the 32 most significant bits of x.
func msb32(x nat) uint32 {
	i := len(x) - 1
	if i < 0 {
		return 0
	}
	if debugFloat && x[i]&(1<<(_W-1)) == 0 {
		panic("x not normalized")
	}
	switch _W {
	case 32:
		return uint32(x[i])
	case 64:
		return uint32(x[i] >> 32)
	}
	panic("unreachable")
}

// msb64 returns the 64 most significant bits of x.
func msb64(x nat) uint64 {
	i := len(x) - 1
	if i < 0 {
		return 0
	}
	if debugFloat && x[i]&(1<<(_W-1)) == 0 {
		panic("x not normalized")
	}
	switch _W {
	case 32:
		v := uint64(x[i]) << 32
		if i > 0 {
			v |= uint64(x[i-1])
		}
		return v
	case 64:
		return uint64(x[i])
	}
	panic("unreachable")
}

// Uint64 returns the unsigned integer resulting from truncating x
// towards zero. If 0 <= x <= math.MaxUint64, the result is Exact
// if x is an integer and Below otherwise.
// The result is (0, Above) for x < 0, and (math.MaxUint64, Below)
// for x > math.MaxUint64.
func (x *Float) Uint64() (uint64, Accuracy) {
	if debugFloat {
		x.validate()
	}

	switch x.form {
	case finite:
		if x.neg {
			return 0, Above
		}
		// 0 < x < +Inf
		if x.exp <= 0 {
			// 0 < x < 1
			return 0, Below
		}
		// 1 <= x < Inf
		if x.exp <= 64 {
			// u = trunc(x) fits into a uint64
			u := msb64(x.mant) >> (64 - uint32(x.exp))
			if x.MinPrec() <= 64 {
				return u, Exact
			}
			return u, Below // x truncated
		}
		// x too large
		return math.MaxUint64, Below

	case zero:
		return 0, Exact

	case inf:
		if x.neg {
			return 0, Above
		}
		return math.MaxUint64, Below
	}

	panic("unreachable")
}

// Int64 returns the integer resulting from truncating x towards zero.
// If math.MinInt64 <= x <= math.MaxInt64, the result is Exact if x is
// an integer, and Above (x < 0) or Below (x > 0) otherwise.
// The result is (math.MinInt64, Above) for x < math.MinInt64,
// and (math.MaxInt64, Below) for x > math.MaxInt64.
func (x *Float) Int64() (int64, Accuracy) {
	if debugFloat {
		x.validate()
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf
		acc := makeAcc(x.neg)
		if x.exp <= 0 {
			// 0 < |x| < 1
			return 0, acc
		}
		// x.exp > 0

		// 1 <= |x| < +Inf
		if x.exp <= 63 {
			// i = trunc(x) fits into an int64 (excluding math.MinInt64)
			i := int64(msb64(x.mant) >> (64 - uint32(x.exp)))
			if x.neg {
				i = -i
			}
			if x.MinPrec() <= uint(x.exp) {
				return i, Exact
			}
			return i, acc // x truncated
		}
		if x.neg {
			// check for special case x == math.MinInt64 (i.e., x == -(0.5 << 64))
			if x.exp == 64 && x.MinPrec() == 1 {
				acc = Exact
			}
			return math.MinInt64, acc
		}
		// x too large
		return math.MaxInt64, Below

	case zero:
		return 0, Exact

	case inf:
		if x.neg {
			return math.MinInt64, Above
		}
		return math.MaxInt64, Below
	}

	panic("unreachable")
}

// Float32 returns the float32 value nearest to x. If x is too small to be
// represented by a float32 (|x| < math.SmallestNonzeroFloat32), the result
// is (0, Below) or (-0, Above), respectively, depending on the sign of x.
// If x is too large to be represented by a float32 (|x| > math.MaxFloat32),
// the result is (+Inf, Above) or (-Inf, Below), depending on the sign of x.
func (x *Float) Float32() (float32, Accuracy) {
	if debugFloat {
		x.validate()
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf

		const (
			fbits = 32                //        float size
			mbits = 23                //        mantissa size (excluding implicit msb)
			ebits = fbits - mbits - 1 //     8  exponent size
			bias  = 1<<(ebits-1) - 1  //   127  exponent bias
			dmin  = 1 - bias - mbits  //  -149  smallest unbiased exponent (denormal)
			emin  = 1 - bias          //  -126  smallest unbiased exponent (normal)
			emax  = bias              //   127  largest unbiased exponent (normal)
		)

		// Float mantissa m is 0.5 <= m < 1.0; compute exponent e for float32 mantissa.
		e := x.exp - 1 // exponent for normal mantissa m with 1.0 <= m < 2.0

		// Compute precision p for float32 mantissa.
		// If the exponent is too small, we have a denormal number before
		// rounding and fewer than p mantissa bits of precision available
		// (the exponent remains fixed but the mantissa gets shifted right).
		p := mbits + 1 // precision of normal float
		if e < emin {
			// recompute precision
			p = mbits + 1 - emin + int(e)
			// If p == 0, the mantissa of x is shifted so much to the right
			// that its msb falls immediately to the right of the float32
			// mantissa space. In other words, if the smallest denormal is
			// considered "1.0", for p == 0, the mantissa value m is >= 0.5.
			// If m > 0.5, it is rounded up to 1.0; i.e., the smallest denormal.
			// If m == 0.5, it is rounded down to even, i.e., 0.0.
			// If p < 0, the mantissa value m is <= "0.25" which is never rounded up.
			if p < 0 /* m <= 0.25 */ || p == 0 && x.mant.sticky(uint(len(x.mant))*_W-1) == 0 /* m == 0.5 */ {
				// underflow to ±0
				if x.neg {
					var z float32
					return -z, Above
				}
				return 0.0, Below
			}
			// otherwise, round up
			// We handle p == 0 explicitly because it's easy and because
			// Float.round doesn't support rounding to 0 bits of precision.
			if p == 0 {
				if x.neg {
					return -math.SmallestNonzeroFloat32, Below
				}
				return math.SmallestNonzeroFloat32, Above
			}
		}
		// p > 0

		// round
		var r Float
		r.prec = uint32(p)
		r.Set(x)
		e = r.exp - 1

		// Rounding may have caused r to overflow to ±Inf
		// (rounding never causes underflows to 0).
		// If the exponent is too large, also overflow to ±Inf.
		if r.form == inf || e > emax {
			// overflow
			if x.neg {
				return float32(math.Inf(-1)), Below
			}
			return float32(math.Inf(+1)), Above
		}
		// e <= emax

		// Determine sign, biased exponent, and mantissa.
		var sign, bexp, mant uint32
		if x.neg {
			sign = 1 << (fbits - 1)
		}

		// Rounding may have caused a denormal number to
		// become normal. Check again.
		if e < emin {
			// denormal number: recompute precision
			// Since rounding may have at best increased precision
			// and we have eliminated p <= 0 early, we know p > 0.
			// bexp == 0 for denormals
			p = mbits + 1 - emin + int(e)
			mant = msb32(r.mant) >> uint(fbits-p)
		} else {
			// normal number: emin <= e <= emax
			bexp = uint32(e+bias) << mbits
			mant = msb32(r.mant) >> ebits & (1<<mbits - 1) // cut off msb (implicit 1 bit)
		}

		return math.Float32frombits(sign | bexp | mant), r.acc

	case zero:
		if x.neg {
			var z float32
			return -z, Exact
		}
		return 0.0, Exact

	case inf:
		if x.neg {
			return float32(math.Inf(-1)), Exact
		}
		return float32(math.Inf(+1)), Exact
	}

	panic("unreachable")
}

// Float64 returns the float64 value nearest to x. If x is too small to be
// represented by a float64 (|x| < math.SmallestNonzeroFloat64), the result
// is (0, Below) or (-0, Above), respectively, depending on the sign of x.
// If x is too large to be represented by a float64 (|x| > math.MaxFloat64),
// the result is (+Inf, Above) or (-Inf, Below), depending on the sign of x.
func (x *Float) Float64() (float64, Accuracy) {
	if debugFloat {
		x.validate()
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf

		const (
			fbits = 64                //        float size
			mbits = 52                //        mantissa size (excluding implicit msb)
			ebits = fbits - mbits - 1 //    11  exponent size
			bias  = 1<<(ebits-1) - 1  //  1023  exponent bias
			dmin  = 1 - bias - mbits  // -1074  smallest unbiased exponent (denormal)
			emin  = 1 - bias          // -1022  smallest unbiased exponent (normal)
			emax  = bias              //  1023  largest unbiased exponent (normal)
		)

		// Float mantissa m is 0.5 <= m < 1.0; compute exponent e for float64 mantissa.
		e := x.exp - 1 // exponent for normal mantissa m with 1.0 <= m < 2.0

		// Compute precision p for float64 mantissa.
		// If the exponent is too small, we have a denormal number before
		// rounding and fewer than p mantissa bits of precision available
		// (the exponent remains fixed but the mantissa gets shifted right).
		p := mbits + 1 // precision of normal float
		if e < emin {
			// recompute precision
			p = mbits + 1 - emin + int(e)
			// If p == 0, the mantissa of x is shifted so much to the right
			// that its msb falls immediately to the right of the float64
			// mantissa space. In other words, if the smallest denormal is
			// considered "1.0", for p == 0, the mantissa value m is >= 0.5.
			// If m > 0.5, it is rounded up to 1.0; i.e., the smallest denormal.
			// If m == 0.5, it is rounded down to even, i.e., 0.0.
			// If p < 0, the mantissa value m is <= "0.25" which is never rounded up.
			if p < 0 /* m <= 0.25 */ || p == 0 && x.mant.sticky(uint(len(x.mant))*_W-1) == 0 /* m == 0.5 */ {
				// underflow to ±0
				if x.neg {
					var z float64
					return -z, Above
				}
				return 0.0, Below
			}
			// otherwise, round up
			// We handle p == 0 explicitly because it's easy and because
			// Float.round doesn't support rounding to 0 bits of precision.
			if p == 0 {
				if x.neg {
					return -math.SmallestNonzeroFloat64, Below
				}
				return math.SmallestNonzeroFloat64, Above
			}
		}
		// p > 0

		// round
		var r Float
		r.prec = uint32(p)
		r.Set(x)
		e = r.exp - 1

		// Rounding may have caused r to overflow to ±Inf
		// (rounding never causes underflows to 0).
		// If the exponent is too large, also overflow to ±Inf.
		if r.form == inf || e > emax {
			// overflow
			if x.neg {
				return math.Inf(-1), Below
			}
			return math.Inf(+1), Above
		}
		// e <= emax

		// Determine sign, biased exponent, and mantissa.
		var sign, bexp, mant uint64
		if x.neg {
			sign = 1 << (fbits - 1)
		}

		// Rounding may have caused a denormal number to
		// become normal. Check again.
		if e < emin {
			// denormal number: recompute precision
			// Since rounding may have at best increased precision
			// and we have eliminated p <= 0 early, we know p > 0.
			// bexp == 0 for denormals
			p = mbits + 1 - emin + int(e)
			mant = msb64(r.mant) >> uint(fbits-p)
		} else {
			// normal number: emin <= e <= emax
			bexp = uint64(e+bias) << mbits
			mant = msb64(r.mant) >> ebits & (1<<mbits - 1) // cut off msb (implicit 1 bit)
		}

		return math.Float64frombits(sign | bexp | mant), r.acc

	case zero:
		if x.neg {
			var z float64
			return -z, Exact
		}
		return 0.0, Exact

	case inf:
		if x.neg {
			return math.Inf(-1), Exact
		}
		return math.Inf(+1), Exact
	}

	panic("unreachable")
}

// Int returns the result of truncating x towards zero;
// or nil if x is an infinity.
// The result is Exact if x.IsInt(); otherwise it is Below
// for x > 0, and Above for x < 0.
// If a non-nil *Int argument z is provided, Int stores
// the result in z instead of allocating a new Int.
func (x *Float) Int(z *Int) (*Int, Accuracy) {
	if debugFloat {
		x.validate()
	}

	if z == nil && x.form <= finite {
		z = new(Int)
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf
		acc := makeAcc(x.neg)
		if x.exp <= 0 {
			// 0 < |x| < 1
			return z.SetInt64(0), acc
		}
		// x.exp > 0

		// 1 <= |x| < +Inf
		// determine minimum required precision for x
		allBits := uint(len(x.mant)) * _W
		exp := uint(x.exp)
		if x.MinPrec() <= exp {
			acc = Exact
		}
		// shift mantissa as needed
		if z == nil {
			z = new(Int)
		}
		z.neg = x.neg
		switch {
		case exp > allBits:
			z.abs = z.abs.shl(x.mant, exp-allBits)
		default:
			z.abs = z.abs.set(x.mant)
		case exp < allBits:
			z.abs = z.abs.shr(x.mant, allBits-exp)
		}
		return z, acc

	case zero:
		return z.SetInt64(0), Exact

	case inf:
		return nil, makeAcc(x.neg)
	}

	panic("unreachable")
}

// Rat returns the rational number corresponding to x;
// or nil if x is an infinity.
// The result is Exact if x is not an Inf.
// If a non-nil *Rat argument z is provided, Rat stores
// the result in z instead of allocating a new Rat.
func (x *Float) Rat(z *Rat) (*Rat, Accuracy) {
	if debugFloat {
		x.validate()
	}

	if z == nil && x.form <= finite {
		z = new(Rat)
	}

	switch x.form {
	case finite:
		// 0 < |x| < +Inf
		allBits := int32(len(x.mant)) * _W
		// build up numerator and denominator
		z.a.neg = x.neg
		switch {
		case x.exp > allBits:
			z.a.abs = z.a.abs.shl(x.mant, uint(x.exp-allBits))
			z.b.abs = z.b.abs[:0] // == 1 (see Rat)
			// z already in normal form
		default:
			z.a.abs = z.a.abs.set(x.mant)
			z.b.abs = z.b.abs[:0] // == 1 (see Rat)
			// z already in normal form
		case x.exp < allBits:
			z.a.abs = z.a.abs.set(x.mant)
			t := z.b.abs.setUint64(1)
			z.b.abs = t.shl(t, uint(allBits-x.exp))
			z.norm()
		}
		return z, Exact

	case zero:
		return z.SetInt64(0), Exact

	case inf:
		return nil, makeAcc(x.neg)
	}

	panic("unreachable")
}

// Abs sets z to the (possibly rounded) value |x| (the absolute value of x)
// and returns z.
func (z *Float) Abs(x *Float) *Float {
	z.Set(x)
	z.neg = false
	return z
}

// Neg sets z to the (possibly rounded) value of x with its sign negated,
// and returns z.
func (z *Float) Neg(x *Float) *Float {
	z.Set(x)
	z.neg = !z.neg
	return z
}

func validateBinaryOperands(x, y *Float) {
	if !debugFloat {
		// avoid performance bugs
		panic("validateBinaryOperands called but debugFloat is not set")
	}
	if len(x.mant) == 0 {
		panic("empty mantissa for x")
	}
	if len(y.mant) == 0 {
		panic("empty mantissa for y")
	}
}

// z = x + y, ignoring signs of x and y for the addition
// but using the sign of z for rounding the result.
// x and y must have a non-empty mantissa and valid exponent.
func (z *Float) uadd(x, y *Float) {
	// Note: This implementation requires 2 shifts most of the
	// time. It is also inefficient if exponents or precisions
	// differ by wide margins. The following article describes
	// an efficient (but much more complicated) implementation
	// compatible with the internal representation used here:
	//
	// Vincent Lefèvre: "The Generic Multiple-Precision Floating-
	// Point Addition With Exact Rounding (as in the MPFR Library)"
	// http://www.vinc17.net/research/papers/rnc6.pdf

	if debugFloat {
		validateBinaryOperands(x, y)
	}

	// compute exponents ex, ey for mantissa with "binary point"
	// on the right (mantissa.0) - use int64 to avoid overflow
	ex := int64(x.exp) - int64(len(x.mant))*_W
	ey := int64(y.exp) - int64(len(y.mant))*_W

	al := alias(z.mant, x.mant) || alias(z.mant, y.mant)

	// TODO(gri) having a combined add-and-shift primitive
	//           could make this code significantly faster
	switch {
	case ex < ey:
		if al {
			t := nat(nil).shl(y.mant, uint(ey-ex))
			z.mant = z.mant.add(x.mant, t)
		} else {
			z.mant = z.mant.shl(y.mant, uint(ey-ex))
			z.mant = z.mant.add(x.mant, z.mant)
		}
	default:
		// ex == ey, no shift needed
		z.mant = z.mant.add(x.mant, y.mant)
	case ex > ey:
		if al {
			t := nat(nil).shl(x.mant, uint(ex-ey))
			z.mant = z.mant.add(t, y.mant)
		} else {
			z.mant = z.mant.shl(x.mant, uint(ex-ey))
			z.mant = z.mant.add(z.mant, y.mant)
		}
		ex = ey
	}
	// len(z.mant) > 0

	z.setExpAndRound(ex+int64(len(z.mant))*_W-fnorm(z.mant), 0)
}

// z = x - y for |x| > |y|, ignoring signs of x and y for the subtraction
// but using the sign of z for rounding the result.
// x and y must have a non-empty mantissa and valid exponent.
func (z *Float) usub(x, y *Float) {
	// This code is symmetric to uadd.
	// We have not factored the common code out because
	// eventually uadd (and usub) should be optimized
	// by special-casing, and the code will diverge.

	if debugFloat {
		validateBinaryOperands(x, y)
	}

	ex := int64(x.exp) - int64(len(x.mant))*_W
	ey := int64(y.exp) - int64(len(y.mant))*_W

	al := alias(z.mant, x.mant) || alias(z.mant, y.mant)

	switch {
	case ex < ey:
		if al {
			t := nat(nil).shl(y.mant, uint(ey-ex))
			z.mant = t.sub(x.mant, t)
		} else {
			z.mant = z.mant.shl(y.mant, uint(ey-ex))
			z.mant = z.mant.sub(x.mant, z.mant)
		}
	default:
		// ex == ey, no shift needed
		z.mant = z.mant.sub(x.mant, y.mant)
	case ex > ey:
		if al {
			t := nat(nil).shl(x.mant, uint(ex-ey))
			z.mant = t.sub(t, y.mant)
		} else {
			z.mant = z.mant.shl(x.mant, uint(ex-ey))
			z.mant = z.mant.sub(z.mant, y.mant)
		}
		ex = ey
	}

	// operands may have canceled each other out
	if len(z.mant) == 0 {
		z.acc = Exact
		z.form = zero
		z.neg = false
		return
	}
	// len(z.mant) > 0

	z.setExpAndRound(ex+int64(len(z.mant))*_W-fnorm(z.mant), 0)
}

// z = x * y, ignoring signs of x and y for the multiplication
// but using the sign of z for rounding the result.
// x and y must have a non-empty mantissa and valid exponent.
func (z *Float) umul(x, y *Float) {
	if debugFloat {
		validateBinaryOperands(x, y)
	}

	// Note: This is doing too much work if the precision
	// of z is less than the sum of the precisions of x
	// and y which is often the case (e.g., if all floats
	// have the same precision).
	// TODO(gri) Optimize this for the common case.

	e := int64(x.exp) + int64(y.exp)
	if x == y {
		z.mant = z.mant.sqr(x.mant)
	} else {
		z.mant = z.mant.mul(x.mant, y.mant)
	}
	z.setExpAndRound(e-fnorm(z.mant), 0)
}

// z = x / y, ignoring signs of x and y for the division
// but using the sign of z for rounding the result.
// x and y must have a non-empty mantissa and valid exponent.
func (z *Float) uquo(x, y *Float) {
	if debugFloat {
		validateBinaryOperands(x, y)
	}

	// mantissa length in words for desired result precision + 1
	// (at least one extra bit so we get the rounding bit after
	// the division)
	n := int(z.prec/_W) + 1

	// compute adjusted x.mant such that we get enough result precision
	xadj := x.mant
	if d := n - len(x.mant) + len(y.mant); d > 0 {
		// d extra words needed => add d "0 digits" to x
		xadj = make(nat, len(x.mant)+d)
		copy(xadj[d:], x.mant)
	}
	// TODO(gri): If we have too many digits (d < 0), we should be able
	// to shorten x for faster division. But we must be extra careful
	// with rounding in that case.

	// Compute d before division since there may be aliasing of x.mant
	// (via xadj) or y.mant with z.mant.
	d := len(xadj) - len(y.mant)

	// divide
	var r nat
	z.mant, r = z.mant.div(nil, xadj, y.mant)
	e := int64(x.exp) - int64(y.exp) - int64(d-len(z.mant))*_W

	// The result is long enough to include (at least) the rounding bit.
	// If there's a non-zero remainder, the corresponding fractional part
	// (if it were computed), would have a non-zero sticky bit (if it were
	// zero, it couldn't have a non-zero remainder).
	var sbit uint
	if len(r) > 0 {
		sbit = 1
	}

	z.setExpAndRound(e-fnorm(z.mant), sbit)
}

// ucmp returns -1, 0, or +1, depending on whether
// |x| < |y|, |x| == |y|, or |x| > |y|.
// x and y must have a non-empty mantissa and valid exponent.
func (x *Float) ucmp(y *Float) int {
	if debugFloat {
		validateBinaryOperands(x, y)
	}

	switch {
	case x.exp < y.exp:
		return -1
	case x.exp > y.exp:
		return +1
	}
	// x.exp == y.exp

	// compare mantissas
	i := len(x.mant)
	j := len(y.mant)
	for i > 0 || j > 0 {
		var xm, ym Word
		if i > 0 {
			i--
			xm = x.mant[i]
		}
		if j > 0 {
			j--
			ym = y.mant[j]
		}
		switch {
		case xm < ym:
			return -1
		case xm > ym:
			return +1
		}
	}

	return 0
}

// Handling of sign bit as defined by IEEE 754-2008, section 6.3:
//
// When neither the inputs nor result are NaN, the sign of a product or
// quotient is the exclusive OR of the operands’ signs; the sign of a sum,
// or of a difference x−y regarded as a sum x+(−y), differs from at most
// one of the addends’ signs; and the sign of the result of conversions,
// the quantize operation, the roundToIntegral operations, and the
// roundToIntegralExact (see 5.3.1) is the sign of the first or only operand.
// These rules shall apply even when operands or results are zero or infinite.
//
// When the sum of two operands with opposite signs (or the difference of
// two operands with like signs) is exactly zero, the sign of that sum (or
// difference) shall be +0 in all rounding-direction attributes except
// roundTowardNegative; under that attribute, the sign of an exact zero
// sum (or difference) shall be −0. However, x+x = x−(−x) retains the same
// sign as x even when x is zero.
//
// See also: https://play.golang.org/p/RtH3UCt5IH

// Add sets z to the rounded sum x+y and returns z. If z's precision is 0,
// it is changed to the larger of x's or y's precision before the operation.
// Rounding is performed according to z's precision and rounding mode; and
// z's accuracy reports the result error relative to the exact (not rounded)
// result. Add panics with ErrNaN if x and y are infinities with opposite
// signs. The value of z is undefined in that case.
func (z *Float) Add(x, y *Float) *Float {
	if debugFloat {
		x.validate()
		y.validate()
	}

	if z.prec == 0 {
		z.prec = umax32(x.prec, y.prec)
	}

	if x.form == finite && y.form == finite {
		// x + y (common case)

		// Below we set z.neg = x.neg, and when z aliases y this will
		// change the y operand's sign. This is fine, because if an
		// operand aliases the receiver it'll be overwritten, but we still
		// want the original x.neg and y.neg values when we evaluate
		// x.neg != y.neg, so we need to save y.neg before setting z.neg.
		yneg := y.neg

		z.neg = x.neg
		if x.neg == yneg {
			// x + y == x + y
			// (-x) + (-y) == -(x + y)
			z.uadd(x, y)
		} else {
			// x + (-y) == x - y == -(y - x)
			// (-x) + y == y - x == -(x - y)
			if x.ucmp(y) > 0 {
				z.usub(x, y)
			} else {
				z.neg = !z.neg
				z.usub(y, x)
			}
		}
		if z.form == zero && z.mode == ToNegativeInf && z.acc == Exact {
			z.neg = true
		}
		return z
	}

	if x.form == inf && y.form == inf && x.neg != y.neg {
		// +Inf + -Inf
		// -Inf + +Inf
		// value of z is undefined but make sure it's valid
		z.acc = Exact
		z.form = zero
		z.neg = false
		panic(ErrNaN{"addition of infinities with opposite signs"})
	}

	if x.form == zero && y.form == zero {
		// ±0 + ±0
		z.acc = Exact
		z.form = zero
		z.neg = x.neg && y.neg // -0 + -0 == -0
		return z
	}

	if x.form == inf || y.form == zero {
		// ±Inf + y
		// x + ±0
		return z.Set(x)
	}

	// ±0 + y
	// x + ±Inf
	return z.Set(y)
}

// Sub sets z to the rounded difference x-y and returns z.
// Precision, rounding, and accuracy reporting are as for Add.
// Sub panics with ErrNaN if x and y are infinities with equal
// signs. The value of z is undefined in that case.
func (z *Float) Sub(x, y *Float) *Float {
	if debugFloat {
		x.validate()
		y.validate()
	}

	if z.prec == 0 {
		z.prec = umax32(x.prec, y.prec)
	}

	if x.form == finite && y.form == finite {
		// x - y (common case)
		yneg := y.neg
		z.neg = x.neg
		if x.neg != yneg {
			// x - (-y) == x + y
			// (-x) - y == -(x + y)
			z.uadd(x, y)
		} else {
			// x - y == x - y == -(y - x)
			// (-x) - (-y) == y - x == -(x - y)
			if x.ucmp(y) > 0 {
				z.usub(x, y)
			} else {
				z.neg = !z.neg
				z.usub(y, x)
			}
		}
		if z.form == zero && z.mode == ToNegativeInf && z.acc == Exact {
			z.neg = true
		}
		return z
	}

	if x.form == inf && y.form == inf && x.neg == y.neg {
		// +Inf - +Inf
		// -Inf - -Inf
		// value of z is undefined but make sure it's valid
		z.acc = Exact
		z.form = zero
		z.neg = false
		panic(ErrNaN{"subtraction of infinities with equal signs"})
	}

	if x.form == zero && y.form == zero {
		// ±0 - ±0
		z.acc = Exact
		z.form = zero
		z.neg = x.neg && !y.neg // -0 - +0 == -0
		return z
	}

	if x.form == inf || y.form == zero {
		// ±Inf - y
		// x - ±0
		return z.Set(x)
	}

	// ±0 - y
	// x - ±Inf
	return z.Neg(y)
}

// Mul sets z to the rounded product x*y and returns z.
// Precision, rounding, and accuracy reporting are as for Add.
// Mul panics with ErrNaN if one operand is zero and the other
// operand an infinity. The value of z is undefined in that case.
func (z *Float) Mul(x, y *Float) *Float {
	if debugFloat {
		x.validate()
		y.validate()
	}

	if z.prec == 0 {
		z.prec = umax32(x.prec, y.prec)
	}

	z.neg = x.neg != y.neg

	if x.form == finite && y.form == finite {
		// x * y (common case)
		z.umul(x, y)
		return z
	}

	z.acc = Exact
	if x.form == zero && y.form == inf || x.form == inf && y.form == zero {
		// ±0 * ±Inf
		// ±Inf * ±0
		// value of z is undefined but make sure it's valid
		z.form = zero
		z.neg = false
		panic(ErrNaN{"multiplication of zero with infinity"})
	}

	if x.form == inf || y.form == inf {
		// ±Inf * y
		// x * ±Inf
		z.form = inf
		return z
	}

	// ±0 * y
	// x * ±0
	z.form = zero
	return z
}

// Quo sets z to the rounded quotient x/y and returns z.
// Precision, rounding, and accuracy reporting are as for Add.
// Quo panics with ErrNaN if both operands are zero or infinities.
// The value of z is undefined in that case.
func (z *Float) Quo(x, y *Float) *Float {
	if debugFloat {
		x.validate()
		y.validate()
	}

	if z.prec == 0 {
		z.prec = umax32(x.prec, y.prec)
	}

	z.neg = x.neg != y.neg

	if x.form == finite && y.form == finite {
		// x / y (common case)
		z.uquo(x, y)
		return z
	}

	z.acc = Exact
	if x.form == zero && y.form == zero || x.form == inf && y.form == inf {
		// ±0 / ±0
		// ±Inf / ±Inf
		// value of z is undefined but make sure it's valid
		z.form = zero
		z.neg = false
		panic(ErrNaN{"division of zero by zero or infinity by infinity"})
	}

	if x.form == zero || y.form == inf {
		// ±0 / y
		// x / ±Inf
		z.form = zero
		return z
	}

	// x / ±0
	// ±Inf / y
	z.form = inf
	return z
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
//   +1 if x >  y
//
func (x *Float) Cmp(y *Float) int {
	if debugFloat {
		x.validate()
		y.validate()
	}

	mx := x.ord()
	my := y.ord()
	switch {
	case mx < my:
		return -1
	case mx > my:
		return +1
	}
	// mx == my

	// only if |mx| == 1 we have to compare the mantissae
	switch mx {
	case -1:
		return y.ucmp(x)
	case +1:
		return x.ucmp(y)
	}

	return 0
}

// ord classifies x and returns:
//
//	-2 if -Inf == x
//	-1 if -Inf < x < 0
//	 0 if x == 0 (signed or unsigned)
//	+1 if 0 < x < +Inf
//	+2 if x == +Inf
//
func (x *Float) ord() int {
	var m int
	switch x.form {
	case finite:
		m = 1
	case zero:
		return 0
	case inf:
		m = 2
	}
	if x.neg {
		m = -m
	}
	return m
}

func umax32(x, y uint32) uint32 {
	if x > y {
		return x
	}
	return y
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// math/big/floatconv.go
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements string-to-Float conversion functions.

var floatZero Float

// SetString sets z to the value of s and returns z and a boolean indicating
// success. s must be a floating-point number of the same format as accepted
// by Parse, with base argument 0. The entire string (not just a prefix) must
// be valid for success. If the operation failed, the value of z is undefined
// but the returned value is nil.
func (z *Float) SetString(s string) (*Float, bool) {
	if f, _, err := z.Parse(s, 0); err == nil {
		return f, true
	}
	return nil, false
}

// scan is like Parse but reads the longest possible prefix representing a valid
// floating point number from an io.ByteScanner rather than a string. It serves
// as the implementation of Parse. It does not recognize ±Inf and does not expect
// EOF at the end.
func (z *Float) scan(r io.ByteScanner, base int) (f *Float, b int, err error) {
	prec := z.prec
	if prec == 0 {
		prec = 64
	}

	// A reasonable value in case of an error.
	z.form = zero

	// sign
	z.neg, err = scanSign(r)
	if err != nil {
		return
	}

	// mantissa
	var fcount int // fractional digit count; valid if <= 0
	z.mant, b, fcount, err = z.mant.scan(r, base, true)
	if err != nil {
		return
	}

	// exponent
	var exp int64
	var ebase int
	exp, ebase, err = scanExponent(r, true, base == 0)
	if err != nil {
		return
	}

	// special-case 0
	if len(z.mant) == 0 {
		z.prec = prec
		z.acc = Exact
		z.form = zero
		f = z
		return
	}
	// len(z.mant) > 0

	// The mantissa may have a radix point (fcount <= 0) and there
	// may be a nonzero exponent exp. The radix point amounts to a
	// division by b**(-fcount). An exponent means multiplication by
	// ebase**exp. Finally, mantissa normalization (shift left) requires
	// a correcting multiplication by 2**(-shiftcount). Multiplications
	// are commutative, so we can apply them in any order as long as there
	// is no loss of precision. We only have powers of 2 and 10, and
	// we split powers of 10 into the product of the same powers of
	// 2 and 5. This reduces the size of the multiplication factor
	// needed for base-10 exponents.

	// normalize mantissa and determine initial exponent contributions
	exp2 := int64(len(z.mant))*_W - fnorm(z.mant)
	exp5 := int64(0)

	// determine binary or decimal exponent contribution of radix point
	if fcount < 0 {
		// The mantissa has a radix point ddd.dddd; and
		// -fcount is the number of digits to the right
		// of '.'. Adjust relevant exponent accordingly.
		d := int64(fcount)
		switch b {
		case 10:
			exp5 = d
			fallthrough // 10**e == 5**e * 2**e
		case 2:
			exp2 += d
		case 8:
			exp2 += d * 3 // octal digits are 3 bits each
		case 16:
			exp2 += d * 4 // hexadecimal digits are 4 bits each
		default:
			panic("unexpected mantissa base")
		}
		// fcount consumed - not needed anymore
	}

	// take actual exponent into account
	switch ebase {
	case 10:
		exp5 += exp
		fallthrough // see fallthrough above
	case 2:
		exp2 += exp
	default:
		panic("unexpected exponent base")
	}
	// exp consumed - not needed anymore

	// apply 2**exp2
	if MinExp <= exp2 && exp2 <= MaxExp {
		z.prec = prec
		z.form = finite
		z.exp = int32(exp2)
		f = z
	} else {
		err = fmt.Errorf("exponent overflow")
		return
	}

	if exp5 == 0 {
		// no decimal exponent contribution
		z.round(0)
		return
	}
	// exp5 != 0

	// apply 5**exp5
	p := new(Float).SetPrec(z.Prec() + 64) // use more bits for p -- TODO(gri) what is the right number?
	if exp5 < 0 {
		z.Quo(z, p.pow5(uint64(-exp5)))
	} else {
		z.Mul(z, p.pow5(uint64(exp5)))
	}

	return
}

// These powers of 5 fit into a uint64.
//
//	for p, q := uint64(0), uint64(1); p < q; p, q = q, q*5 {
//		fmt.Println(q)
//	}
//
var pow5tab = [...]uint64{
	1,
	5,
	25,
	125,
	625,
	3125,
	15625,
	78125,
	390625,
	1953125,
	9765625,
	48828125,
	244140625,
	1220703125,
	6103515625,
	30517578125,
	152587890625,
	762939453125,
	3814697265625,
	19073486328125,
	95367431640625,
	476837158203125,
	2384185791015625,
	11920928955078125,
	59604644775390625,
	298023223876953125,
	1490116119384765625,
	7450580596923828125,
}

// pow5 sets z to 5**n and returns z.
// n must not be negative.
func (z *Float) pow5(n uint64) *Float {
	const m = uint64(len(pow5tab) - 1)
	if n <= m {
		return z.SetUint64(pow5tab[n])
	}
	// n > m

	z.SetUint64(pow5tab[m])
	n -= m

	// use more bits for f than for z
	// TODO(gri) what is the right number?
	f := new(Float).SetPrec(z.Prec() + 64).SetUint64(5)

	for n > 0 {
		if n&1 != 0 {
			z.Mul(z, f)
		}
		f.Mul(f, f)
		n >>= 1
	}

	return z
}

// Parse parses s which must contain a text representation of a floating-
// point number with a mantissa in the given conversion base (the exponent
// is always a decimal number), or a string representing an infinite value.
//
// For base 0, an underscore character ``_'' may appear between a base
// prefix and an adjacent digit, and between successive digits; such
// underscores do not change the value of the number, or the returned
// digit count. Incorrect placement of underscores is reported as an
// error if there are no other errors. If base != 0, underscores are
// not recognized and thus terminate scanning like any other character
// that is not a valid radix point or digit.
//
// It sets z to the (possibly rounded) value of the corresponding floating-
// point value, and returns z, the actual base b, and an error err, if any.
// The entire string (not just a prefix) must be consumed for success.
// If z's precision is 0, it is changed to 64 before rounding takes effect.
// The number must be of the form:
//
//     number    = [ sign ] ( float | "inf" | "Inf" ) .
//     sign      = "+" | "-" .
//     float     = ( mantissa | prefix pmantissa ) [ exponent ] .
//     prefix    = "0" [ "b" | "B" | "o" | "O" | "x" | "X" ] .
//     mantissa  = digits "." [ digits ] | digits | "." digits .
//     pmantissa = [ "_" ] digits "." [ digits ] | [ "_" ] digits | "." digits .
//     exponent  = ( "e" | "E" | "p" | "P" ) [ sign ] digits .
//     digits    = digit { [ "_" ] digit } .
//     digit     = "0" ... "9" | "a" ... "z" | "A" ... "Z" .
//
// The base argument must be 0, 2, 8, 10, or 16. Providing an invalid base
// argument will lead to a run-time panic.
//
// For base 0, the number prefix determines the actual base: A prefix of
// ``0b'' or ``0B'' selects base 2, ``0o'' or ``0O'' selects base 8, and
// ``0x'' or ``0X'' selects base 16. Otherwise, the actual base is 10 and
// no prefix is accepted. The octal prefix "0" is not supported (a leading
// "0" is simply considered a "0").
//
// A "p" or "P" exponent indicates a base 2 (rather then base 10) exponent;
// for instance, "0x1.fffffffffffffp1023" (using base 0) represents the
// maximum float64 value. For hexadecimal mantissae, the exponent character
// must be one of 'p' or 'P', if present (an "e" or "E" exponent indicator
// cannot be distinguished from a mantissa digit).
//
// The returned *Float f is nil and the value of z is valid but not
// defined if an error is reported.
//
func (z *Float) Parse(s string, base int) (f *Float, b int, err error) {
	// scan doesn't handle ±Inf
	if len(s) == 3 && (s == "Inf" || s == "inf") {
		f = z.SetInf(false)
		return
	}
	if len(s) == 4 && (s[0] == '+' || s[0] == '-') && (s[1:] == "Inf" || s[1:] == "inf") {
		f = z.SetInf(s[0] == '-')
		return
	}

	r := strings.NewReader(s)
	if f, b, err = z.scan(r, base); err != nil {
		return
	}

	// entire string must have been consumed
	if ch, err2 := r.ReadByte(); err2 == nil {
		err = fmt.Errorf("expected end of string, found %q", ch)
	} else if err2 != io.EOF {
		err = err2
	}

	return
}

// ParseFloat is like f.Parse(s, base) with f set to the given precision
// and rounding mode.
func ParseFloat(s string, base int, prec uint, mode RoundingMode) (f *Float, b int, err error) {
	return new(Float).SetPrec(prec).SetMode(mode).Parse(s, base)
}

var _ fmt.Scanner = (*Float)(nil) // *Float must implement fmt.Scanner

// Scan is a support routine for fmt.Scanner; it sets z to the value of
// the scanned number. It accepts formats whose verbs are supported by
// fmt.Scan for floating point values, which are:
// 'b' (binary), 'e', 'E', 'f', 'F', 'g' and 'G'.
// Scan doesn't handle ±Inf.
func (z *Float) Scan(s fmt.ScanState, ch rune) error {
	s.SkipSpace()
	_, _, err := z.scan(byteReader{s}, 0)
	return err
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// math/big/ftoa.go
// Text converts the floating-point number x to a string according
// to the given format and precision prec. The format is one of:
//
//	'e'	-d.dddde±dd, decimal exponent, at least two (possibly 0) exponent digits
//	'E'	-d.ddddE±dd, decimal exponent, at least two (possibly 0) exponent digits
//	'f'	-ddddd.dddd, no exponent
//	'g'	like 'e' for large exponents, like 'f' otherwise
//	'G'	like 'E' for large exponents, like 'f' otherwise
//	'x'	-0xd.dddddp±dd, hexadecimal mantissa, decimal power of two exponent
//	'p'	-0x.dddp±dd, hexadecimal mantissa, decimal power of two exponent (non-standard)
//	'b'	-ddddddp±dd, decimal mantissa, decimal power of two exponent (non-standard)
//
// For the power-of-two exponent formats, the mantissa is printed in normalized form:
//
//	'x'	hexadecimal mantissa in [1, 2), or 0
//	'p'	hexadecimal mantissa in [½, 1), or 0
//	'b'	decimal integer mantissa using x.Prec() bits, or 0
//
// Note that the 'x' form is the one used by most other languages and libraries.
//
// If format is a different character, Text returns a "%" followed by the
// unrecognized format character.
//
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', 'G', and 'x' formats.
// For 'e', 'E', 'f', and 'x', it is the number of digits after the decimal point.
// For 'g' and 'G' it is the total number of digits. A negative precision selects
// the smallest number of decimal digits necessary to identify the value x uniquely
// using x.Prec() mantissa bits.
// The prec value is ignored for the 'b' and 'p' formats.
func (x *Float) Text(format byte, prec int) string {
	cap := 10 // TODO(gri) determine a good/better value here
	if prec > 0 {
		cap += prec
	}
	return string(x.Append(make([]byte, 0, cap), format, prec))
}

// String formats x like x.Text('g', 10).
// (String must be called explicitly, Float.Format does not support %s verb.)
func (x *Float) String() string {
	return x.Text('g', 10)
}

// Append appends to buf the string form of the floating-point number x,
// as generated by x.Text, and returns the extended buffer.
func (x *Float) Append(buf []byte, fmt byte, prec int) []byte {
	// sign
	if x.neg {
		buf = append(buf, '-')
	}

	// Inf
	if x.form == inf {
		if !x.neg {
			buf = append(buf, '+')
		}
		return append(buf, "Inf"...)
	}

	// pick off easy formats
	switch fmt {
	case 'b':
		return x.fmtB(buf)
	case 'p':
		return x.fmtP(buf)
	case 'x':
		return x.fmtX(buf, prec)
	}

	// Algorithm:
	//   1) convert Float to multiprecision decimal
	//   2) round to desired precision
	//   3) read digits out and format

	// 1) convert Float to multiprecision decimal
	var d decimal // == 0.0
	if x.form == finite {
		// x != 0
		d.init(x.mant, int(x.exp)-x.mant.bitLen())
	}

	// 2) round to desired precision
	shortest := false
	if prec < 0 {
		shortest = true
		roundShortest(&d, x)
		// Precision for shortest representation mode.
		switch fmt {
		case 'e', 'E':
			prec = len(d.mant) - 1
		case 'f':
			prec = max(len(d.mant)-d.exp, 0)
		case 'g', 'G':
			prec = len(d.mant)
		}
	} else {
		// round appropriately
		switch fmt {
		case 'e', 'E':
			// one digit before and number of digits after decimal point
			d.round(1 + prec)
		case 'f':
			// number of digits before and after decimal point
			d.round(d.exp + prec)
		case 'g', 'G':
			if prec == 0 {
				prec = 1
			}
			d.round(prec)
		}
	}

	// 3) read digits out and format
	switch fmt {
	case 'e', 'E':
		return fmtE(buf, fmt, prec, d)
	case 'f':
		return fmtF(buf, prec, d)
	case 'g', 'G':
		// trim trailing fractional zeros in %e format
		eprec := prec
		if eprec > len(d.mant) && len(d.mant) >= d.exp {
			eprec = len(d.mant)
		}
		// %e is used if the exponent from the conversion
		// is less than -4 or greater than or equal to the precision.
		// If precision was the shortest possible, use eprec = 6 for
		// this decision.
		if shortest {
			eprec = 6
		}
		exp := d.exp - 1
		if exp < -4 || exp >= eprec {
			if prec > len(d.mant) {
				prec = len(d.mant)
			}
			return fmtE(buf, fmt+'e'-'g', prec-1, d)
		}
		if prec > d.exp {
			prec = len(d.mant)
		}
		return fmtF(buf, max(prec-d.exp, 0), d)
	}

	// unknown format
	if x.neg {
		buf = buf[:len(buf)-1] // sign was added prematurely - remove it again
	}
	return append(buf, '%', fmt)
}

func roundShortest(d *decimal, x *Float) {
	// if the mantissa is zero, the number is zero - stop now
	if len(d.mant) == 0 {
		return
	}

	// Approach: All numbers in the interval [x - 1/2ulp, x + 1/2ulp]
	// (possibly exclusive) round to x for the given precision of x.
	// Compute the lower and upper bound in decimal form and find the
	// shortest decimal number d such that lower <= d <= upper.

	// TODO(gri) strconv/ftoa.do describes a shortcut in some cases.
	// See if we can use it (in adjusted form) here as well.

	// 1) Compute normalized mantissa mant and exponent exp for x such
	// that the lsb of mant corresponds to 1/2 ulp for the precision of
	// x (i.e., for mant we want x.prec + 1 bits).
	mant := nat(nil).set(x.mant)
	exp := int(x.exp) - mant.bitLen()
	s := mant.bitLen() - int(x.prec+1)
	switch {
	case s < 0:
		mant = mant.shl(mant, uint(-s))
	case s > 0:
		mant = mant.shr(mant, uint(+s))
	}
	exp += s
	// x = mant * 2**exp with lsb(mant) == 1/2 ulp of x.prec

	// 2) Compute lower bound by subtracting 1/2 ulp.
	var lower decimal
	var tmp nat
	lower.init(tmp.sub(mant, natOne), exp)

	// 3) Compute upper bound by adding 1/2 ulp.
	var upper decimal
	upper.init(tmp.add(mant, natOne), exp)

	// The upper and lower bounds are possible outputs only if
	// the original mantissa is even, so that ToNearestEven rounding
	// would round to the original mantissa and not the neighbors.
	inclusive := mant[0]&2 == 0 // test bit 1 since original mantissa was shifted by 1

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for i, m := range d.mant {
		l := lower.at(i)
		u := upper.at(i)

		// Okay to round down (truncate) if lower has a different digit
		// or if lower is inclusive and is exactly the result of rounding
		// down (i.e., and we have reached the final digit of lower).
		okdown := l != m || inclusive && i+1 == len(lower.mant)

		// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		okup := m != u && (inclusive || m+1 < u || i+1 < len(upper.mant))

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		switch {
		case okdown && okup:
			d.round(i + 1)
			return
		case okdown:
			d.roundDown(i + 1)
			return
		case okup:
			d.roundUp(i + 1)
			return
		}
	}
}

// %e: d.ddddde±dd
func fmtE(buf []byte, fmt byte, prec int, d decimal) []byte {
	// first digit
	ch := byte('0')
	if len(d.mant) > 0 {
		ch = d.mant[0]
	}
	buf = append(buf, ch)

	// .moredigits
	if prec > 0 {
		buf = append(buf, '.')
		i := 1
		m := min(len(d.mant), prec+1)
		if i < m {
			buf = append(buf, d.mant[i:m]...)
			i = m
		}
		for ; i <= prec; i++ {
			buf = append(buf, '0')
		}
	}

	// e±
	buf = append(buf, fmt)
	var exp int64
	if len(d.mant) > 0 {
		exp = int64(d.exp) - 1 // -1 because first digit was printed before '.'
	}
	if exp < 0 {
		ch = '-'
		exp = -exp
	} else {
		ch = '+'
	}
	buf = append(buf, ch)

	// dd...d
	if exp < 10 {
		buf = append(buf, '0') // at least 2 exponent digits
	}
	return strconv.AppendInt(buf, exp, 10)
}

// %f: ddddddd.ddddd
func fmtF(buf []byte, prec int, d decimal) []byte {
	// integer, padded with zeros as needed
	if d.exp > 0 {
		m := min(len(d.mant), d.exp)
		buf = append(buf, d.mant[:m]...)
		for ; m < d.exp; m++ {
			buf = append(buf, '0')
		}
	} else {
		buf = append(buf, '0')
	}

	// fraction
	if prec > 0 {
		buf = append(buf, '.')
		for i := 0; i < prec; i++ {
			buf = append(buf, d.at(d.exp+i))
		}
	}

	return buf
}

// fmtB appends the string of x in the format mantissa "p" exponent
// with a decimal mantissa and a binary exponent, or 0" if x is zero,
// and returns the extended buffer.
// The mantissa is normalized such that is uses x.Prec() bits in binary
// representation.
// The sign of x is ignored, and x must not be an Inf.
// (The caller handles Inf before invoking fmtB.)
func (x *Float) fmtB(buf []byte) []byte {
	if x.form == zero {
		return append(buf, '0')
	}

	if debugFloat && x.form != finite {
		panic("non-finite float")
	}
	// x != 0

	// adjust mantissa to use exactly x.prec bits
	m := x.mant
	switch w := uint32(len(x.mant)) * _W; {
	case w < x.prec:
		m = nat(nil).shl(m, uint(x.prec-w))
	case w > x.prec:
		m = nat(nil).shr(m, uint(w-x.prec))
	}

	buf = append(buf, m.utoa(10)...)
	buf = append(buf, 'p')
	e := int64(x.exp) - int64(x.prec)
	if e >= 0 {
		buf = append(buf, '+')
	}
	return strconv.AppendInt(buf, e, 10)
}

// fmtX appends the string of x in the format "0x1." mantissa "p" exponent
// with a hexadecimal mantissa and a binary exponent, or "0x0p0" if x is zero,
// and returns the extended buffer.
// A non-zero mantissa is normalized such that 1.0 <= mantissa < 2.0.
// The sign of x is ignored, and x must not be an Inf.
// (The caller handles Inf before invoking fmtX.)
func (x *Float) fmtX(buf []byte, prec int) []byte {
	if x.form == zero {
		buf = append(buf, "0x0"...)
		if prec > 0 {
			buf = append(buf, '.')
			for i := 0; i < prec; i++ {
				buf = append(buf, '0')
			}
		}
		buf = append(buf, "p+00"...)
		return buf
	}

	if debugFloat && x.form != finite {
		panic("non-finite float")
	}

	// round mantissa to n bits
	var n uint
	if prec < 0 {
		n = 1 + (x.MinPrec()-1+3)/4*4 // round MinPrec up to 1 mod 4
	} else {
		n = 1 + 4*uint(prec)
	}
	// n%4 == 1
	x = new(Float).SetPrec(n).SetMode(x.mode).Set(x)

	// adjust mantissa to use exactly n bits
	m := x.mant
	switch w := uint(len(x.mant)) * _W; {
	case w < n:
		m = nat(nil).shl(m, n-w)
	case w > n:
		m = nat(nil).shr(m, w-n)
	}
	exp64 := int64(x.exp) - 1 // avoid wrap-around

	hm := m.utoa(16)
	if debugFloat && hm[0] != '1' {
		panic("incorrect mantissa: " + string(hm))
	}
	buf = append(buf, "0x1"...)
	if len(hm) > 1 {
		buf = append(buf, '.')
		buf = append(buf, hm[1:]...)
	}

	buf = append(buf, 'p')
	if exp64 >= 0 {
		buf = append(buf, '+')
	} else {
		exp64 = -exp64
		buf = append(buf, '-')
	}
	// Force at least two exponent digits, to match fmt.
	if exp64 < 10 {
		buf = append(buf, '0')
	}
	return strconv.AppendInt(buf, exp64, 10)
}

// fmtP appends the string of x in the format "0x." mantissa "p" exponent
// with a hexadecimal mantissa and a binary exponent, or "0" if x is zero,
// and returns the extended buffer.
// The mantissa is normalized such that 0.5 <= 0.mantissa < 1.0.
// The sign of x is ignored, and x must not be an Inf.
// (The caller handles Inf before invoking fmtP.)
func (x *Float) fmtP(buf []byte) []byte {
	if x.form == zero {
		return append(buf, '0')
	}

	if debugFloat && x.form != finite {
		panic("non-finite float")
	}
	// x != 0

	// remove trailing 0 words early
	// (no need to convert to hex 0's and trim later)
	m := x.mant
	i := 0
	for i < len(m) && m[i] == 0 {
		i++
	}
	m = m[i:]

	buf = append(buf, "0x."...)
	buf = append(buf, bytes.TrimRight(m.utoa(16), "0")...)
	buf = append(buf, 'p')
	if x.exp >= 0 {
		buf = append(buf, '+')
	}
	return strconv.AppendInt(buf, int64(x.exp), 10)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var _ fmt.Formatter = &floatZero // *Float must implement fmt.Formatter

// Format implements fmt.Formatter. It accepts all the regular
// formats for floating-point numbers ('b', 'e', 'E', 'f', 'F',
// 'g', 'G', 'x') as well as 'p' and 'v'. See (*Float).Text for the
// interpretation of 'p'. The 'v' format is handled like 'g'.
// Format also supports specification of the minimum precision
// in digits, the output field width, as well as the format flags
// '+' and ' ' for sign control, '0' for space or zero padding,
// and '-' for left or right justification. See the fmt package
// for details.
func (x *Float) Format(s fmt.State, format rune) {
	prec, hasPrec := s.Precision()
	if !hasPrec {
		prec = 6 // default precision for 'e', 'f'
	}

	switch format {
	case 'e', 'E', 'f', 'b', 'p', 'x':
		// nothing to do
	case 'F':
		// (*Float).Text doesn't support 'F'; handle like 'f'
		format = 'f'
	case 'v':
		// handle like 'g'
		format = 'g'
		fallthrough
	case 'g', 'G':
		if !hasPrec {
			prec = -1 // default precision for 'g', 'G'
		}
	default:
		fmt.Fprintf(s, "%%!%c(*big.Float=%s)", format, x.String())
		return
	}
	var buf []byte
	buf = x.Append(buf, byte(format), prec)
	if len(buf) == 0 {
		buf = []byte("?") // should never happen, but don't crash
	}
	// len(buf) > 0

	var sign string
	switch {
	case buf[0] == '-':
		sign = "-"
		buf = buf[1:]
	case buf[0] == '+':
		// +Inf
		sign = "+"
		if s.Flag(' ') {
			sign = " "
		}
		buf = buf[1:]
	case s.Flag('+'):
		sign = "+"
	case s.Flag(' '):
		sign = " "
	}

	var padding int
	if width, hasWidth := s.Width(); hasWidth && width > len(sign)+len(buf) {
		padding = width - len(sign) - len(buf)
	}

	switch {
	case s.Flag('0') && !x.IsInf():
		// 0-padding on left
		writeMultiple(s, sign, 1)
		writeMultiple(s, "0", padding)
		s.Write(buf)
	case s.Flag('-'):
		// padding on right
		writeMultiple(s, sign, 1)
		s.Write(buf)
		writeMultiple(s, " ", padding)
	default:
		// padding on left
		writeMultiple(s, " ", padding)
		writeMultiple(s, sign, 1)
		s.Write(buf)
	}
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements signed multi-precision integers.

// An Int represents a signed multi-precision integer.
// The zero value for an Int represents the value 0.
//
// Operations always take pointer arguments (*Int) rather
// than Int values, and each unique Int value requires
// its own unique *Int pointer. To "copy" an Int value,
// an existing (or newly allocated) Int must be set to
// a new value using the Int.Set method; shallow copies
// of Ints are not supported and may lead to errors.
type Int struct {
	neg bool // sign
	abs nat  // absolute value of the integer
}

var intOne = &Int{false, natOne}

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
//
func (x *Int) Sign() int {
	if len(x.abs) == 0 {
		return 0
	}
	if x.neg {
		return -1
	}
	return 1
}

// SetInt64 sets z to x and returns z.
func (z *Int) SetInt64(x int64) *Int {
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	z.abs = z.abs.setUint64(uint64(x))
	z.neg = neg
	return z
}

// SetUint64 sets z to x and returns z.
func (z *Int) SetUint64(x uint64) *Int {
	z.abs = z.abs.setUint64(x)
	z.neg = false
	return z
}

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *Int {
	return new(Int).SetInt64(x)
}

// Set sets z to x and returns z.
func (z *Int) Set(x *Int) *Int {
	if z != x {
		z.abs = z.abs.set(x.abs)
		z.neg = x.neg
	}
	return z
}

// Bits provides raw (unchecked but fast) access to x by returning its
// absolute value as a little-endian Word slice. The result and x share
// the same underlying array.
// Bits is intended to support implementation of missing low-level Int
// functionality outside this package; it should be avoided otherwise.
func (x *Int) Bits() []Word {
	return x.abs
}

// SetBits provides raw (unchecked but fast) access to z by setting its
// value to abs, interpreted as a little-endian Word slice, and returning
// z. The result and abs share the same underlying array.
// SetBits is intended to support implementation of missing low-level Int
// functionality outside this package; it should be avoided otherwise.
func (z *Int) SetBits(abs []Word) *Int {
	z.abs = nat(abs).norm()
	z.neg = false
	return z
}

// Abs sets z to |x| (the absolute value of x) and returns z.
func (z *Int) Abs(x *Int) *Int {
	z.Set(x)
	z.neg = false
	return z
}

// Neg sets z to -x and returns z.
func (z *Int) Neg(x *Int) *Int {
	z.Set(x)
	z.neg = len(z.abs) > 0 && !z.neg // 0 has no sign
	return z
}

// Add sets z to the sum x+y and returns z.
func (z *Int) Add(x, y *Int) *Int {
	neg := x.neg
	if x.neg == y.neg {
		// x + y == x + y
		// (-x) + (-y) == -(x + y)
		z.abs = z.abs.add(x.abs, y.abs)
	} else {
		// x + (-y) == x - y == -(y - x)
		// (-x) + y == y - x == -(x - y)
		if x.abs.cmp(y.abs) >= 0 {
			z.abs = z.abs.sub(x.abs, y.abs)
		} else {
			neg = !neg
			z.abs = z.abs.sub(y.abs, x.abs)
		}
	}
	z.neg = len(z.abs) > 0 && neg // 0 has no sign
	return z
}

// Sub sets z to the difference x-y and returns z.
func (z *Int) Sub(x, y *Int) *Int {
	neg := x.neg
	if x.neg != y.neg {
		// x - (-y) == x + y
		// (-x) - y == -(x + y)
		z.abs = z.abs.add(x.abs, y.abs)
	} else {
		// x - y == x - y == -(y - x)
		// (-x) - (-y) == y - x == -(x - y)
		if x.abs.cmp(y.abs) >= 0 {
			z.abs = z.abs.sub(x.abs, y.abs)
		} else {
			neg = !neg
			z.abs = z.abs.sub(y.abs, x.abs)
		}
	}
	z.neg = len(z.abs) > 0 && neg // 0 has no sign
	return z
}

// Mul sets z to the product x*y and returns z.
func (z *Int) Mul(x, y *Int) *Int {
	// x * y == x * y
	// x * (-y) == -(x * y)
	// (-x) * y == -(x * y)
	// (-x) * (-y) == x * y
	if x == y {
		z.abs = z.abs.sqr(x.abs)
		z.neg = false
		return z
	}
	z.abs = z.abs.mul(x.abs, y.abs)
	z.neg = len(z.abs) > 0 && x.neg != y.neg // 0 has no sign
	return z
}

// MulRange sets z to the product of all integers
// in the range [a, b] inclusively and returns z.
// If a > b (empty range), the result is 1.
func (z *Int) MulRange(a, b int64) *Int {
	switch {
	case a > b:
		return z.SetInt64(1) // empty range
	case a <= 0 && b >= 0:
		return z.SetInt64(0) // range includes 0
	}
	// a <= b && (b < 0 || a > 0)

	neg := false
	if a < 0 {
		neg = (b-a)&1 == 0
		a, b = -b, -a
	}

	z.abs = z.abs.mulRange(uint64(a), uint64(b))
	z.neg = neg
	return z
}

// Binomial sets z to the binomial coefficient of (n, k) and returns z.
func (z *Int) Binomial(n, k int64) *Int {
	// reduce the number of multiplications by reducing k
	if n/2 < k && k <= n {
		k = n - k // Binomial(n, k) == Binomial(n, n-k)
	}
	var a, b Int
	a.MulRange(n-k+1, n)
	b.MulRange(1, k)
	return z.Quo(&a, &b)
}

// Quo sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Quo implements truncated division (like Go); see QuoRem for more details.
func (z *Int) Quo(x, y *Int) *Int {
	z.abs, _ = z.abs.div(nil, x.abs, y.abs)
	z.neg = len(z.abs) > 0 && x.neg != y.neg // 0 has no sign
	return z
}

// Rem sets z to the remainder x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Rem implements truncated modulus (like Go); see QuoRem for more details.
func (z *Int) Rem(x, y *Int) *Int {
	_, z.abs = nat(nil).div(z.abs, x.abs, y.abs)
	z.neg = len(z.abs) > 0 && x.neg // 0 has no sign
	return z
}

// QuoRem sets z to the quotient x/y and r to the remainder x%y
// and returns the pair (z, r) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// QuoRem implements T-division and modulus (like Go):
//
//	q = x/y      with the result truncated to zero
//	r = x - y*q
//
// (See Daan Leijen, ``Division and Modulus for Computer Scientists''.)
// See DivMod for Euclidean division and modulus (unlike Go).
//
func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int) {
	z.abs, r.abs = z.abs.div(r.abs, x.abs, y.abs)
	z.neg, r.neg = len(z.abs) > 0 && x.neg != y.neg, len(r.abs) > 0 && x.neg // 0 has no sign
	return z, r
}

// Div sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see DivMod for more details.
func (z *Int) Div(x, y *Int) *Int {
	y_neg := y.neg // z may be an alias for y
	var r Int
	z.QuoRem(x, y, &r)
	if r.neg {
		if y_neg {
			z.Add(z, intOne)
		} else {
			z.Sub(z, intOne)
		}
	}
	return z
}

// Mod sets z to the modulus x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean modulus (unlike Go); see DivMod for more details.
func (z *Int) Mod(x, y *Int) *Int {
	y0 := y // save y
	if z == y || alias(z.abs, y.abs) {
		y0 = new(Int).Set(y)
	}
	var q Int
	q.QuoRem(x, y, z)
	if z.neg {
		if y0.neg {
			z.Sub(z, y0)
		} else {
			z.Add(z, y0)
		}
	}
	return z
}

// DivMod sets z to the quotient x div y and m to the modulus x mod y
// and returns the pair (z, m) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// DivMod implements Euclidean division and modulus (unlike Go):
//
//	q = x div y  such that
//	m = x - y*q  with 0 <= m < |y|
//
// (See Raymond T. Boute, ``The Euclidean definition of the functions
// div and mod''. ACM Transactions on Programming Languages and
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
// ACM press.)
// See QuoRem for T-division and modulus (like Go).
//
func (z *Int) DivMod(x, y, m *Int) (*Int, *Int) {
	y0 := y // save y
	if z == y || alias(z.abs, y.abs) {
		y0 = new(Int).Set(y)
	}
	z.QuoRem(x, y, m)
	if m.neg {
		if y0.neg {
			z.Add(z, intOne)
			m.Sub(m, y0)
		} else {
			z.Sub(z, intOne)
			m.Add(m, y0)
		}
	}
	return z, m
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func (x *Int) Cmp(y *Int) (r int) {
	// x cmp y == x cmp y
	// x cmp (-y) == x
	// (-x) cmp y == y
	// (-x) cmp (-y) == -(x cmp y)
	switch {
	case x == y:
		// nothing to do
	case x.neg == y.neg:
		r = x.abs.cmp(y.abs)
		if x.neg {
			r = -r
		}
	case x.neg:
		r = -1
	default:
		r = 1
	}
	return
}

// CmpAbs compares the absolute values of x and y and returns:
//
//   -1 if |x| <  |y|
//    0 if |x| == |y|
//   +1 if |x| >  |y|
//
func (x *Int) CmpAbs(y *Int) int {
	return x.abs.cmp(y.abs)
}

// low32 returns the least significant 32 bits of x.
func low32(x nat) uint32 {
	if len(x) == 0 {
		return 0
	}
	return uint32(x[0])
}

// low64 returns the least significant 64 bits of x.
func low64(x nat) uint64 {
	if len(x) == 0 {
		return 0
	}
	v := uint64(x[0])
	if _W == 32 && len(x) > 1 {
		return uint64(x[1])<<32 | v
	}
	return v
}

// Int64 returns the int64 representation of x.
// If x cannot be represented in an int64, the result is undefined.
func (x *Int) Int64() int64 {
	v := int64(low64(x.abs))
	if x.neg {
		v = -v
	}
	return v
}

// Uint64 returns the uint64 representation of x.
// If x cannot be represented in a uint64, the result is undefined.
func (x *Int) Uint64() uint64 {
	return low64(x.abs)
}

// IsInt64 reports whether x can be represented as an int64.
func (x *Int) IsInt64() bool {
	if len(x.abs) <= 64/_W {
		w := int64(low64(x.abs))
		return w >= 0 || x.neg && w == -w
	}
	return false
}

// IsUint64 reports whether x can be represented as a uint64.
func (x *Int) IsUint64() bool {
	return !x.neg && len(x.abs) <= 64/_W
}

// SetString sets z to the value of s, interpreted in the given base,
// and returns z and a boolean indicating success. The entire string
// (not just a prefix) must be valid for success. If SetString fails,
// the value of z is undefined but the returned value is nil.
//
// The base argument must be 0 or a value between 2 and MaxBase.
// For base 0, the number prefix determines the actual base: A prefix of
// ``0b'' or ``0B'' selects base 2, ``0'', ``0o'' or ``0O'' selects base 8,
// and ``0x'' or ``0X'' selects base 16. Otherwise, the selected base is 10
// and no prefix is accepted.
//
// For bases <= 36, lower and upper case letters are considered the same:
// The letters 'a' to 'z' and 'A' to 'Z' represent digit values 10 to 35.
// For bases > 36, the upper case letters 'A' to 'Z' represent the digit
// values 36 to 61.
//
// For base 0, an underscore character ``_'' may appear between a base
// prefix and an adjacent digit, and between successive digits; such
// underscores do not change the value of the number.
// Incorrect placement of underscores is reported as an error if there
// are no other errors. If base != 0, underscores are not recognized
// and act like any other character that is not a valid digit.
//
func (z *Int) SetString(s string, base int) (*Int, bool) {
	return z.setFromScanner(strings.NewReader(s), base)
}

// setFromScanner implements SetString given an io.BytesScanner.
// For documentation see comments of SetString.
func (z *Int) setFromScanner(r io.ByteScanner, base int) (*Int, bool) {
	if _, _, err := z.scan(r, base); err != nil {
		return nil, false
	}
	// entire content must have been consumed
	if _, err := r.ReadByte(); err != io.EOF {
		return nil, false
	}
	return z, true // err == io.EOF => scan consumed all content of r
}

// SetBytes interprets buf as the bytes of a big-endian unsigned
// integer, sets z to that value, and returns z.
func (z *Int) SetBytes(buf []byte) *Int {
	z.abs = z.abs.setBytes(buf)
	z.neg = false
	return z
}

// Bytes returns the absolute value of x as a big-endian byte slice.
//
// To use a fixed length slice, or a preallocated one, use FillBytes.
func (x *Int) Bytes() []byte {
	buf := make([]byte, len(x.abs)*_S)
	return buf[x.abs.bytes(buf):]
}

// FillBytes sets buf to the absolute value of x, storing it as a zero-extended
// big-endian byte slice, and returns buf.
//
// If the absolute value of x doesn't fit in buf, FillBytes will panic.
func (x *Int) FillBytes(buf []byte) []byte {
	// Clear whole buffer. (This gets optimized into a memclr.)
	for i := range buf {
		buf[i] = 0
	}
	x.abs.bytes(buf)
	return buf
}

// BitLen returns the length of the absolute value of x in bits.
// The bit length of 0 is 0.
func (x *Int) BitLen() int {
	return x.abs.bitLen()
}

// TrailingZeroBits returns the number of consecutive least significant zero
// bits of |x|.
func (x *Int) TrailingZeroBits() uint {
	return x.abs.trailingZeroBits()
}

// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
// If m == nil or m == 0, z = x**y unless y <= 0 then z = 1. If m != 0, y < 0,
// and x and m are not relatively prime, z is unchanged and nil is returned.
//
// Modular exponentiation of inputs of a particular size is not a
// cryptographically constant-time operation.
func (z *Int) Exp(x, y, m *Int) *Int {
	// See Knuth, volume 2, section 4.6.3.
	xWords := x.abs
	if y.neg {
		if m == nil || len(m.abs) == 0 {
			return z.SetInt64(1)
		}
		// for y < 0: x**y mod m == (x**(-1))**|y| mod m
		inverse := new(Int).ModInverse(x, m)
		if inverse == nil {
			return nil
		}
		xWords = inverse.abs
	}
	yWords := y.abs

	var mWords nat
	if m != nil {
		mWords = m.abs // m.abs may be nil for m == 0
	}

	z.abs = z.abs.expNN(xWords, yWords, mWords)
	z.neg = len(z.abs) > 0 && x.neg && len(yWords) > 0 && yWords[0]&1 == 1 // 0 has no sign
	if z.neg && len(mWords) > 0 {
		// make modulus result positive
		z.abs = z.abs.sub(mWords, z.abs) // z == x**y mod |m| && 0 <= z < |m|
		z.neg = false
	}

	return z
}

// GCD sets z to the greatest common divisor of a and b and returns z.
// If x or y are not nil, GCD sets their value such that z = a*x + b*y.
//
// a and b may be positive, zero or negative. (Before Go 1.14 both had
// to be > 0.) Regardless of the signs of a and b, z is always >= 0.
//
// If a == b == 0, GCD sets z = x = y = 0.
//
// If a == 0 and b != 0, GCD sets z = |b|, x = 0, y = sign(b) * 1.
//
// If a != 0 and b == 0, GCD sets z = |a|, x = sign(a) * 1, y = 0.
func (z *Int) GCD(x, y, a, b *Int) *Int {
	if len(a.abs) == 0 || len(b.abs) == 0 {
		lenA, lenB, negA, negB := len(a.abs), len(b.abs), a.neg, b.neg
		if lenA == 0 {
			z.Set(b)
		} else {
			z.Set(a)
		}
		z.neg = false
		if x != nil {
			if lenA == 0 {
				x.SetUint64(0)
			} else {
				x.SetUint64(1)
				x.neg = negA
			}
		}
		if y != nil {
			if lenB == 0 {
				y.SetUint64(0)
			} else {
				y.SetUint64(1)
				y.neg = negB
			}
		}
		return z
	}

	return z.lehmerGCD(x, y, a, b)
}

// lehmerSimulate attempts to simulate several Euclidean update steps
// using the leading digits of A and B.  It returns u0, u1, v0, v1
// such that A and B can be updated as:
//		A = u0*A + v0*B
//		B = u1*A + v1*B
// Requirements: A >= B and len(B.abs) >= 2
// Since we are calculating with full words to avoid overflow,
// we use 'even' to track the sign of the cosequences.
// For even iterations: u0, v1 >= 0 && u1, v0 <= 0
// For odd  iterations: u0, v1 <= 0 && u1, v0 >= 0
func lehmerSimulate(A, B *Int) (u0, u1, v0, v1 Word, even bool) {
	// initialize the digits
	var a1, a2, u2, v2 Word

	m := len(B.abs) // m >= 2
	n := len(A.abs) // n >= m >= 2

	// extract the top Word of bits from A and B
	h := nlz(A.abs[n-1])
	a1 = A.abs[n-1]<<h | A.abs[n-2]>>(_W-h)
	// B may have implicit zero words in the high bits if the lengths differ
	switch {
	case n == m:
		a2 = B.abs[n-1]<<h | B.abs[n-2]>>(_W-h)
	case n == m+1:
		a2 = B.abs[n-2] >> (_W - h)
	default:
		a2 = 0
	}

	// Since we are calculating with full words to avoid overflow,
	// we use 'even' to track the sign of the cosequences.
	// For even iterations: u0, v1 >= 0 && u1, v0 <= 0
	// For odd  iterations: u0, v1 <= 0 && u1, v0 >= 0
	// The first iteration starts with k=1 (odd).
	even = false
	// variables to track the cosequences
	u0, u1, u2 = 0, 1, 0
	v0, v1, v2 = 0, 0, 1

	// Calculate the quotient and cosequences using Collins' stopping condition.
	// Note that overflow of a Word is not possible when computing the remainder
	// sequence and cosequences since the cosequence size is bounded by the input size.
	// See section 4.2 of Jebelean for details.
	for a2 >= v2 && a1-a2 >= v1+v2 {
		q, r := a1/a2, a1%a2
		a1, a2 = a2, r
		u0, u1, u2 = u1, u2, u1+q*u2
		v0, v1, v2 = v1, v2, v1+q*v2
		even = !even
	}
	return
}

// lehmerUpdate updates the inputs A and B such that:
//		A = u0*A + v0*B
//		B = u1*A + v1*B
// where the signs of u0, u1, v0, v1 are given by even
// For even == true: u0, v1 >= 0 && u1, v0 <= 0
// For even == false: u0, v1 <= 0 && u1, v0 >= 0
// q, r, s, t are temporary variables to avoid allocations in the multiplication
func lehmerUpdate(A, B, q, r, s, t *Int, u0, u1, v0, v1 Word, even bool) {

	t.abs = t.abs.setWord(u0)
	s.abs = s.abs.setWord(v0)
	t.neg = !even
	s.neg = even

	t.Mul(A, t)
	s.Mul(B, s)

	r.abs = r.abs.setWord(u1)
	q.abs = q.abs.setWord(v1)
	r.neg = even
	q.neg = !even

	r.Mul(A, r)
	q.Mul(B, q)

	A.Add(t, s)
	B.Add(r, q)
}

// euclidUpdate performs a single step of the Euclidean GCD algorithm
// if extended is true, it also updates the cosequence Ua, Ub
func euclidUpdate(A, B, Ua, Ub, q, r, s, t *Int, extended bool) {
	q, r = q.QuoRem(A, B, r)

	*A, *B, *r = *B, *r, *A

	if extended {
		// Ua, Ub = Ub, Ua - q*Ub
		t.Set(Ub)
		s.Mul(Ub, q)
		Ub.Sub(Ua, s)
		Ua.Set(t)
	}
}

// lehmerGCD sets z to the greatest common divisor of a and b,
// which both must be != 0, and returns z.
// If x or y are not nil, their values are set such that z = a*x + b*y.
// See Knuth, The Art of Computer Programming, Vol. 2, Section 4.5.2, Algorithm L.
// This implementation uses the improved condition by Collins requiring only one
// quotient and avoiding the possibility of single Word overflow.
// See Jebelean, "Improving the multiprecision Euclidean algorithm",
// Design and Implementation of Symbolic Computation Systems, pp 45-58.
// The cosequences are updated according to Algorithm 10.45 from
// Cohen et al. "Handbook of Elliptic and Hyperelliptic Curve Cryptography" pp 192.
func (z *Int) lehmerGCD(x, y, a, b *Int) *Int {
	var A, B, Ua, Ub *Int

	A = new(Int).Abs(a)
	B = new(Int).Abs(b)

	extended := x != nil || y != nil

	if extended {
		// Ua (Ub) tracks how many times input a has been accumulated into A (B).
		Ua = new(Int).SetInt64(1)
		Ub = new(Int)
	}

	// temp variables for multiprecision update
	q := new(Int)
	r := new(Int)
	s := new(Int)
	t := new(Int)

	// ensure A >= B
	if A.abs.cmp(B.abs) < 0 {
		A, B = B, A
		Ub, Ua = Ua, Ub
	}

	// loop invariant A >= B
	for len(B.abs) > 1 {
		// Attempt to calculate in single-precision using leading words of A and B.
		u0, u1, v0, v1, even := lehmerSimulate(A, B)

		// multiprecision Step
		if v0 != 0 {
			// Simulate the effect of the single-precision steps using the cosequences.
			// A = u0*A + v0*B
			// B = u1*A + v1*B
			lehmerUpdate(A, B, q, r, s, t, u0, u1, v0, v1, even)

			if extended {
				// Ua = u0*Ua + v0*Ub
				// Ub = u1*Ua + v1*Ub
				lehmerUpdate(Ua, Ub, q, r, s, t, u0, u1, v0, v1, even)
			}

		} else {
			// Single-digit calculations failed to simulate any quotients.
			// Do a standard Euclidean step.
			euclidUpdate(A, B, Ua, Ub, q, r, s, t, extended)
		}
	}

	if len(B.abs) > 0 {
		// extended Euclidean algorithm base case if B is a single Word
		if len(A.abs) > 1 {
			// A is longer than a single Word, so one update is needed.
			euclidUpdate(A, B, Ua, Ub, q, r, s, t, extended)
		}
		if len(B.abs) > 0 {
			// A and B are both a single Word.
			aWord, bWord := A.abs[0], B.abs[0]
			if extended {
				var ua, ub, va, vb Word
				ua, ub = 1, 0
				va, vb = 0, 1
				even := true
				for bWord != 0 {
					q, r := aWord/bWord, aWord%bWord
					aWord, bWord = bWord, r
					ua, ub = ub, ua+q*ub
					va, vb = vb, va+q*vb
					even = !even
				}

				t.abs = t.abs.setWord(ua)
				s.abs = s.abs.setWord(va)
				t.neg = !even
				s.neg = even

				t.Mul(Ua, t)
				s.Mul(Ub, s)

				Ua.Add(t, s)
			} else {
				for bWord != 0 {
					aWord, bWord = bWord, aWord%bWord
				}
			}
			A.abs[0] = aWord
		}
	}
	negA := a.neg
	if y != nil {
		// avoid aliasing b needed in the division below
		if y == b {
			B.Set(b)
		} else {
			B = b
		}
		// y = (z - a*x)/b
		y.Mul(a, Ua) // y can safely alias a
		if negA {
			y.neg = !y.neg
		}
		y.Sub(A, y)
		y.Div(y, B)
	}

	if x != nil {
		*x = *Ua
		if negA {
			x.neg = !x.neg
		}
	}

	*z = *A

	return z
}

// Rand sets z to a pseudo-random number in [0, n) and returns z.
//
// As this uses the math/rand package, it must not be used for
// security-sensitive work. Use crypto/rand.Int instead.
func (z *Int) Rand(rnd *rand.Rand, n *Int) *Int {
	z.neg = false
	if n.neg || len(n.abs) == 0 {
		z.abs = nil
		return z
	}
	z.abs = z.abs.random(rnd, n.abs, n.abs.bitLen())
	return z
}

// ModInverse sets z to the multiplicative inverse of g in the ring ℤ/nℤ
// and returns z. If g and n are not relatively prime, g has no multiplicative
// inverse in the ring ℤ/nℤ.  In this case, z is unchanged and the return value
// is nil.
func (z *Int) ModInverse(g, n *Int) *Int {
	// GCD expects parameters a and b to be > 0.
	if n.neg {
		var n2 Int
		n = n2.Neg(n)
	}
	if g.neg {
		var g2 Int
		g = g2.Mod(g, n)
	}
	var d, x Int
	d.GCD(&x, nil, g, n)

	// if and only if d==1, g and n are relatively prime
	if d.Cmp(intOne) != 0 {
		return nil
	}

	// x and y are such that g*x + n*y = 1, therefore x is the inverse element,
	// but it may be negative, so convert to the range 0 <= z < |n|
	if x.neg {
		z.Add(&x, n)
	} else {
		z.Set(&x)
	}
	return z
}

// Jacobi returns the Jacobi symbol (x/y), either +1, -1, or 0.
// The y argument must be an odd integer.
func Jacobi(x, y *Int) int {
	if len(y.abs) == 0 || y.abs[0]&1 == 0 {
		panic(fmt.Sprintf("big: invalid 2nd argument to Int.Jacobi: need odd integer but got %s", y))
	}

	// We use the formulation described in chapter 2, section 2.4,
	// "The Yacas Book of Algorithms":
	// http://yacas.sourceforge.net/Algo.book.pdf

	var a, b, c Int
	a.Set(x)
	b.Set(y)
	j := 1

	if b.neg {
		if a.neg {
			j = -1
		}
		b.neg = false
	}

	for {
		if b.Cmp(intOne) == 0 {
			return j
		}
		if len(a.abs) == 0 {
			return 0
		}
		a.Mod(&a, &b)
		if len(a.abs) == 0 {
			return 0
		}
		// a > 0

		// handle factors of 2 in 'a'
		s := a.abs.trailingZeroBits()
		if s&1 != 0 {
			bmod8 := b.abs[0] & 7
			if bmod8 == 3 || bmod8 == 5 {
				j = -j
			}
		}
		c.Rsh(&a, s) // a = 2^s*c

		// swap numerator and denominator
		if b.abs[0]&3 == 3 && c.abs[0]&3 == 3 {
			j = -j
		}
		a.Set(&b)
		b.Set(&c)
	}
}

// modSqrt3Mod4 uses the identity
//      (a^((p+1)/4))^2  mod p
//   == u^(p+1)          mod p
//   == u^2              mod p
// to calculate the square root of any quadratic residue mod p quickly for 3
// mod 4 primes.
func (z *Int) modSqrt3Mod4Prime(x, p *Int) *Int {
	e := new(Int).Add(p, intOne) // e = p + 1
	e.Rsh(e, 2)                  // e = (p + 1) / 4
	z.Exp(x, e, p)               // z = x^e mod p
	return z
}

// modSqrt5Mod8 uses Atkin's observation that 2 is not a square mod p
//   alpha ==  (2*a)^((p-5)/8)    mod p
//   beta  ==  2*a*alpha^2        mod p  is a square root of -1
//   b     ==  a*alpha*(beta-1)   mod p  is a square root of a
// to calculate the square root of any quadratic residue mod p quickly for 5
// mod 8 primes.
func (z *Int) modSqrt5Mod8Prime(x, p *Int) *Int {
	// p == 5 mod 8 implies p = e*8 + 5
	// e is the quotient and 5 the remainder on division by 8
	e := new(Int).Rsh(p, 3)  // e = (p - 5) / 8
	tx := new(Int).Lsh(x, 1) // tx = 2*x
	alpha := new(Int).Exp(tx, e, p)
	beta := new(Int).Mul(alpha, alpha)
	beta.Mod(beta, p)
	beta.Mul(beta, tx)
	beta.Mod(beta, p)
	beta.Sub(beta, intOne)
	beta.Mul(beta, x)
	beta.Mod(beta, p)
	beta.Mul(beta, alpha)
	z.Mod(beta, p)
	return z
}

// modSqrtTonelliShanks uses the Tonelli-Shanks algorithm to find the square
// root of a quadratic residue modulo any prime.
func (z *Int) modSqrtTonelliShanks(x, p *Int) *Int {
	// Break p-1 into s*2^e such that s is odd.
	var s Int
	s.Sub(p, intOne)
	e := s.abs.trailingZeroBits()
	s.Rsh(&s, e)

	// find some non-square n
	var n Int
	n.SetInt64(2)
	for Jacobi(&n, p) != -1 {
		n.Add(&n, intOne)
	}

	// Core of the Tonelli-Shanks algorithm. Follows the description in
	// section 6 of "Square roots from 1; 24, 51, 10 to Dan Shanks" by Ezra
	// Brown:
	// https://www.maa.org/sites/default/files/pdf/upload_library/22/Polya/07468342.di020786.02p0470a.pdf
	var y, b, g, t Int
	y.Add(&s, intOne)
	y.Rsh(&y, 1)
	y.Exp(x, &y, p)  // y = x^((s+1)/2)
	b.Exp(x, &s, p)  // b = x^s
	g.Exp(&n, &s, p) // g = n^s
	r := e
	for {
		// find the least m such that ord_p(b) = 2^m
		var m uint
		t.Set(&b)
		for t.Cmp(intOne) != 0 {
			t.Mul(&t, &t).Mod(&t, p)
			m++
		}

		if m == 0 {
			return z.Set(&y)
		}

		t.SetInt64(0).SetBit(&t, int(r-m-1), 1).Exp(&g, &t, p)
		// t = g^(2^(r-m-1)) mod p
		g.Mul(&t, &t).Mod(&g, p) // g = g^(2^(r-m)) mod p
		y.Mul(&y, &t).Mod(&y, p)
		b.Mul(&b, &g).Mod(&b, p)
		r = m
	}
}

// ModSqrt sets z to a square root of x mod p if such a square root exists, and
// returns z. The modulus p must be an odd prime. If x is not a square mod p,
// ModSqrt leaves z unchanged and returns nil. This function panics if p is
// not an odd integer.
func (z *Int) ModSqrt(x, p *Int) *Int {
	switch Jacobi(x, p) {
	case -1:
		return nil // x is not a square mod p
	case 0:
		return z.SetInt64(0) // sqrt(0) mod p = 0
	case 1:
		break
	}
	if x.neg || x.Cmp(p) >= 0 { // ensure 0 <= x < p
		x = new(Int).Mod(x, p)
	}

	switch {
	case p.abs[0]%4 == 3:
		// Check whether p is 3 mod 4, and if so, use the faster algorithm.
		return z.modSqrt3Mod4Prime(x, p)
	case p.abs[0]%8 == 5:
		// Check whether p is 5 mod 8, use Atkin's algorithm.
		return z.modSqrt5Mod8Prime(x, p)
	default:
		// Otherwise, use Tonelli-Shanks.
		return z.modSqrtTonelliShanks(x, p)
	}
}

// Lsh sets z = x << n and returns z.
func (z *Int) Lsh(x *Int, n uint) *Int {
	z.abs = z.abs.shl(x.abs, n)
	z.neg = x.neg
	return z
}

// Rsh sets z = x >> n and returns z.
func (z *Int) Rsh(x *Int, n uint) *Int {
	if x.neg {
		// (-x) >> s == ^(x-1) >> s == ^((x-1) >> s) == -(((x-1) >> s) + 1)
		t := z.abs.sub(x.abs, natOne) // no underflow because |x| > 0
		t = t.shr(t, n)
		z.abs = t.add(t, natOne)
		z.neg = true // z cannot be zero if x is negative
		return z
	}

	z.abs = z.abs.shr(x.abs, n)
	z.neg = false
	return z
}

// Bit returns the value of the i'th bit of x. That is, it
// returns (x>>i)&1. The bit index i must be >= 0.
func (x *Int) Bit(i int) uint {
	if i == 0 {
		// optimization for common case: odd/even test of x
		if len(x.abs) > 0 {
			return uint(x.abs[0] & 1) // bit 0 is same for -x
		}
		return 0
	}
	if i < 0 {
		panic("negative bit index")
	}
	if x.neg {
		t := nat(nil).sub(x.abs, natOne)
		return t.bit(uint(i)) ^ 1
	}

	return x.abs.bit(uint(i))
}

// SetBit sets z to x, with x's i'th bit set to b (0 or 1).
// That is, if b is 1 SetBit sets z = x | (1 << i);
// if b is 0 SetBit sets z = x &^ (1 << i). If b is not 0 or 1,
// SetBit will panic.
func (z *Int) SetBit(x *Int, i int, b uint) *Int {
	if i < 0 {
		panic("negative bit index")
	}
	if x.neg {
		t := z.abs.sub(x.abs, natOne)
		t = t.setBit(t, uint(i), b^1)
		z.abs = t.add(t, natOne)
		z.neg = len(z.abs) > 0
		return z
	}
	z.abs = z.abs.setBit(x.abs, uint(i), b)
	z.neg = false
	return z
}

// And sets z = x & y and returns z.
func (z *Int) And(x, y *Int) *Int {
	if x.neg == y.neg {
		if x.neg {
			// (-x) & (-y) == ^(x-1) & ^(y-1) == ^((x-1) | (y-1)) == -(((x-1) | (y-1)) + 1)
			x1 := nat(nil).sub(x.abs, natOne)
			y1 := nat(nil).sub(y.abs, natOne)
			z.abs = z.abs.add(z.abs.or(x1, y1), natOne)
			z.neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x & y == x & y
		z.abs = z.abs.and(x.abs, y.abs)
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // & is symmetric
	}

	// x & (-y) == x & ^(y-1) == x &^ (y-1)
	y1 := nat(nil).sub(y.abs, natOne)
	z.abs = z.abs.andNot(x.abs, y1)
	z.neg = false
	return z
}

// AndNot sets z = x &^ y and returns z.
func (z *Int) AndNot(x, y *Int) *Int {
	if x.neg == y.neg {
		if x.neg {
			// (-x) &^ (-y) == ^(x-1) &^ ^(y-1) == ^(x-1) & (y-1) == (y-1) &^ (x-1)
			x1 := nat(nil).sub(x.abs, natOne)
			y1 := nat(nil).sub(y.abs, natOne)
			z.abs = z.abs.andNot(y1, x1)
			z.neg = false
			return z
		}

		// x &^ y == x &^ y
		z.abs = z.abs.andNot(x.abs, y.abs)
		z.neg = false
		return z
	}

	if x.neg {
		// (-x) &^ y == ^(x-1) &^ y == ^(x-1) & ^y == ^((x-1) | y) == -(((x-1) | y) + 1)
		x1 := nat(nil).sub(x.abs, natOne)
		z.abs = z.abs.add(z.abs.or(x1, y.abs), natOne)
		z.neg = true // z cannot be zero if x is negative and y is positive
		return z
	}

	// x &^ (-y) == x &^ ^(y-1) == x & (y-1)
	y1 := nat(nil).sub(y.abs, natOne)
	z.abs = z.abs.and(x.abs, y1)
	z.neg = false
	return z
}

// Or sets z = x | y and returns z.
func (z *Int) Or(x, y *Int) *Int {
	if x.neg == y.neg {
		if x.neg {
			// (-x) | (-y) == ^(x-1) | ^(y-1) == ^((x-1) & (y-1)) == -(((x-1) & (y-1)) + 1)
			x1 := nat(nil).sub(x.abs, natOne)
			y1 := nat(nil).sub(y.abs, natOne)
			z.abs = z.abs.add(z.abs.and(x1, y1), natOne)
			z.neg = true // z cannot be zero if x and y are negative
			return z
		}

		// x | y == x | y
		z.abs = z.abs.or(x.abs, y.abs)
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // | is symmetric
	}

	// x | (-y) == x | ^(y-1) == ^((y-1) &^ x) == -(^((y-1) &^ x) + 1)
	y1 := nat(nil).sub(y.abs, natOne)
	z.abs = z.abs.add(z.abs.andNot(y1, x.abs), natOne)
	z.neg = true // z cannot be zero if one of x or y is negative
	return z
}

// Xor sets z = x ^ y and returns z.
func (z *Int) Xor(x, y *Int) *Int {
	if x.neg == y.neg {
		if x.neg {
			// (-x) ^ (-y) == ^(x-1) ^ ^(y-1) == (x-1) ^ (y-1)
			x1 := nat(nil).sub(x.abs, natOne)
			y1 := nat(nil).sub(y.abs, natOne)
			z.abs = z.abs.xor(x1, y1)
			z.neg = false
			return z
		}

		// x ^ y == x ^ y
		z.abs = z.abs.xor(x.abs, y.abs)
		z.neg = false
		return z
	}

	// x.neg != y.neg
	if x.neg {
		x, y = y, x // ^ is symmetric
	}

	// x ^ (-y) == x ^ ^(y-1) == ^(x ^ (y-1)) == -((x ^ (y-1)) + 1)
	y1 := nat(nil).sub(y.abs, natOne)
	z.abs = z.abs.add(z.abs.xor(x.abs, y1), natOne)
	z.neg = true // z cannot be zero if only one of x or y is negative
	return z
}

// Not sets z = ^x and returns z.
func (z *Int) Not(x *Int) *Int {
	if x.neg {
		// ^(-x) == ^(^(x-1)) == x-1
		z.abs = z.abs.sub(x.abs, natOne)
		z.neg = false
		return z
	}

	// ^x == -x-1 == -(x+1)
	z.abs = z.abs.add(x.abs, natOne)
	z.neg = true // z cannot be zero if x is positive
	return z
}

// Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
// It panics if x is negative.
func (z *Int) Sqrt(x *Int) *Int {
	if x.neg {
		panic("square root of negative number")
	}
	z.neg = false
	z.abs = z.abs.sqrt(x.abs)
	return z
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements unsigned multi-precision integers (natural
// numbers). They are the building blocks for the implementation
// of signed integers, rationals, and floating-point numbers.
//
// Caution: This implementation relies on the function "alias"
//          which assumes that (nat) slice capacities are never
//          changed (no 3-operand slice expressions). If that
//          changes, alias needs to be updated for correctness.

// An unsigned integer x of the form
//
//   x = x[n-1]*_B^(n-1) + x[n-2]*_B^(n-2) + ... + x[1]*_B + x[0]
//
// with 0 <= x[i] < _B and 0 <= i < n is stored in a slice of length n,
// with the digits x[i] as the slice elements.
//
// A number is normalized if the slice contains no leading 0 digits.
// During arithmetic operations, denormalized values may occur but are
// always normalized before returning the final result. The normalized
// representation of 0 is the empty or nil slice (length = 0).
//
type nat []Word

var (
	natOne  = nat{1}
	natTwo  = nat{2}
	natFive = nat{5}
	natTen  = nat{10}
)

func (z nat) clear() {
	for i := range z {
		z[i] = 0
	}
}

func (z nat) norm() nat {
	i := len(z)
	for i > 0 && z[i-1] == 0 {
		i--
	}
	return z[0:i]
}

func (z nat) make(n int) nat {
	if n <= cap(z) {
		return z[:n] // reuse z
	}
	if n == 1 {
		// Most nats start small and stay that way; don't over-allocate.
		return make(nat, 1)
	}
	// Choosing a good value for e has significant performance impact
	// because it increases the chance that a value can be reused.
	const e = 4 // extra capacity
	return make(nat, n, n+e)
}

func (z nat) setWord(x Word) nat {
	if x == 0 {
		return z[:0]
	}
	z = z.make(1)
	z[0] = x
	return z
}

func (z nat) setUint64(x uint64) nat {
	// single-word value
	if w := Word(x); uint64(w) == x {
		return z.setWord(w)
	}
	// 2-word value
	z = z.make(2)
	z[1] = Word(x >> 32)
	z[0] = Word(x)
	return z
}

func (z nat) set(x nat) nat {
	z = z.make(len(x))
	copy(z, x)
	return z
}

func (z nat) add(x, y nat) nat {
	m := len(x)
	n := len(y)

	switch {
	case m < n:
		return z.add(y, x)
	case m == 0:
		// n == 0 because m >= n; result is 0
		return z[:0]
	case n == 0:
		// result is x
		return z.set(x)
	}
	// m > 0

	z = z.make(m + 1)
	c := addVV(z[0:n], x, y)
	if m > n {
		c = addVW(z[n:m], x[n:], c)
	}
	z[m] = c

	return z.norm()
}

func (z nat) sub(x, y nat) nat {
	m := len(x)
	n := len(y)

	switch {
	case m < n:
		panic("underflow")
	case m == 0:
		// n == 0 because m >= n; result is 0
		return z[:0]
	case n == 0:
		// result is x
		return z.set(x)
	}
	// m > 0

	z = z.make(m)
	c := subVV(z[0:n], x, y)
	if m > n {
		c = subVW(z[n:], x[n:], c)
	}
	if c != 0 {
		panic("underflow")
	}

	return z.norm()
}

func (x nat) cmp(y nat) (r int) {
	m := len(x)
	n := len(y)
	if m != n || m == 0 {
		switch {
		case m < n:
			r = -1
		case m > n:
			r = 1
		}
		return
	}

	i := m - 1
	for i > 0 && x[i] == y[i] {
		i--
	}

	switch {
	case x[i] < y[i]:
		r = -1
	case x[i] > y[i]:
		r = 1
	}
	return
}

func (z nat) mulAddWW(x nat, y, r Word) nat {
	m := len(x)
	if m == 0 || y == 0 {
		return z.setWord(r) // result is r
	}
	// m > 0

	z = z.make(m + 1)
	z[m] = mulAddVWW(z[0:m], x, y, r)

	return z.norm()
}

// basicMul multiplies x and y and leaves the result in z.
// The (non-normalized) result is placed in z[0 : len(x) + len(y)].
func basicMul(z, x, y nat) {
	z[0 : len(x)+len(y)].clear() // initialize z
	for i, d := range y {
		if d != 0 {
			z[len(x)+i] = addMulVVW(z[i:i+len(x)], x, d)
		}
	}
}

// montgomery computes z mod m = x*y*2**(-n*_W) mod m,
// assuming k = -1/m mod 2**_W.
// z is used for storing the result which is returned;
// z must not alias x, y or m.
// See Gueron, "Efficient Software Implementations of Modular Exponentiation".
// https://eprint.iacr.org/2011/239.pdf
// In the terminology of that paper, this is an "Almost Montgomery Multiplication":
// x and y are required to satisfy 0 <= z < 2**(n*_W) and then the result
// z is guaranteed to satisfy 0 <= z < 2**(n*_W), but it may not be < m.
func (z nat) montgomery(x, y, m nat, k Word, n int) nat {
	// This code assumes x, y, m are all the same length, n.
	// (required by addMulVVW and the for loop).
	// It also assumes that x, y are already reduced mod m,
	// or else the result will not be properly reduced.
	if len(x) != n || len(y) != n || len(m) != n {
		panic("math/big: mismatched montgomery number lengths")
	}
	z = z.make(n * 2)
	z.clear()
	var c Word
	for i := 0; i < n; i++ {
		d := y[i]
		c2 := addMulVVW(z[i:n+i], x, d)
		t := z[i] * k
		c3 := addMulVVW(z[i:n+i], m, t)
		cx := c + c2
		cy := cx + c3
		z[n+i] = cy
		if cx < c2 || cy < c3 {
			c = 1
		} else {
			c = 0
		}
	}
	if c != 0 {
		subVV(z[:n], z[n:], m)
	} else {
		copy(z[:n], z[n:])
	}
	return z[:n]
}

// Fast version of z[0:n+n>>1].add(z[0:n+n>>1], x[0:n]) w/o bounds checks.
// Factored out for readability - do not use outside karatsuba.
func karatsubaAdd(z, x nat, n int) {
	if c := addVV(z[0:n], z, x); c != 0 {
		addVW(z[n:n+n>>1], z[n:], c)
	}
}

// Like karatsubaAdd, but does subtract.
func karatsubaSub(z, x nat, n int) {
	if c := subVV(z[0:n], z, x); c != 0 {
		subVW(z[n:n+n>>1], z[n:], c)
	}
}

// Operands that are shorter than karatsubaThreshold are multiplied using
// "grade school" multiplication; for longer operands the Karatsuba algorithm
// is used.
var karatsubaThreshold = 40 // computed by calibrate_test.go

// karatsuba multiplies x and y and leaves the result in z.
// Both x and y must have the same length n and n must be a
// power of 2. The result vector z must have len(z) >= 6*n.
// The (non-normalized) result is placed in z[0 : 2*n].
func karatsuba(z, x, y nat) {
	n := len(y)

	// Switch to basic multiplication if numbers are odd or small.
	// (n is always even if karatsubaThreshold is even, but be
	// conservative)
	if n&1 != 0 || n < karatsubaThreshold || n < 2 {
		basicMul(z, x, y)
		return
	}
	// n&1 == 0 && n >= karatsubaThreshold && n >= 2

	// Karatsuba multiplication is based on the observation that
	// for two numbers x and y with:
	//
	//   x = x1*b + x0
	//   y = y1*b + y0
	//
	// the product x*y can be obtained with 3 products z2, z1, z0
	// instead of 4:
	//
	//   x*y = x1*y1*b*b + (x1*y0 + x0*y1)*b + x0*y0
	//       =    z2*b*b +              z1*b +    z0
	//
	// with:
	//
	//   xd = x1 - x0
	//   yd = y0 - y1
	//
	//   z1 =      xd*yd                    + z2 + z0
	//      = (x1-x0)*(y0 - y1)             + z2 + z0
	//      = x1*y0 - x1*y1 - x0*y0 + x0*y1 + z2 + z0
	//      = x1*y0 -    z2 -    z0 + x0*y1 + z2 + z0
	//      = x1*y0                 + x0*y1

	// split x, y into "digits"
	n2 := n >> 1              // n2 >= 1
	x1, x0 := x[n2:], x[0:n2] // x = x1*b + y0
	y1, y0 := y[n2:], y[0:n2] // y = y1*b + y0

	// z is used for the result and temporary storage:
	//
	//   6*n     5*n     4*n     3*n     2*n     1*n     0*n
	// z = [z2 copy|z0 copy| xd*yd | yd:xd | x1*y1 | x0*y0 ]
	//
	// For each recursive call of karatsuba, an unused slice of
	// z is passed in that has (at least) half the length of the
	// caller's z.

	// compute z0 and z2 with the result "in place" in z
	karatsuba(z, x0, y0)     // z0 = x0*y0
	karatsuba(z[n:], x1, y1) // z2 = x1*y1

	// compute xd (or the negative value if underflow occurs)
	s := 1 // sign of product xd*yd
	xd := z[2*n : 2*n+n2]
	if subVV(xd, x1, x0) != 0 { // x1-x0
		s = -s
		subVV(xd, x0, x1) // x0-x1
	}

	// compute yd (or the negative value if underflow occurs)
	yd := z[2*n+n2 : 3*n]
	if subVV(yd, y0, y1) != 0 { // y0-y1
		s = -s
		subVV(yd, y1, y0) // y1-y0
	}

	// p = (x1-x0)*(y0-y1) == x1*y0 - x1*y1 - x0*y0 + x0*y1 for s > 0
	// p = (x0-x1)*(y0-y1) == x0*y0 - x0*y1 - x1*y0 + x1*y1 for s < 0
	p := z[n*3:]
	karatsuba(p, xd, yd)

	// save original z2:z0
	// (ok to use upper half of z since we're done recursing)
	r := z[n*4:]
	copy(r, z[:n*2])

	// add up all partial products
	//
	//   2*n     n     0
	// z = [ z2  | z0  ]
	//   +    [ z0  ]
	//   +    [ z2  ]
	//   +    [  p  ]
	//
	karatsubaAdd(z[n2:], r, n)
	karatsubaAdd(z[n2:], r[n:], n)
	if s > 0 {
		karatsubaAdd(z[n2:], p, n)
	} else {
		karatsubaSub(z[n2:], p, n)
	}
}

// alias reports whether x and y share the same base array.
// Note: alias assumes that the capacity of underlying arrays
//       is never changed for nat values; i.e. that there are
//       no 3-operand slice expressions in this code (or worse,
//       reflect-based operations to the same effect).
func alias(x, y nat) bool {
	return cap(x) > 0 && cap(y) > 0 && &x[0:cap(x)][cap(x)-1] == &y[0:cap(y)][cap(y)-1]
}

// addAt implements z += x<<(_W*i); z must be long enough.
// (we don't use nat.add because we need z to stay the same
// slice, and we don't need to normalize z after each addition)
func addAt(z, x nat, i int) {
	if n := len(x); n > 0 {
		if c := addVV(z[i:i+n], z[i:], x); c != 0 {
			j := i + n
			if j < len(z) {
				addVW(z[j:], z[j:], c)
			}
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// karatsubaLen computes an approximation to the maximum k <= n such that
// k = p<<i for a number p <= threshold and an i >= 0. Thus, the
// result is the largest number that can be divided repeatedly by 2 before
// becoming about the value of threshold.
func karatsubaLen(n, threshold int) int {
	i := uint(0)
	for n > threshold {
		n >>= 1
		i++
	}
	return n << i
}

func (z nat) mul(x, y nat) nat {
	m := len(x)
	n := len(y)

	switch {
	case m < n:
		return z.mul(y, x)
	case m == 0 || n == 0:
		return z[:0]
	case n == 1:
		return z.mulAddWW(x, y[0], 0)
	}
	// m >= n > 1

	// determine if z can be reused
	if alias(z, x) || alias(z, y) {
		z = nil // z is an alias for x or y - cannot reuse
	}

	// use basic multiplication if the numbers are small
	if n < karatsubaThreshold {
		z = z.make(m + n)
		basicMul(z, x, y)
		return z.norm()
	}
	// m >= n && n >= karatsubaThreshold && n >= 2

	// determine Karatsuba length k such that
	//
	//   x = xh*b + x0  (0 <= x0 < b)
	//   y = yh*b + y0  (0 <= y0 < b)
	//   b = 1<<(_W*k)  ("base" of digits xi, yi)
	//
	k := karatsubaLen(n, karatsubaThreshold)
	// k <= n

	// multiply x0 and y0 via Karatsuba
	x0 := x[0:k]              // x0 is not normalized
	y0 := y[0:k]              // y0 is not normalized
	z = z.make(max(6*k, m+n)) // enough space for karatsuba of x0*y0 and full result of x*y
	karatsuba(z, x0, y0)
	z = z[0 : m+n]  // z has final length but may be incomplete
	z[2*k:].clear() // upper portion of z is garbage (and 2*k <= m+n since k <= n <= m)

	// If xh != 0 or yh != 0, add the missing terms to z. For
	//
	//   xh = xi*b^i + ... + x2*b^2 + x1*b (0 <= xi < b)
	//   yh =                         y1*b (0 <= y1 < b)
	//
	// the missing terms are
	//
	//   x0*y1*b and xi*y0*b^i, xi*y1*b^(i+1) for i > 0
	//
	// since all the yi for i > 1 are 0 by choice of k: If any of them
	// were > 0, then yh >= b^2 and thus y >= b^2. Then k' = k*2 would
	// be a larger valid threshold contradicting the assumption about k.
	//
	if k < n || m != n {
		tp := getNat(3 * k)
		t := *tp

		// add x0*y1*b
		x0 := x0.norm()
		y1 := y[k:]       // y1 is normalized because y is
		t = t.mul(x0, y1) // update t so we don't lose t's underlying array
		addAt(z, t, k)

		// add xi*y0<<i, xi*y1*b<<(i+k)
		y0 := y0.norm()
		for i := k; i < len(x); i += k {
			xi := x[i:]
			if len(xi) > k {
				xi = xi[:k]
			}
			xi = xi.norm()
			t = t.mul(xi, y0)
			addAt(z, t, i)
			t = t.mul(xi, y1)
			addAt(z, t, i+k)
		}

		putNat(tp)
	}

	return z.norm()
}

// basicSqr sets z = x*x and is asymptotically faster than basicMul
// by about a factor of 2, but slower for small arguments due to overhead.
// Requirements: len(x) > 0, len(z) == 2*len(x)
// The (non-normalized) result is placed in z.
func basicSqr(z, x nat) {
	n := len(x)
	tp := getNat(2 * n)
	t := *tp // temporary variable to hold the products
	t.clear()
	z[1], z[0] = mulWW(x[0], x[0]) // the initial square
	for i := 1; i < n; i++ {
		d := x[i]
		// z collects the squares x[i] * x[i]
		z[2*i+1], z[2*i] = mulWW(d, d)
		// t collects the products x[i] * x[j] where j < i
		t[2*i] = addMulVVW(t[i:2*i], x[0:i], d)
	}
	t[2*n-1] = shlVU(t[1:2*n-1], t[1:2*n-1], 1) // double the j < i products
	addVV(z, z, t)                              // combine the result
	putNat(tp)
}

// karatsubaSqr squares x and leaves the result in z.
// len(x) must be a power of 2 and len(z) >= 6*len(x).
// The (non-normalized) result is placed in z[0 : 2*len(x)].
//
// The algorithm and the layout of z are the same as for karatsuba.
func karatsubaSqr(z, x nat) {
	n := len(x)

	if n&1 != 0 || n < karatsubaSqrThreshold || n < 2 {
		basicSqr(z[:2*n], x)
		return
	}

	n2 := n >> 1
	x1, x0 := x[n2:], x[0:n2]

	karatsubaSqr(z, x0)
	karatsubaSqr(z[n:], x1)

	// s = sign(xd*yd) == -1 for xd != 0; s == 1 for xd == 0
	xd := z[2*n : 2*n+n2]
	if subVV(xd, x1, x0) != 0 {
		subVV(xd, x0, x1)
	}

	p := z[n*3:]
	karatsubaSqr(p, xd)

	r := z[n*4:]
	copy(r, z[:n*2])

	karatsubaAdd(z[n2:], r, n)
	karatsubaAdd(z[n2:], r[n:], n)
	karatsubaSub(z[n2:], p, n) // s == -1 for p != 0; s == 1 for p == 0
}

// Operands that are shorter than basicSqrThreshold are squared using
// "grade school" multiplication; for operands longer than karatsubaSqrThreshold
// we use the Karatsuba algorithm optimized for x == y.
var basicSqrThreshold = 20      // computed by calibrate_test.go
var karatsubaSqrThreshold = 260 // computed by calibrate_test.go

// z = x*x
func (z nat) sqr(x nat) nat {
	n := len(x)
	switch {
	case n == 0:
		return z[:0]
	case n == 1:
		d := x[0]
		z = z.make(2)
		z[1], z[0] = mulWW(d, d)
		return z.norm()
	}

	if alias(z, x) {
		z = nil // z is an alias for x - cannot reuse
	}

	if n < basicSqrThreshold {
		z = z.make(2 * n)
		basicMul(z, x, x)
		return z.norm()
	}
	if n < karatsubaSqrThreshold {
		z = z.make(2 * n)
		basicSqr(z, x)
		return z.norm()
	}

	// Use Karatsuba multiplication optimized for x == y.
	// The algorithm and layout of z are the same as for mul.

	// z = (x1*b + x0)^2 = x1^2*b^2 + 2*x1*x0*b + x0^2

	k := karatsubaLen(n, karatsubaSqrThreshold)

	x0 := x[0:k]
	z = z.make(max(6*k, 2*n))
	karatsubaSqr(z, x0) // z = x0^2
	z = z[0 : 2*n]
	z[2*k:].clear()

	if k < n {
		tp := getNat(2 * k)
		t := *tp
		x0 := x0.norm()
		x1 := x[k:]
		t = t.mul(x0, x1)
		addAt(z, t, k)
		addAt(z, t, k) // z = 2*x1*x0*b + x0^2
		t = t.sqr(x1)
		addAt(z, t, 2*k) // z = x1^2*b^2 + 2*x1*x0*b + x0^2
		putNat(tp)
	}

	return z.norm()
}

// mulRange computes the product of all the unsigned integers in the
// range [a, b] inclusively. If a > b (empty range), the result is 1.
func (z nat) mulRange(a, b uint64) nat {
	switch {
	case a == 0:
		// cut long ranges short (optimization)
		return z.setUint64(0)
	case a > b:
		return z.setUint64(1)
	case a == b:
		return z.setUint64(a)
	case a+1 == b:
		return z.mul(nat(nil).setUint64(a), nat(nil).setUint64(b))
	}
	m := (a + b) / 2
	return z.mul(nat(nil).mulRange(a, m), nat(nil).mulRange(m+1, b))
}

// q = (x-r)/y, with 0 <= r < y
func (z nat) divW(x nat, y Word) (q nat, r Word) {
	m := len(x)
	switch {
	case y == 0:
		panic("division by zero")
	case y == 1:
		q = z.set(x) // result is x
		return
	case m == 0:
		q = z[:0] // result is 0
		return
	}
	// m > 0
	z = z.make(m)
	r = divWVW(z, 0, x, y)
	q = z.norm()
	return
}

func (z nat) div(z2, u, v nat) (q, r nat) {
	if len(v) == 0 {
		panic("division by zero")
	}

	if u.cmp(v) < 0 {
		q = z[:0]
		r = z2.set(u)
		return
	}

	if len(v) == 1 {
		var r2 Word
		q, r2 = z.divW(u, v[0])
		r = z2.setWord(r2)
		return
	}

	q, r = z.divLarge(z2, u, v)
	return
}

// getNat returns a *nat of len n. The contents may not be zero.
// The pool holds *nat to avoid allocation when converting to interface{}.
func getNat(n int) *nat {
	var z *nat
	if v := natPool.Get(); v != nil {
		z = v.(*nat)
	}
	if z == nil {
		z = new(nat)
	}
	*z = z.make(n)
	return z
}

func putNat(x *nat) {
	natPool.Put(x)
}

var natPool sync.Pool

// q = (uIn-r)/vIn, with 0 <= r < vIn
// Uses z as storage for q, and u as storage for r if possible.
// See Knuth, Volume 2, section 4.3.1, Algorithm D.
// Preconditions:
//    len(vIn) >= 2
//    len(uIn) >= len(vIn)
//    u must not alias z
func (z nat) divLarge(u, uIn, vIn nat) (q, r nat) {
	n := len(vIn)
	m := len(uIn) - n

	// D1.
	shift := nlz(vIn[n-1])
	// do not modify vIn, it may be used by another goroutine simultaneously
	vp := getNat(n)
	v := *vp
	shlVU(v, vIn, shift)

	// u may safely alias uIn or vIn, the value of uIn is used to set u and vIn was already used
	u = u.make(len(uIn) + 1)
	u[len(uIn)] = shlVU(u[0:len(uIn)], uIn, shift)

	// z may safely alias uIn or vIn, both values were used already
	if alias(z, u) {
		z = nil // z is an alias for u - cannot reuse
	}
	q = z.make(m + 1)

	if n < divRecursiveThreshold {
		q.divBasic(u, v)
	} else {
		q.divRecursive(u, v)
	}
	putNat(vp)

	q = q.norm()
	shrVU(u, u, shift)
	r = u.norm()

	return q, r
}

// divBasic performs word-by-word division of u by v.
// The quotient is written in pre-allocated q.
// The remainder overwrites input u.
//
// Precondition:
// - q is large enough to hold the quotient u / v
//   which has a maximum length of len(u)-len(v)+1.
func (q nat) divBasic(u, v nat) {
	n := len(v)
	m := len(u) - n

	qhatvp := getNat(n + 1)
	qhatv := *qhatvp

	// D2.
	vn1 := v[n-1]
	for j := m; j >= 0; j-- {
		// D3.
		qhat := Word(_M)
		var ujn Word
		if j+n < len(u) {
			ujn = u[j+n]
		}
		if ujn != vn1 {
			var rhat Word
			qhat, rhat = divWW(ujn, u[j+n-1], vn1)

			// x1 | x2 = q̂v_{n-2}
			vn2 := v[n-2]
			x1, x2 := mulWW(qhat, vn2)
			// test if q̂v_{n-2} > br̂ + u_{j+n-2}
			ujn2 := u[j+n-2]
			for greaterThan(x1, x2, rhat, ujn2) {
				qhat--
				prevRhat := rhat
				rhat += vn1
				// v[n-1] >= 0, so this tests for overflow.
				if rhat < prevRhat {
					break
				}
				x1, x2 = mulWW(qhat, vn2)
			}
		}

		// D4.
		// Compute the remainder u - (q̂*v) << (_W*j).
		// The subtraction may overflow if q̂ estimate was off by one.
		qhatv[n] = mulAddVWW(qhatv[0:n], v, qhat, 0)
		qhl := len(qhatv)
		if j+qhl > len(u) && qhatv[n] == 0 {
			qhl--
		}
		c := subVV(u[j:j+qhl], u[j:], qhatv)
		if c != 0 {
			c := addVV(u[j:j+n], u[j:], v)
			// If n == qhl, the carry from subVV and the carry from addVV
			// cancel out and don't affect u[j+n].
			if n < qhl {
				u[j+n] += c
			}
			qhat--
		}

		if j == m && m == len(q) && qhat == 0 {
			continue
		}
		q[j] = qhat
	}

	putNat(qhatvp)
}

const divRecursiveThreshold = 100

// divRecursive performs word-by-word division of u by v.
// The quotient is written in pre-allocated z.
// The remainder overwrites input u.
//
// Precondition:
// - len(z) >= len(u)-len(v)
//
// See Burnikel, Ziegler, "Fast Recursive Division", Algorithm 1 and 2.
func (z nat) divRecursive(u, v nat) {
	// Recursion depth is less than 2 log2(len(v))
	// Allocate a slice of temporaries to be reused across recursion.
	recDepth := 2 * bits.Len(uint(len(v)))
	// large enough to perform Karatsuba on operands as large as v
	tmp := getNat(3 * len(v))
	temps := make([]*nat, recDepth)
	z.clear()
	z.divRecursiveStep(u, v, 0, tmp, temps)
	for _, n := range temps {
		if n != nil {
			putNat(n)
		}
	}
	putNat(tmp)
}

// divRecursiveStep computes the division of u by v.
// - z must be large enough to hold the quotient
// - the quotient will overwrite z
// - the remainder will overwrite u
func (z nat) divRecursiveStep(u, v nat, depth int, tmp *nat, temps []*nat) {
	u = u.norm()
	v = v.norm()

	if len(u) == 0 {
		z.clear()
		return
	}
	n := len(v)
	if n < divRecursiveThreshold {
		z.divBasic(u, v)
		return
	}
	m := len(u) - n
	if m < 0 {
		return
	}

	// Produce the quotient by blocks of B words.
	// Division by v (length n) is done using a length n/2 division
	// and a length n/2 multiplication for each block. The final
	// complexity is driven by multiplication complexity.
	B := n / 2

	// Allocate a nat for qhat below.
	if temps[depth] == nil {
		temps[depth] = getNat(n)
	} else {
		*temps[depth] = temps[depth].make(B + 1)
	}

	j := m
	for j > B {
		// Divide u[j-B:j+n] by vIn. Keep remainder in u
		// for next block.
		//
		// The following property will be used (Lemma 2):
		// if u = u1 << s + u0
		//    v = v1 << s + v0
		// then floor(u1/v1) >= floor(u/v)
		//
		// Moreover, the difference is at most 2 if len(v1) >= len(u/v)
		// We choose s = B-1 since len(v)-B >= B+1 >= len(u/v)
		s := (B - 1)
		// Except for the first step, the top bits are always
		// a division remainder, so the quotient length is <= n.
		uu := u[j-B:]

		qhat := *temps[depth]
		qhat.clear()
		qhat.divRecursiveStep(uu[s:B+n], v[s:], depth+1, tmp, temps)
		qhat = qhat.norm()
		// Adjust the quotient:
		//    u = u_h << s + u_l
		//    v = v_h << s + v_l
		//  u_h = q̂ v_h + rh
		//    u = q̂ (v - v_l) + rh << s + u_l
		// After the above step, u contains a remainder:
		//    u = rh << s + u_l
		// and we need to subtract q̂ v_l
		//
		// But it may be a bit too large, in which case q̂ needs to be smaller.
		qhatv := tmp.make(3 * n)
		qhatv.clear()
		qhatv = qhatv.mul(qhat, v[:s])
		for i := 0; i < 2; i++ {
			e := qhatv.cmp(uu.norm())
			if e <= 0 {
				break
			}
			subVW(qhat, qhat, 1)
			c := subVV(qhatv[:s], qhatv[:s], v[:s])
			if len(qhatv) > s {
				subVW(qhatv[s:], qhatv[s:], c)
			}
			addAt(uu[s:], v[s:], 0)
		}
		if qhatv.cmp(uu.norm()) > 0 {
			panic("impossible")
		}
		c := subVV(uu[:len(qhatv)], uu[:len(qhatv)], qhatv)
		if c > 0 {
			subVW(uu[len(qhatv):], uu[len(qhatv):], c)
		}
		addAt(z, qhat, j-B)
		j -= B
	}

	// Now u < (v<<B), compute lower bits in the same way.
	// Choose shift = B-1 again.
	s := B - 1
	qhat := *temps[depth]
	qhat.clear()
	qhat.divRecursiveStep(u[s:].norm(), v[s:], depth+1, tmp, temps)
	qhat = qhat.norm()
	qhatv := tmp.make(3 * n)
	qhatv.clear()
	qhatv = qhatv.mul(qhat, v[:s])
	// Set the correct remainder as before.
	for i := 0; i < 2; i++ {
		if e := qhatv.cmp(u.norm()); e > 0 {
			subVW(qhat, qhat, 1)
			c := subVV(qhatv[:s], qhatv[:s], v[:s])
			if len(qhatv) > s {
				subVW(qhatv[s:], qhatv[s:], c)
			}
			addAt(u[s:], v[s:], 0)
		}
	}
	if qhatv.cmp(u.norm()) > 0 {
		panic("impossible")
	}
	c := subVV(u[0:len(qhatv)], u[0:len(qhatv)], qhatv)
	if c > 0 {
		c = subVW(u[len(qhatv):], u[len(qhatv):], c)
	}
	if c > 0 {
		panic("impossible")
	}

	// Done!
	addAt(z, qhat.norm(), 0)
}

// Length of x in bits. x must be normalized.
func (x nat) bitLen() int {
	if i := len(x) - 1; i >= 0 {
		return i*_W + bits.Len(uint(x[i]))
	}
	return 0
}

// trailingZeroBits returns the number of consecutive least significant zero
// bits of x.
func (x nat) trailingZeroBits() uint {
	if len(x) == 0 {
		return 0
	}
	var i uint
	for x[i] == 0 {
		i++
	}
	// x[i] != 0
	return i*_W + uint(bits.TrailingZeros(uint(x[i])))
}

func same(x, y nat) bool {
	return len(x) == len(y) && len(x) > 0 && &x[0] == &y[0]
}

// z = x << s
func (z nat) shl(x nat, s uint) nat {
	if s == 0 {
		if same(z, x) {
			return z
		}
		if !alias(z, x) {
			return z.set(x)
		}
	}

	m := len(x)
	if m == 0 {
		return z[:0]
	}
	// m > 0

	n := m + int(s/_W)
	z = z.make(n + 1)
	z[n] = shlVU(z[n-m:n], x, s%_W)
	z[0 : n-m].clear()

	return z.norm()
}

// z = x >> s
func (z nat) shr(x nat, s uint) nat {
	if s == 0 {
		if same(z, x) {
			return z
		}
		if !alias(z, x) {
			return z.set(x)
		}
	}

	m := len(x)
	n := m - int(s/_W)
	if n <= 0 {
		return z[:0]
	}
	// n > 0

	z = z.make(n)
	shrVU(z, x[m-n:], s%_W)

	return z.norm()
}

func (z nat) setBit(x nat, i uint, b uint) nat {
	j := int(i / _W)
	m := Word(1) << (i % _W)
	n := len(x)
	switch b {
	case 0:
		z = z.make(n)
		copy(z, x)
		if j >= n {
			// no need to grow
			return z
		}
		z[j] &^= m
		return z.norm()
	case 1:
		if j >= n {
			z = z.make(j + 1)
			z[n:].clear()
		} else {
			z = z.make(n)
		}
		copy(z, x)
		z[j] |= m
		// no need to normalize
		return z
	}
	panic("set bit is not 0 or 1")
}

// bit returns the value of the i'th bit, with lsb == bit 0.
func (x nat) bit(i uint) uint {
	j := i / _W
	if j >= uint(len(x)) {
		return 0
	}
	// 0 <= j < len(x)
	return uint(x[j] >> (i % _W) & 1)
}

// sticky returns 1 if there's a 1 bit within the
// i least significant bits, otherwise it returns 0.
func (x nat) sticky(i uint) uint {
	j := i / _W
	if j >= uint(len(x)) {
		if len(x) == 0 {
			return 0
		}
		return 1
	}
	// 0 <= j < len(x)
	for _, x := range x[:j] {
		if x != 0 {
			return 1
		}
	}
	if x[j]<<(_W-i%_W) != 0 {
		return 1
	}
	return 0
}

func (z nat) and(x, y nat) nat {
	m := len(x)
	n := len(y)
	if m > n {
		m = n
	}
	// m <= n

	z = z.make(m)
	for i := 0; i < m; i++ {
		z[i] = x[i] & y[i]
	}

	return z.norm()
}

func (z nat) andNot(x, y nat) nat {
	m := len(x)
	n := len(y)
	if n > m {
		n = m
	}
	// m >= n

	z = z.make(m)
	for i := 0; i < n; i++ {
		z[i] = x[i] &^ y[i]
	}
	copy(z[n:m], x[n:m])

	return z.norm()
}

func (z nat) or(x, y nat) nat {
	m := len(x)
	n := len(y)
	s := x
	if m < n {
		n, m = m, n
		s = y
	}
	// m >= n

	z = z.make(m)
	for i := 0; i < n; i++ {
		z[i] = x[i] | y[i]
	}
	copy(z[n:m], s[n:m])

	return z.norm()
}

func (z nat) xor(x, y nat) nat {
	m := len(x)
	n := len(y)
	s := x
	if m < n {
		n, m = m, n
		s = y
	}
	// m >= n

	z = z.make(m)
	for i := 0; i < n; i++ {
		z[i] = x[i] ^ y[i]
	}
	copy(z[n:m], s[n:m])

	return z.norm()
}

// greaterThan reports whether (x1<<_W + x2) > (y1<<_W + y2)
func greaterThan(x1, x2, y1, y2 Word) bool {
	return x1 > y1 || x1 == y1 && x2 > y2
}

// modW returns x % d.
func (x nat) modW(d Word) (r Word) {
	// TODO(agl): we don't actually need to store the q value.
	var q nat
	q = q.make(len(x))
	return divWVW(q, 0, x, d)
}

// random creates a random integer in [0..limit), using the space in z if
// possible. n is the bit length of limit.
func (z nat) random(rand *rand.Rand, limit nat, n int) nat {
	if alias(z, limit) {
		z = nil // z is an alias for limit - cannot reuse
	}
	z = z.make(len(limit))

	bitLengthOfMSW := uint(n % _W)
	if bitLengthOfMSW == 0 {
		bitLengthOfMSW = _W
	}
	mask := Word((1 << bitLengthOfMSW) - 1)

	for {
		switch _W {
		case 32:
			for i := range z {
				z[i] = Word(rand.Uint32())
			}
		case 64:
			for i := range z {
				z[i] = Word(rand.Uint32()) | Word(rand.Uint32())<<32
			}
		default:
			panic("unknown word size")
		}
		z[len(limit)-1] &= mask
		if z.cmp(limit) < 0 {
			break
		}
	}

	return z.norm()
}

// If m != 0 (i.e., len(m) != 0), expNN sets z to x**y mod m;
// otherwise it sets z to x**y. The result is the value of z.
func (z nat) expNN(x, y, m nat) nat {
	if alias(z, x) || alias(z, y) {
		// We cannot allow in-place modification of x or y.
		z = nil
	}

	// x**y mod 1 == 0
	if len(m) == 1 && m[0] == 1 {
		return z.setWord(0)
	}
	// m == 0 || m > 1

	// x**0 == 1
	if len(y) == 0 {
		return z.setWord(1)
	}
	// y > 0

	// x**1 mod m == x mod m
	if len(y) == 1 && y[0] == 1 && len(m) != 0 {
		_, z = nat(nil).div(z, x, m)
		return z
	}
	// y > 1

	if len(m) != 0 {
		// We likely end up being as long as the modulus.
		z = z.make(len(m))
	}
	z = z.set(x)

	// If the base is non-trivial and the exponent is large, we use
	// 4-bit, windowed exponentiation. This involves precomputing 14 values
	// (x^2...x^15) but then reduces the number of multiply-reduces by a
	// third. Even for a 32-bit exponent, this reduces the number of
	// operations. Uses Montgomery method for odd moduli.
	if x.cmp(natOne) > 0 && len(y) > 1 && len(m) > 0 {
		if m[0]&1 == 1 {
			return z.expNNMontgomery(x, y, m)
		}
		return z.expNNWindowed(x, y, m)
	}

	v := y[len(y)-1] // v > 0 because y is normalized and y > 0
	shift := nlz(v) + 1
	v <<= shift
	var q nat

	const mask = 1 << (_W - 1)

	// We walk through the bits of the exponent one by one. Each time we
	// see a bit, we square, thus doubling the power. If the bit is a one,
	// we also multiply by x, thus adding one to the power.

	w := _W - int(shift)
	// zz and r are used to avoid allocating in mul and div as
	// otherwise the arguments would alias.
	var zz, r nat
	for j := 0; j < w; j++ {
		zz = zz.sqr(z)
		zz, z = z, zz

		if v&mask != 0 {
			zz = zz.mul(z, x)
			zz, z = z, zz
		}

		if len(m) != 0 {
			zz, r = zz.div(r, z, m)
			zz, r, q, z = q, z, zz, r
		}

		v <<= 1
	}

	for i := len(y) - 2; i >= 0; i-- {
		v = y[i]

		for j := 0; j < _W; j++ {
			zz = zz.sqr(z)
			zz, z = z, zz

			if v&mask != 0 {
				zz = zz.mul(z, x)
				zz, z = z, zz
			}

			if len(m) != 0 {
				zz, r = zz.div(r, z, m)
				zz, r, q, z = q, z, zz, r
			}

			v <<= 1
		}
	}

	return z.norm()
}

// expNNWindowed calculates x**y mod m using a fixed, 4-bit window.
func (z nat) expNNWindowed(x, y, m nat) nat {
	// zz and r are used to avoid allocating in mul and div as otherwise
	// the arguments would alias.
	var zz, r nat

	const n = 4
	// powers[i] contains x^i.
	var powers [1 << n]nat
	powers[0] = natOne
	powers[1] = x
	for i := 2; i < 1<<n; i += 2 {
		p2, p, p1 := &powers[i/2], &powers[i], &powers[i+1]
		*p = p.sqr(*p2)
		zz, r = zz.div(r, *p, m)
		*p, r = r, *p
		*p1 = p1.mul(*p, x)
		zz, r = zz.div(r, *p1, m)
		*p1, r = r, *p1
	}

	z = z.setWord(1)

	for i := len(y) - 1; i >= 0; i-- {
		yi := y[i]
		for j := 0; j < _W; j += n {
			if i != len(y)-1 || j != 0 {
				// Unrolled loop for significant performance
				// gain. Use go test -bench=".*" in crypto/rsa
				// to check performance before making changes.
				zz = zz.sqr(z)
				zz, z = z, zz
				zz, r = zz.div(r, z, m)
				z, r = r, z

				zz = zz.sqr(z)
				zz, z = z, zz
				zz, r = zz.div(r, z, m)
				z, r = r, z

				zz = zz.sqr(z)
				zz, z = z, zz
				zz, r = zz.div(r, z, m)
				z, r = r, z

				zz = zz.sqr(z)
				zz, z = z, zz
				zz, r = zz.div(r, z, m)
				z, r = r, z
			}

			zz = zz.mul(z, powers[yi>>(_W-n)])
			zz, z = z, zz
			zz, r = zz.div(r, z, m)
			z, r = r, z

			yi <<= n
		}
	}

	return z.norm()
}

// expNNMontgomery calculates x**y mod m using a fixed, 4-bit window.
// Uses Montgomery representation.
func (z nat) expNNMontgomery(x, y, m nat) nat {
	numWords := len(m)

	// We want the lengths of x and m to be equal.
	// It is OK if x >= m as long as len(x) == len(m).
	if len(x) > numWords {
		_, x = nat(nil).div(nil, x, m)
		// Note: now len(x) <= numWords, not guaranteed ==.
	}
	if len(x) < numWords {
		rr := make(nat, numWords)
		copy(rr, x)
		x = rr
	}

	// Ideally the precomputations would be performed outside, and reused
	// k0 = -m**-1 mod 2**_W. Algorithm from: Dumas, J.G. "On Newton–Raphson
	// Iteration for Multiplicative Inverses Modulo Prime Powers".
	k0 := 2 - m[0]
	t := m[0] - 1
	for i := 1; i < _W; i <<= 1 {
		t *= t
		k0 *= (t + 1)
	}
	k0 = -k0

	// RR = 2**(2*_W*len(m)) mod m
	RR := nat(nil).setWord(1)
	zz := nat(nil).shl(RR, uint(2*numWords*_W))
	_, RR = nat(nil).div(RR, zz, m)
	if len(RR) < numWords {
		zz = zz.make(numWords)
		copy(zz, RR)
		RR = zz
	}
	// one = 1, with equal length to that of m
	one := make(nat, numWords)
	one[0] = 1

	const n = 4
	// powers[i] contains x^i
	var powers [1 << n]nat
	powers[0] = powers[0].montgomery(one, RR, m, k0, numWords)
	powers[1] = powers[1].montgomery(x, RR, m, k0, numWords)
	for i := 2; i < 1<<n; i++ {
		powers[i] = powers[i].montgomery(powers[i-1], powers[1], m, k0, numWords)
	}

	// initialize z = 1 (Montgomery 1)
	z = z.make(numWords)
	copy(z, powers[0])

	zz = zz.make(numWords)

	// same windowed exponent, but with Montgomery multiplications
	for i := len(y) - 1; i >= 0; i-- {
		yi := y[i]
		for j := 0; j < _W; j += n {
			if i != len(y)-1 || j != 0 {
				zz = zz.montgomery(z, z, m, k0, numWords)
				z = z.montgomery(zz, zz, m, k0, numWords)
				zz = zz.montgomery(z, z, m, k0, numWords)
				z = z.montgomery(zz, zz, m, k0, numWords)
			}
			zz = zz.montgomery(z, powers[yi>>(_W-n)], m, k0, numWords)
			z, zz = zz, z
			yi <<= n
		}
	}
	// convert to regular number
	zz = zz.montgomery(z, one, m, k0, numWords)

	// One last reduction, just in case.
	// See golang.org/issue/13907.
	if zz.cmp(m) >= 0 {
		// Common case is m has high bit set; in that case,
		// since zz is the same length as m, there can be just
		// one multiple of m to remove. Just subtract.
		// We think that the subtract should be sufficient in general,
		// so do that unconditionally, but double-check,
		// in case our beliefs are wrong.
		// The div is not expected to be reached.
		zz = zz.sub(zz, m)
		if zz.cmp(m) >= 0 {
			_, zz = nat(nil).div(nil, zz, m)
		}
	}

	return zz.norm()
}

// bytes writes the value of z into buf using big-endian encoding.
// The value of z is encoded in the slice buf[i:]. If the value of z
// cannot be represented in buf, bytes panics. The number i of unused
// bytes at the beginning of buf is returned as result.
func (z nat) bytes(buf []byte) (i int) {
	i = len(buf)
	for _, d := range z {
		for j := 0; j < _S; j++ {
			i--
			if i >= 0 {
				buf[i] = byte(d)
			} else if byte(d) != 0 {
				panic("math/big: buffer too small to fit value")
			}
			d >>= 8
		}
	}

	if i < 0 {
		i = 0
	}
	for i < len(buf) && buf[i] == 0 {
		i++
	}

	return
}

// bigEndianWord returns the contents of buf interpreted as a big-endian encoded Word value.
func bigEndianWord(buf []byte) Word {
	if _W == 64 {
		return Word(binary.BigEndian.Uint64(buf))
	}
	return Word(binary.BigEndian.Uint32(buf))
}

// setBytes interprets buf as the bytes of a big-endian unsigned
// integer, sets z to that value, and returns z.
func (z nat) setBytes(buf []byte) nat {
	z = z.make((len(buf) + _S - 1) / _S)

	i := len(buf)
	for k := 0; i >= _S; k++ {
		z[k] = bigEndianWord(buf[i-_S : i])
		i -= _S
	}
	if i > 0 {
		var d Word
		for s := uint(0); i > 0; s += 8 {
			d |= Word(buf[i-1]) << s
			i--
		}
		z[len(z)-1] = d
	}

	return z.norm()
}

// sqrt sets z = ⌊√x⌋
func (z nat) sqrt(x nat) nat {
	if x.cmp(natOne) <= 0 {
		return z.set(x)
	}
	if alias(z, x) {
		z = nil
	}

	// Start with value known to be too large and repeat "z = ⌊(z + ⌊x/z⌋)/2⌋" until it stops getting smaller.
	// See Brent and Zimmermann, Modern Computer Arithmetic, Algorithm 1.13 (SqrtInt).
	// https://members.loria.fr/PZimmermann/mca/pub226.html
	// If x is one less than a perfect square, the sequence oscillates between the correct z and z+1;
	// otherwise it converges to the correct z and stays there.
	var z1, z2 nat
	z1 = z
	z1 = z1.setUint64(1)
	z1 = z1.shl(z1, uint(x.bitLen()+1)/2) // must be ≥ √x
	for n := 0; ; n++ {
		z2, _ = z2.div(nil, x, z1)
		z2 = z2.add(z2, z1)
		z2 = z2.shr(z2, 1)
		if z2.cmp(z1) >= 0 {
			// z1 is answer.
			// Figure out whether z1 or z2 is currently aliased to z by looking at loop count.
			if n&1 == 0 {
				return z1
			}
			return z.set(z1)
		}
		z1, z2 = z2, z1
	}
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides Go implementations of elementary multi-precision
// arithmetic operations on word vectors. These have the suffix _g.
// These are needed for platforms without assembly implementations of these routines.
// This file also contains elementary operations that can be implemented
// sufficiently efficiently in Go.

// A Word represents a single digit of a multi-precision unsigned integer.
type Word uint

const (
	_S = _W / 8 // word size in bytes

	_W = bits.UintSize // word size in bits
	_B = 1 << _W       // digit base
	_M = _B - 1        // digit mask
)

// Many of the loops in this file are of the form
//   for i := 0; i < len(z) && i < len(x) && i < len(y); i++
// i < len(z) is the real condition.
// However, checking i < len(x) && i < len(y) as well is faster than
// having the compiler do a bounds check in the body of the loop;
// remarkably it is even faster than hoisting the bounds check
// out of the loop, by doing something like
//   _, _ = x[len(z)-1], y[len(z)-1]
// There are other ways to hoist the bounds check out of the loop,
// but the compiler's BCE isn't powerful enough for them (yet?).
// See the discussion in CL 164966.

// ----------------------------------------------------------------------------
// Elementary operations on words
//
// These operations are used by the vector operations below.

// z1<<_W + z0 = x*y
func mulWW_g(x, y Word) (z1, z0 Word) {
	hi, lo := bits.Mul(uint(x), uint(y))
	return Word(hi), Word(lo)
}

// z1<<_W + z0 = x*y + c
func mulAddWWW_g(x, y, c Word) (z1, z0 Word) {
	hi, lo := bits.Mul(uint(x), uint(y))
	var cc uint
	lo, cc = bits.Add(lo, uint(c), 0)
	return Word(hi + cc), Word(lo)
}

// nlz returns the number of leading zeros in x.
// Wraps bits.LeadingZeros call for convenience.
func nlz(x Word) uint {
	return uint(bits.LeadingZeros(uint(x)))
}

// q = (u1<<_W + u0 - r)/v
func divWW_g(u1, u0, v Word) (q, r Word) {
	qq, rr := bits.Div(uint(u1), uint(u0), uint(v))
	return Word(qq), Word(rr)
}

// The resulting carry c is either 0 or 1.
func addVV_g(z, x, y []Word) (c Word) {
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x) && i < len(y); i++ {
		zi, cc := bits.Add(uint(x[i]), uint(y[i]), uint(c))
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

// The resulting carry c is either 0 or 1.
func subVV_g(z, x, y []Word) (c Word) {
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x) && i < len(y); i++ {
		zi, cc := bits.Sub(uint(x[i]), uint(y[i]), uint(c))
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

// The resulting carry c is either 0 or 1.
func addVW_g(z, x []Word, y Word) (c Word) {
	c = y
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		zi, cc := bits.Add(uint(x[i]), uint(c), 0)
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

// addVWlarge is addVW, but intended for large z.
// The only difference is that we check on every iteration
// whether we are done with carries,
// and if so, switch to a much faster copy instead.
// This is only a good idea for large z,
// because the overhead of the check and the function call
// outweigh the benefits when z is small.
func addVWlarge(z, x []Word, y Word) (c Word) {
	c = y
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		if c == 0 {
			copy(z[i:], x[i:])
			return
		}
		zi, cc := bits.Add(uint(x[i]), uint(c), 0)
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

func subVW_g(z, x []Word, y Word) (c Word) {
	c = y
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		zi, cc := bits.Sub(uint(x[i]), uint(c), 0)
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

// subVWlarge is to subVW as addVWlarge is to addVW.
func subVWlarge(z, x []Word, y Word) (c Word) {
	c = y
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		if c == 0 {
			copy(z[i:], x[i:])
			return
		}
		zi, cc := bits.Sub(uint(x[i]), uint(c), 0)
		z[i] = Word(zi)
		c = Word(cc)
	}
	return
}

func shlVU_g(z, x []Word, s uint) (c Word) {
	if s == 0 {
		copy(z, x)
		return
	}
	if len(z) == 0 {
		return
	}
	s &= _W - 1 // hint to the compiler that shifts by s don't need guard code
	ŝ := _W - s
	ŝ &= _W - 1 // ditto
	c = x[len(z)-1] >> ŝ
	for i := len(z) - 1; i > 0; i-- {
		z[i] = x[i]<<s | x[i-1]>>ŝ
	}
	z[0] = x[0] << s
	return
}

func shrVU_g(z, x []Word, s uint) (c Word) {
	if s == 0 {
		copy(z, x)
		return
	}
	if len(z) == 0 {
		return
	}
	s &= _W - 1 // hint to the compiler that shifts by s don't need guard code
	ŝ := _W - s
	ŝ &= _W - 1 // ditto
	c = x[0] << ŝ
	for i := 0; i < len(z)-1; i++ {
		z[i] = x[i]>>s | x[i+1]<<ŝ
	}
	z[len(z)-1] = x[len(z)-1] >> s
	return
}

func mulAddVWW_g(z, x []Word, y, r Word) (c Word) {
	c = r
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		c, z[i] = mulAddWWW_g(x[i], y, c)
	}
	return
}

func addMulVVW_g(z, x []Word, y Word) (c Word) {
	// The comment near the top of this file discusses this for loop condition.
	for i := 0; i < len(z) && i < len(x); i++ {
		z1, z0 := mulAddWWW_g(x[i], y, z[i])
		lo, cc := bits.Add(uint(z0), uint(c), 0)
		c, z[i] = Word(cc), Word(lo)
		c += z1
	}
	return
}

func divWVW_g(z []Word, xn Word, x []Word, y Word) (r Word) {
	r = xn
	for i := len(z) - 1; i >= 0; i-- {
		z[i], r = divWW_g(r, x[i], y)
	}
	return
}
