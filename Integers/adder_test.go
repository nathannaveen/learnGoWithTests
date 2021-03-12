package Integers

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("got '%d' but expected '%d'", sum, expected)
	}
}

func TestSumTwoExample(t *testing.T) {
	sum := Add(1, 5)
	expected := 6
	if sum != expected {
		t.Errorf("got '%d' but expected '%d'", sum, expected)
	}
}
