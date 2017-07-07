package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	expected := 233168

	actual := solution()

	if actual != expected {
		t.Errorf("Expected was %d, but got %d", expected, actual)
	}
}
