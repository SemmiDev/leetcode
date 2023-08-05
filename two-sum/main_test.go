package main

import "testing"

func TestTwoSum(t *testing.T) {
	tc := []struct {
		arr    []int
		target int
		result []int
	}{
		{[]int{1, 3, 5, 6, 7, 8}, 9, []int{1, 3}},
		{[]int{1, 3, 5, 6, 7, 8}, 10, []int{1, 4}},
		{[]int{1, 3, 5, 6, 7, 8}, 12, []int{2, 4}},
		{[]int{1, 3, 5, 6, 7, 8}, 13, []int{3, 4}},
	}

	for _, v := range tc {
		result := twoSum(v.arr, v.target)
		if !compareSlice(result, v.result) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", v.arr, v.target, result, v.result)
		}
	}
}

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
