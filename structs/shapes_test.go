package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkShapes := func(t testing.TB, got, hasArea float64) {
		t.Helper()
		if got != hasArea {
			t.Errorf("got %.2f hasArea %.2f", got, hasArea)
		}
	}

	t.Run("a", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		got := Perimeter(rectangle)
		hasArea := 40.0
		checkShapes(t, got, hasArea)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
