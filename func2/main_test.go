package main

import "testing"

func TestMain(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}
}
