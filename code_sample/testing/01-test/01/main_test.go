package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2, 3, 4}, 10},
		test{[]int{1, 2, 4, 4}, 11},
		test{[]int{1, 2, 2, 4}, 9},
		test{[]int{0, 0, 3, 4}, 7},
	}

	for _, td := range tests {
		x := mySum(td.data...)
		if x != td.answer {
			t.Error("Expected", td.answer, "Got", x)
		}
	}

}
