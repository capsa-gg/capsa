package main

import "testing"

func TestExample(t *testing.T) {
	result := Example()
	if result != 0 {
		t.Error("unexpected failure")
	}
}
