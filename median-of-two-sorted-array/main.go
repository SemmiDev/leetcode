package main

import (
	"sort"
)

func findMedianSortedArrays(nums1, nums2 []float64) float64 {
	nums1 = append(nums1, nums2...)

	sort.Float64s(nums1)

	if len(nums1) == 0 {
		return 0
	}

	if len(nums1)%2 == 0 {
		return (nums1[len(nums1)/2-1] + nums1[len(nums1)/2]) / 2
	}

	return nums1[len(nums1)/2]
}
