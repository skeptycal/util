// Package points provides simple utilities for dealing with points
// in a data series.
package points

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Point represents an x,y data point with a data value.
type Point struct {
	X    float64
	Y    float64
	Data []byte
}

// PointSet represents a sequence of connected data points.
type PointSet struct {
	Name   string // name of this data set
	Color  string // display color for this set
	Xlabel string
	Ylabel string
	Unit   string   // unit of measurement for display purposes
	Points []*Point // points in the data set
}

func (p *PointSet) addPoint(np *Point) {
	p.Points = append(p.Points, np)
}

func (p *PointSet) Add(x, y float64, data []byte) {
	p.addPoint(&Point{x, y, data})
}

// Reset releases the underlying array data to the garbage
// collector (by setting it to nil) but leaves other values intact.
//
// The official Go wiki recommends using nil slices over empty slices.
//
// […] the nil slice is the preferred style.
//
// Note that there are limited circumstances where a non-nil but zero-length slice is preferred, such as when encoding JSON objects (a nil slice encodes to null, while []string{} encodes to the JSON array []).
//
// When designing interfaces, avoid making a distinction between a nil slice and a non-nil, zero-length slice, as this can lead to subtle programming errors.
// Reference: https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
func (p *PointSet) Reset() {
	p.Points = nil
}

// Clear clears the points slice but keeps the memory allocated
// and leaves the underlying array data and other values intact.
func (p *PointSet) Clear() {
	p.Points = p.Points[:0]
}

// MakeDataSeries creates a series of x, y points (with no data value)
// based on func, which is a function of the form f(x) = y, that
// accepts a single float64 value and returns a single float64 value.
//
// Values that are already in the PointSet are not overwritten.
//
// The x values are determined by the for loop values min, max, and step.
//
// e.g.
//
//  func y(x float64) float64 { return 5 * x + 3 }
//  ps.MakeDataSeries(2,50,2,y())
func (p *PointSet) MakeDataSeries(min, max, step float64, f func(float64) float64) {
	for i := min; i <= max; i += step {
		p.addPoint(&Point{i, f(i), nil})
	}
}

func (p *PointSet) PrintDataTable(w io.Writer, cols int) {
	if w == nil {
		w = os.Stderr
	}
	if cols == 0 {
		cols = 5
	}
	sb := strings.Builder{}
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf("PointSet name: %v\n", p.Name))
	sb.WriteString(fmt.Sprintf("PointSet color: %v\n", p.Color))
	sb.WriteString("PointSet values:\n")
	sb.WriteString("  ----------------\n")
	sb.WriteString(fmt.Sprintf("%6v : (%v,%v)\n", "n", "x", "y"))
	sb.WriteString("  ----------------\n")
	var s string
	for i, pt := range p.Points {
		s = fmt.Sprintf("%4d : (%0.1f,%0.1f)", i, pt.X, pt.Y)
		sb.WriteString(fmt.Sprintf("%-30s", s))
		if (i+1)%cols == 0 {
			sb.WriteString("\n")
		} else {
			fmt.Print("\t")
		}
	}
	sb.WriteString("\n")
	fmt.Fprint(w, sb.String())
}

// NewPointSet returns an empty point set.
//
// Default values are:
//  name: PointSet
//  color: black
//  unit: ""
func NewPointSet(name, color, unit string) *PointSet {
	if name == "" {
		name = "PointSet"
	}
	if color == "" {
		color = "black"
	}

	return &PointSet{
		Name:   name,
		Color:  color,
		Unit:   unit,
		Points: make([]*Point, 0),
	}
}
