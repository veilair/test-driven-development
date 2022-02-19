package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 40.0}
	got := rectangle.Perimeter()
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 15, Height: 6}, hasArea: 45.0},
	}
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// }
	// t.Run("rectangles", func(t *testing.T) {

	// 	rectangle := Rectangle{10.0, 40.0}
	// 	got := rectangle.Area()
	// 	want := 400.00

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 		checkArea(t, rectangle, 400.0)
	// 	}
	// })

	// t.Run("Circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	got := circle.Area()
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 		checkArea(t, circle, 314.1592653589793)
	// 	}

	// })

	t.Run("Table Driven Tests", func(t *testing.T) {
		for _, tt := range areaTests {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %.2f want %.2f", got, tt.hasArea)
			}
		}
	})
}
