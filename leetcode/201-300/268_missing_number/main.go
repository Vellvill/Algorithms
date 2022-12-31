package main

import (
	"algorithms/base/core/random"
	"fmt"
)

func main() {
	var testCase []int
	l := random.GetInt(100) + 2
	for i := 1; i < l; i++ {
		testCase = append(testCase, i)
	} //порядок не важен, считается ар прогрессия
	testCase[random.GetInt(l-2)] = 0
	fmt.Println(testCase)
	fmt.Println(missingNumber(testCase))
}

func missingNumber(nums []int) int {
	var s int
	var n = len(nums)
	for i := range nums {
		s += nums[i]
	}
	return n*(n+1)/2 - s
}
