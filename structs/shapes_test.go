package structs

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{
			name:         "Rectangle",
			shape:        Rectangle{Width: 10, Height: 10},
			hasPerimeter: 40,
		},
		{
			name:         "Circle",
			shape:        Circle{Radius: 10},
			hasPerimeter: 62.83185307179586,
		},
		{
			name:         "Isosceles Triangle",
			shape:        IsoscelesTriangle{Base: 6.0, Height: 12.0},
			hasPerimeter: 30,
		},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.hasPerimeter {
			t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasPerimeter)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "Rectangle",
			shape:   Rectangle{Width: 12, Height: 6},
			hasArea: 72,
		},
		{
			name:    "Circle",
			shape:   Circle{Radius: 10},
			hasArea: 314.1592653589793,
		},
		{
			name:    "Isosceles Triangle",
			shape:   IsoscelesTriangle{Base: 12, Height: 6},
			hasArea: 36.0,
		},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %.2f hasArea %.2f", tt.name, got, tt.hasArea)
		}
	}
}
