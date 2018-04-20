package Structs_Methods_Interfaces

import (
	"testing"
)

type Shape interface {
	Area() float64
}

func checkArea(t *testing.T, s Shape, expected float64) {
	t.Helper()
	if actual := s.Area(); actual != expected {
		t.Errorf("%#v Expected: %.2f Actual: %.2f", s, expected, actual)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	t.Run("Test the calculating of perimeter", func(t *testing.T) {
		actual := rectangle.Perimeter()
		expected := 40.0

		if actual != expected {
			t.Errorf("Expected: %.2f \nActual: %.2f", expected, actual)
		}
	})
}


func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		expected float64
	}{
		{ "Test Rectangle area",Rectangle{10.0, 10.0}, 100.0},
		{"Test Circle area", Circle{10}, 314.1592653589793},
		{"Test Triangle area", Triangle{12,6}, 36},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.expected)
		})
	}
}
