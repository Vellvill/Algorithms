package main

import (
    . "algorithms/base/random"
    "fmt"
)

func main() {
    arr := CompareArrInt(GetRandArray(GetRandomInt(100), GetRandomInt(100), TypeInt))
    fmt.Println(arr)
    res := findKthLargest(arr, 4)
    fmt.Println(arr)
    fmt.Println(res)
}

func findKthLargest(nums []int, k int) int {
    qs(nums)

    return nums[k-1]
}

func qs(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }

    left, right := 0, len(nums)-1

    pivot := (left + right) >> 1

    nums[pivot], nums[right] = nums[right], nums[pivot]

    for i := range nums {
        if nums[i] > nums[right] {
            nums[i], nums[left] = nums[left], nums[i]
            left++
        }
    }

    nums[left], nums[right] = nums[right], nums[left]

    qs(nums[:left])
    qs(nums[left+1:])

    return nums
}
