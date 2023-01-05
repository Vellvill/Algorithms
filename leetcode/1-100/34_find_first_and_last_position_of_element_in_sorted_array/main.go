package main

//enable for random structs
import (
	_ "algorithms/base/core/random"
	"fmt"
)

//34_find_first_and_last_position_of_element_in_sorted_array
func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)>>1

		if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		} else if nums[m] == target {

			leftOc, rightOc := m, m

			if m+1 < len(nums) && nums[m+1] == target {
				l = m
				r = len(nums) - 1

				for l <= r {
					m := l + (r-l)>>1

					if nums[l] == target && l+1 < len(nums) && nums[l+1] == target {
						l++
					} else if nums[l] == target {
						rightOc = l
						break
					} else if target < nums[m] {
						r = m - 1
					}
				}

			}

			if m-1 >= 0 && nums[m-1] == target {
				r = m
				l = 0
				for l <= r {
					m := l + (r-l)>>1

					if nums[r] == target && r-1 >= 0 && nums[r-1] == target {
						r--
					} else if nums[r] == target {
						leftOc = r
						break
					} else if target > nums[m] {
						l = m + 1
					}
				}
			}

			return []int{leftOc, rightOc}
		}
	}
	return []int{-1, -1}
}
