package samples_test

import (
	"github.com/Bilbottom/leetcode-go/internal/samples"
	"math"
	"testing"
)

func TestCirclePerimeter(t *testing.T) {
	c := samples.Circle{Radius: 5}
	want := 2 * 5 * math.Pi
	if c.Perimeter() != want {
		t.Errorf("%+v.Perimeter() = %.2f, want %.2f", c, c.Perimeter(), want)
	}
}

func TestCircleArea(t *testing.T) {
	c := samples.Circle{Radius: 5}
	want := math.Pi * math.Pow(5, 2)
	if c.Area() != want {
		t.Errorf("%+v.Area() = %.2f, want %.2f", c, c.Area(), want)
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tr := samples.Triangle{SideA: 3, SideB: 4, SideC: 5}
	want := 3.0 + 4.0 + 5.0
	if tr.Perimeter() != want {
		t.Errorf("%+v.Perimeter() = %.2f, want %.2f", tr, tr.Perimeter(), want)
	}
}

func TestTriangleArea(t *testing.T) {
	tr := samples.Triangle{SideA: 3, SideB: 4, SideC: 5}
	want := 6.0
	if tr.Area() != want {
		t.Errorf("%+v.Area() = %.2f, want %.2f", tr, tr.Area(), want)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := samples.Rectangle{Width: 3, Height: 4}
	want := 2.0 * (3 + 4)
	if r.Perimeter() != want {
		t.Errorf("%+v.Perimeter() = %.2f, want %.2f", r, r.Perimeter(), want)
	}
}

func TestRectangleArea(t *testing.T) {
	r := samples.Rectangle{Width: 3, Height: 4}
	want := 3.0 * 4.0
	if r.Area() != want {
		t.Errorf("%+v.Area() = %.2f, want %.2f", r, r.Area(), want)
	}
}

func TestShapeDetails(t *testing.T) {
	circle := &samples.Circle{Radius: 5}
	want := "Shape *samples.Circle has perimeter 31.42 and area 78.54\n"
	if got := samples.ShapeDetails(circle); got != want {
		t.Errorf("ShapeDetails(%+v) = %q, want %q", circle, got, want)
	}

	triangle := &samples.Triangle{SideA: 3, SideB: 4, SideC: 5}
	want = "Shape *samples.Triangle has perimeter 12.00 and area 6.00\n"
	if got := samples.ShapeDetails(triangle); got != want {
		t.Errorf("ShapeDetails(%+v) = %q, want %q", triangle, got, want)
	}

	rectangle := &samples.Rectangle{Width: 3, Height: 4}
	want = "Shape *samples.Rectangle has perimeter 14.00 and area 12.00\n"
	if got := samples.ShapeDetails(rectangle); got != want {
		t.Errorf("ShapeDetails(%+v) = %q, want %q", rectangle, got, want)
	}
}
