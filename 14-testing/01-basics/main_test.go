package main

import "testing"

func TestMySum(t *testing.T) {
	v := mySum(10, 2)
	if v != 12 {
		t.Error("Expected", 12, "Got", v)
	}
}
