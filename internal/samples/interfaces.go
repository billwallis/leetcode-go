package samples

import (
	"fmt"
	"math"
)

// Shape is a 2D geometric shape which has a perimeter and an area.
type Shape interface {
	Perimeter() float64
	Area() float64
}

// Circle is a shape defined by its radius.
type Circle struct {
	Radius float64
}

func (circle *Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle *Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

// Triangle is a shape defined by its three sides.
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (triangle *Triangle) Perimeter() float64 {
	return triangle.SideA + triangle.SideB + triangle.SideC
}

func (triangle *Triangle) Area() float64 {
	// Heron's formula
	s := triangle.Perimeter() / 2
	return math.Sqrt(s * (s - triangle.SideA) * (s - triangle.SideB) * (s - triangle.SideC))
}

// Rectangle is a shape defined by its width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle *Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle *Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

// ShapeDetails prints the perimeter and area of a shape.
func ShapeDetails(shape Shape) string {
	perimeter, area := shape.Perimeter(), shape.Area()
	return fmt.Sprintf("Shape %T has perimeter %.2f and area %.2f\n", shape, perimeter, area)
}
