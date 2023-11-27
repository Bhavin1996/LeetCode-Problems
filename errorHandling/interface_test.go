package main

import (
	"testing"
)

func NewRect() rectangle {
	return rectangle{
		length: 10.0,
		height: 5,
	}
}

func TestCalculate(t *testing.T) {
	rect := rectangle{
		length: 5,
		height: 10,
	}

	expectedArea := 60

	result := rect.Area()

	if result != expectedArea {
		t.Errorf("Expected area: %d, got: %d", expectedArea, result)
	}

	// Now, test the Calculate function
	Calculate(rect)
	// Note: To capture the printed output, you might need to redirect stdout in your test
}
