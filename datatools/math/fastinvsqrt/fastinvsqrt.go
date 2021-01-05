// package fastinvsqrt
/*
note - start here: https://commandcenter.blogspot.com/2012/04/byte-order-fallacy.html

IEEE 754

uses IEEE 754 standard floating point number with
clever optimizations; uses only normalised numbers

    0   0000000 000000000000000000000000
    |            |               |--> 23 bit mantissa
    |            |-------------> 8 bit exponent
    |----------------------> 1 sign bit (0 means positive)

bit representation is
    2^23 * E + M (shift E by 23 bits and add M)
decimal number is
    (1 + M/2^23) * 2^(E-127)

Mantissa

23 bit mantissa - in binary the mantissa is unique; the only non-zero number before the decimal point is 1 ... so it is always 1 ... this means that the 1 is assumed and does not need to be represented.

Instead of 1 and 22 decimal (or binary places?)
    0.00000000000000000000000
we get the full 23 decimal places (or ... whatever)
    .000000000000000000000000

Range of numbers is 0 to 2^23-1

Exponent

The exponent is shifted by -127 to allow negative exponents

Instead of
    0 .. 255
we get a range of
    -127 to 128

Sign Bit

The sign bit is ignored in this algorithm; Real square roots
are only for positive numbers ...

References

based on code from quake3 algorithm

Reference: https://www.youtube.com/watch?v=p8u_k2LIZyo

*/
package fastinvsqrt

import (
	"fmt"
	"math"
	"math/big"
	"unsafe"

	log "github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/diff/fd"
)

type Any interface{}

//  i = (data[3]<<0) | (data[2]<<8) | (data[1]<<16) | (data[0]<<24);
// Ref: https://commandcenter.blogspot.com/2012/04/byte-order-fallacy.html
type Bits [4]byte

func (b Bits) String() string {
	return fmt.Sprintf("%v", Any(b))
}

// Sign returns the sign bit of the Bits number.
// 0 or 1 (-1 represents an error)
func (b Bits) Sign() uint32 {
	val, ok := Any(b).(uint32)
	if ok {
		return val & signBitMask
	}
	log.Errorf("cannot convert %b to int32", b)
	return 0
}

func (b *Bits) Decode() float32 { return DecodeBits(b) }

// Decode returns an IEEE 754 floating point number.
func DecodeBits(b *Bits) float32 { return *(*float32)(unsafe.Pointer(&b)) }

// Encode returns an new bitmapped IEEE 754 float32
func EncodeBits(f float32) uint32 { return *(*uint32)(unsafe.Pointer(&f)) }

func ByteToInt(n uint64) float64 {
	return math.Float64frombits(n)
}

// invSqrtBasic is a simple and slow implementation
func invSqrtBasic(x float64) float64 {
	return 1 / math.Sqrt(x)
}

// func sqrtEstimate(n float64, t tolerance) float64 {
// 	maxError := n * t // max absolute error
// 	d0 := n * 0.5
// 	fmt.Printf("n: %f", n)
// 	fmt.Printf("p: %f", precisionConstant)
// 	fmt.Printf("maxError: %f", maxError)
// 	fmt.Printf("d0: %f", d0)
// 	for p < maxError {
// 		d = d + dx
// 	}
// }

// slope is a slow approximation of the slope of f between 2 points
// it contains float64 values, function calls, and division operations,
// all of which are enormously expensive
// but ... it is a quick and dirty approximation that is simple to
// implement and understand
func slope(x1, x2 float64, f func(float64) float64) float64 {
	dy := f(x2) - f(x1)
	dx := x2 - x1
	return dy / dx
}

func AddAny(things ...interface{}) interface{} {
	var x, y big.Float
	// nan := math.NaN()

	for _, v := range things {
		if v, ok := v.(float64); ok && math.IsNaN(v) {
			return math.NaN()
		}
		if _, _, err := x.Parse(fmt.Sprint(v), 10); err != nil {
			return nil
		}
		y.Add(&y, &x)
	}
	return y.String()
}

// findRoot returns one root of f
// goal: function that has a root at the point where the inverse square root is located ...
// the Quake3 algorithm was interesting to me so I am experimenting with
// things along the same line
// func findRoot(x0 float64, f func(float64) float64) float64 {
// 	var x, y, y0, dy, lastdy, d2y float64
// 	var tryThis bool = true

// 	// initial y0
// 	y0 = f(x0)

// 	// artibrary x  ... (x+ 1) * 2 (todo - is this the best x1 ?)
// 	x = (x0 + 1.0) * 2
// 	y = f(x)

// 	dx = x - x1
// 	dy = y - y1
// 	d2y = dy - lastdy
// 	lastdy = dy

// 	absError := x * precisionConstant

// for dy < absError { // find dy close to zero

// 	y0 = y

// 	switch (d2y := dy - lastdy); {
// 	case x > 0.0: // to hot
// 	case x < 1: // too cold
// 	default: // just right
// 	}
// }
// }

func Fib100Digit() (big.Int, bool) {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Initialize limit as 10^99, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	// fmt.Println(a) // 100-digit Fibonacci number

	// Test a for primality.
	// (ProbablyPrimes' argument sets the number of Miller-Rabin
	// rounds to be performed. 20 is a good value.)
	// fmt.Println(a.ProbablyPrime(20))

	return *a, a.ProbablyPrime(20)
}

func showFD() {
	// https://godoc.org/gonum.org/v1/gonum/diff/fd
	f := func(x float64) float64 {
		return math.Sin(x)
	}
	// Compute the first derivative of f at 0 using the default settings.
	fmt.Println("f'(0) ≈", fd.Derivative(f, 0, nil))
	// Compute the first derivative of f at 0 using the forward approximation
	// with a custom step size.
	df := fd.Derivative(f, 0, &fd.Settings{
		Formula: fd.Forward,
		Step:    1e-3,
	})
	fmt.Println("f'(0) ≈", df)

	f = func(x float64) float64 {
		return math.Pow(math.Cos(x), 3)
	}
	// Compute the second derivative of f at 0 using
	// the centered approximation, concurrent evaluation,
	// and a known function value at x.
	df = fd.Derivative(f, 0, &fd.Settings{
		Formula:     fd.Central2nd,
		Concurrent:  true,
		OriginKnown: true,
		OriginValue: f(0),
	})
	fmt.Println("f''(0) ≈", df)
}
