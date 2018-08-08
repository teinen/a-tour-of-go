package main

import "testing"

func TestSplit(t *testing.T) {
	actualX, actualY := split(25)
	expectedX, expectedY := 11, 14

	if actualX != expectedX {
		t.Errorf("got %v, want %v", actualX, expectedX)
	}

	if actualY != expectedY {
		t.Errorf("got %v, want %v", actualY, expectedY)
	}
}
