package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2}, 3},
		test{[]int{4, 5, 3}, 12},
		test{[]int{10, 22, 43}, 75},
		test{[]int{9, 5, 2}, 16},
	}
	for _, x := range tests {
		v := mySum(x.data...)
		if v != x.answer {
			t.Error("EXPECTED", x.answer, "GOT", v)
		}

	}
}
