package main

import (
	"fmt"
	"math"
	"math/big"

	"gonum.org/v1/gonum/diff/fd"
)

func main() {
	// sqrtEstimate(5, onePercent)
	// fmt.Printf("", findRoot(1, invSqrtBasic))
	fmt.Println(Fib100Digit())
	showFD()
}

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
