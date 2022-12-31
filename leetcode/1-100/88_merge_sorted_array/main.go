package main

//enable for random structs
import (
	_ "algorithms/base/random"
	"fmt"
)

//88_merge_sorted_array
func main() {
	a1 := []int{1, 2, 3, 0, 0, 0}
	merge(a1, 3, []int{2, 5, 6}, 3)
	fmt.Println(a1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for t, i, j := m+n-1, m-1, n-1; t >= 0; t-- {
		if i >= 0 && j >= 0 && nums1[i] > nums2[j] {
			nums1[t] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[t] = nums2[j]
			j--
		}
	}
}
