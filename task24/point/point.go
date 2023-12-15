package point

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x, y,
	}
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

// Расстояние между двумя точками, используя теорему Пифагора AB^2 = AC^2 + CB^2
func (p Point) Distance(point *Point) float64 {
	return math.Sqrt(math.Pow(p.x-point.x, 2) + math.Pow(p.y-point.y, 2))
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%f, %f)", p.x, p.y)
}
