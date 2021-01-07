package main

import (
	"github.com/skeptycal/util/datatools/math/fastinvsqrt"
	"github.com/skeptycal/util/datatools/math/points"
)

func y1(x float64) float64 { return 5*x + 3 }
func y2(x float64) float64 { return 0.3*x - 3 }

func main() {
	var b fastinvsqrt.Bits = 0x0FFF0FFF
	b.PrintMethods(nil)

	ps := points.NewPointSet("y = 5x + 3", "", "")
	ps.MakeDataSeries(3.14, 97.3, 1.5, y1)
	ps.PrintDataTable(nil, 5)

	ps = points.NewPointSet("y = 0.3*x - 3", "2nd equation", "blue")
	ps.MakeDataSeries(0.1, 32.5, 0.24, y2)
	ps.PrintDataTable(nil, 5)
}
