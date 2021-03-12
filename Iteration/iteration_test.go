package Iteration

import "testing"

func TestRepeat(t *testing.T) {
	output := Repeat("a")
	expected := "aaaaa"

	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
