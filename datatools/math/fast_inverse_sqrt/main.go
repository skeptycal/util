package main

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"gonum.org/v1/gonum/diff/fd"
)

const (
	// arbitrary 1 ppt tolerance // TODO - move to config
	precisionConstant = onePPT
)

func Fib100Digit() {
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
	fmt.Println(a) // 100-digit Fibonacci number

	// Test a for primality.
	// (ProbablyPrimes' argument sets the number of Miller-Rabin
	// rounds to be performed. 20 is a good value.)
	fmt.Println(a.ProbablyPrime(20))

}

func invSqrtBasic(x float64) float64 {
	return 1 / math.Sqrt(x)
}

func sqrtEstimate(n float64, t tolerance) float64 {
	maxError := n * t // max absolute error
	d0 := n * 0.5
	fmt.Printf("n: %f", n)
	fmt.Printf("p: %f", precisionConstant)
	fmt.Printf("maxError: %f", maxError)
	fmt.Printf("d0: %f", d0)
	for p < maxError {
		d = d + dx
	}
}

func main() {
	sqrtEstimate(5, onePercent)
	fmt.Printf("", findRoot(1, invSqrtBasic))
	Fib100Digit()
}

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

func addAny(things ...interface{}) interface{} {
	var x, y big.Float

	for _, v := range things {
		if _, _, err := x.Parse(fmt.Sprint(v), 10); err != nil {
			return nil
		}
		y.Add(&y, &x)
	}
	return y.String()
}

func NewtonIter(x float64, f func(float64) float64) float64 {
	var xn, y, dy float64

	y = f(x)
	dy = y - f()

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

func printDay() {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}
}

// tolerance constants describe allowable
// tolerances for estimates
type tolerance = float64

// func (t *tolerance) String() string {
// 	return fmt.Sprintf("%.3F", t)
// }

const (
	fivePercent = 5e-2
	twoPercent  = 2e-2
	onePercent  = 1e-2
	onePPT      = 1e-3
	onePPM      = 1e-6
	onePPB      = 1e-9
)
