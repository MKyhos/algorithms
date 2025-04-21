package rtree

import "math"

type (
	Point     []float64
	Rectangle struct {
		Min, Max Point
	}
	Rectangles []Rectangle
)

func NewRectangle(p Point, epsilon float64) Rectangle {
	min := make(Point, len(p))
	max := make(Point, len(p))
	for i := range p {
		min[i] = p[i] - epsilon
		max[i] = p[i] + epsilon
	}
	return Rectangle{Min: min, Max: max}
}

func (r *Rectangle) Contains(p Point) bool {
	for i := range p {
		if p[i] < r.Min[i] || p[i] > r.Max[i] {
			return false
		}
	}
	return true
}

func (r *Rectangle) Overlaps(other *Rectangle) bool {
	for i := range r.Min {
		if r.Max[i] < other.Min[i] || r.Min[i] > other.Max[i] {
			return false
		}
	}
	return true
}

func (r *Rectangle) Overlap(other *Rectangle) float64 {
	overlap := 1.0
	for i := range r.Min {
		overlapDim := math.Min(r.Max[i], other.Max[i]) - math.Max(r.Min[i], other.Min[i])
		if overlapDim <= 0 {
			return 0
		}
		overlap *= overlapDim
	}
	return overlap
}

func (r *Rectangle) Enlarge(other *Rectangle) {
	for i := range r.Min {
		r.Min[i] = math.Min(r.Min[i], other.Min[i])
		r.Max[i] = math.Max(r.Max[i], other.Max[i])
	}
}

func (r *Rectangle) GetCentroid(axis int) float64 {
	return (r.Min[axis] + r.Max[axis]) / 2
}

func (r *Rectangle) Area() float64 {
	area := 1.0
	for i := range r.Min {
		area *= (r.Max[i] - r.Min[i])
	}
	return area
}
