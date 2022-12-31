package main

//enable for random structs
import (
	_ "algorithms/base/random"
	"fmt"
)

//33_search_in_rotated_sorted_array
func main() {
	fmt.Println(search([]int{3, 1}, 3))
}

func search(ar []int, t int) int {
	if len(ar) == 0 {
		return -1
	} else if len(ar) == 1 {
		if ar[0] == t {
			return 0
		}
		return -1
	}

	left, right := 0, len(ar)-1

	for left < right {
		mid := (left + right) >> 1

		if ar[mid] > ar[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if t >= ar[0] && t <= ar[left] {
		left, right = 0, left
		for left <= right {
			mid := (left + right) >> 1

			if ar[mid] == t {
				return mid
			} else if t > ar[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	} else {
		left, right = left, len(ar)-1
		for left <= right {
			mid := (left + right) >> 1

			if ar[mid] == t {
				return mid
			} else if t > ar[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

//func search(nums []int, target int) int {
//	if len(nums) == 0 {
//		return 0
//	} else if len(nums) == 1 {
//		if nums[0] == target {
//			return 0
//		}
//		return -1
//	}
//
//	var t int
//
//	for i := 0; i < len(nums); i++ {
//		if i+1 < len(nums) && nums[i] != nums[i+1]-1 {
//			if nums[i+1] >= target && nums[i] < target {
//				t = binary(nums[i+1:], target)
//				if t >= 0 {
//					return i + 1 + t
//				}
//				return -1
//			} else {
//				t = binary(nums[0:i+1], target)
//				if t >= 0 {
//					return i - t
//				}
//				return -1
//			}
//		}
//	}
//
//	return binary(nums, target)
//}
//
//func binary(ar []int, t int) int {
//	l, r := 0, len(ar)-1
//	for l <= r {
//		m := (l + r) / 2
//		if ar[m] == t {
//			return m
//		} else if t > ar[m] {
//			l = m + 1
//		} else {
//			r = m - 1
//		}
//	}
//	return -1
//}
