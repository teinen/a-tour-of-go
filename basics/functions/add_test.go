package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	actual := add(15, 25)
	expected := 40

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
