package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2g want %.2g", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("triangles", func(t *testing.T) {
		triangle := Triangle{2, 2}
		checkArea(t, triangle, 2)
	})
	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{5.0, 10.0}
	// 	got := rectangle.Area()
	// 	want := 50.0

	// 	if got != want {
	// 		t.Errorf("got %.2g want %.2g", got, want)
	// 	}
	// })
	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{5.0}
	// 	got := circle.Area()
	// 	want := math.Pi * 25.0

	// 	if got != want {
	// 		t.Errorf("got %.2g want %.2g", got, want)
	// 	}
	// })
}
