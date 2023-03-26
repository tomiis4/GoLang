package main

import "testing"

func TestAbs(t *testing.T) {
	// -1, 0, 1

	if abs(-1) < 0 {
		t.Error("Negative value was found in abs() with", -1)
	}
	if abs(0) < 0 {
		t.Error("Negative value was found in abs() with", 0)
	}
	if abs(1) < 0 {
		t.Error("Negative value was found in abs() with", 1)
	}
}
