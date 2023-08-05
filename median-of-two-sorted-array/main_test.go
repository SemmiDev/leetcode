package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tc := []struct {
		nums1  []float64
		nums2  []float64
		result float64
	}{
		{[]float64{1, 3}, []float64{2}, 2},
		{[]float64{1, 2}, []float64{3, 4}, 2.5},
		{[]float64{0, 0}, []float64{0, 0}, 0},
		{[]float64{}, []float64{1}, 1},
		{[]float64{2}, []float64{}, 2},
	}

	for _, v := range tc {
		result := findMedianSortedArrays(v.nums1, v.nums2)
		if result != v.result {
			t.Errorf("findMedianSortedArrays(%v, %v) = %f, want %f", v.nums1, v.nums2, result, v.result)
		}
	}
}
