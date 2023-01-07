package main

//enable for random structs
import _ "algorithms/base/core/random"

//35_search_insert_position
func main() {

}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)>>1

		if nums[m] == target {
			return m
		} else if l == r {
			if target < nums[l] {
				return l
			} else if target > nums[l] {
				return l + 1
			}
		} else if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		}
	}

	return l
}
