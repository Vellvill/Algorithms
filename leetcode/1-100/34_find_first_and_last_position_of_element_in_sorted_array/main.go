package main

//enable for random structs
import (
	_ "algorithms/base/core/random"
	"fmt"
)

//34_find_first_and_last_position_of_element_in_sorted_array
func main() {
	fmt.Println(searchRange([]int{0, 0, 0}))
}

func searchRange(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	h := make(map[int]int)
	maxL := 0

	for i := range nums {
		if nums[i] == 1 {
			if _, ok := h[nums[i]]; ok {
				h[nums[i]]++
			} else {
				h[nums[i]] = 1
			}
			maxL = max(maxL, h[nums[i]])
		} else {
			delete(h, 1)
		}
	}

	return maxL
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
