package main

/*
	arr = [1,3,5,6,7,8]
	target = 9
*/

func twoSum(arr []int, target int) []int {
	// k-> arr value, v-> arr index
	mapper := make(map[int]int)

	for i, v := range arr {
		sub := target - v
		if _, ok := mapper[sub]; ok {
			return []int{mapper[sub], i}
		} else {
			mapper[v] = i
		}
	}
	return []int{}
}
