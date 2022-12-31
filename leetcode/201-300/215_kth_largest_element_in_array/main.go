package main

import (
	. "algorithms/base/core/random"
	"fmt"
)

func main() {
	arr := CompareArrInt(GetArray(GetInt(100), GetInt(100), TypeInt))
	fmt.Println(arr)
	res := findKthLargest(arr, 4)
	fmt.Println(arr)
	fmt.Println(res)
}

func findKthLargest(nums []int, k int) int {
	qs(nums)

	return nums[k-1]
}

//quicksort for sorting array
func qs(nums []int) []int {
	if len(nums) < 2 { //recursive || 1 el array will go back
		return nums
	}

	left, right := 0, len(nums)-1 //setting l and r indexes

	pivot := (left + right) >> 1 //seting pivot index (could be random)

	nums[pivot], nums[right] = nums[right], nums[pivot] //moving pivot to the right

	for i := range nums { //pushing ith elements to the left side of pivot if they are less than pivot
		if nums[i] > nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left] //moving pivot on smallest element

	qs(nums[:left])   //qs left part of array
	qs(nums[left+1:]) //qs right part of array

	return nums //returning result
}
