package points

func Series(start, end, step interface{}) []Point {
	return nil
}

func makeDataSeries(min, max float64, f func()) (p *PointSet) {
	for i := min; i <= max; i++ {
		p.Add(i, f(i))
	}
}

// Point represents an x,y data point with a data value.
type Point struct {
	x    float64
	y    float64
	data []byte
}

// PointSet represents a sequence of connected data points.
type PointSet struct {
	name   string   // name of this data set
	color  string   // display color for this set
	unit   string   // unit of measurement for display purposes
	points []*Point // points in the data set
}

func (p *PointSet) AddPoint(np *Point) {
	p.points = append(p.points, np)
}

func (p *PointSet) Add(x, y float64, data []byte) {
	p.AddPoint(&Point{x, y, data})
}

func NewPointSet(name, color, unit string, points []*Point) *PointSet {
	return &PointSet{
		name:   name,
		color:  color,
		unit:   unit,
		points: points,
	}
}
