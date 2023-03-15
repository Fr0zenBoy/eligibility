package logic

import (
	"testing"
)

func TestCalc(t *testing.T) {
	result := GetPercentege(4000, 5000)
	expected := 80.0

	if result != expected {
		t.Errorf("Result: %.2f, Expected: %.2f", result, expected)
	}
}
