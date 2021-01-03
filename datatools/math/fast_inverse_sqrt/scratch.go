package main

import (
	"fmt"
	"math"
	"math/big"
)

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
