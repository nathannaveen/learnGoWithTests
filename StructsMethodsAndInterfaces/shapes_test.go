package StructsMethodsAndInterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := rectangle{10.0, 10.0}
	got := Perimeter(rectangle(rec))
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: rectangle{width: 12, height: 6}, want: 72.0},
		{shape: Circle{radius: 10}, want: 314.1592653589793},
		{shape: Triangle{base: 12, height: 6}, want: 36.0},
	}

	for _, i := range areaTests {
		got := i.shape.Area()
		if got != i.want {
			t.Errorf("got %g want %g", got, i.want)
		}
	}

}
