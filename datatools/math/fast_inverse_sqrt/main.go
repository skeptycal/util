package main

import (
	"fmt"
	"math"
)

const (
	precisionConstant = 0.001 // 1 ppt tolerance
)

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

func sqrtBasic(n float64) float64 {
	return 1 / math.Sqrt(n)
}

func sqrtEstimate(n float64, t tolerance) float64 {
	maxError := n * t // max error
	d0 := dx << 1
	fmt.Printf("n: %f", n)
	fmt.Printf("p: %f", precisionConstant)
	fmt.Printf("maxError: %f", maxError)
	fmt.Printf("d0: %f", d0)
	for d < maxError {
		d = d << 1
	}
}

func main() {

}
