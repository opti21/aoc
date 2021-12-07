package utils

import (
	"testing"
)

func TestBinaryStrToInt(t *testing.T) {
	testInt := BinaryStringToInt("10110")
	var expected int64 = 22

	if testInt !=  expected {
		t.Errorf("expected %q got %q", testInt, expected)
	}
}