package main

import (
	"algorithms/base/random"
)

func main() {
	println(twoSum(
		random.CompareArrInt(random.GetArray(100, 10, 1)),
		random.GetInt(10)))
}

func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for i := range nums {
		if j, ok := h[target-nums[i]]; ok {
			return []int{i, j}
		}
		h[nums[i]] = i
	}
	return nil
}
