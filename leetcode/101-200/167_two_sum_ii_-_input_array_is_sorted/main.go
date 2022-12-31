package main

//enable for random structs
import (
	"algorithms/base/core/tests"
	"fmt"
)

type testCase struct {
	ar     []int
	target int
	output []int
}

var testCases = []testCase{
	{
		ar:     []int{2, 7, 11, 15},
		target: 9,
		output: []int{1, 2},
	},
	{
		ar:     []int{2, 3, 4},
		target: 6,
		output: []int{1, 3},
	},
	{
		ar:     []int{-1, 0},
		target: -1,
		output: []int{1, 2},
	},
}

//Input: numbers = [2,7,11,15], target = 9
//Output: [1,2]
//Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

//Input: numbers = [2,3,4], target = 6
//Output: [1,3]
//Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

//Input: numbers = [-1,0], target = -1
//Output: [1,2]
//Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

//167_two_sum_ii_-_input_array_is_sorted
func main() {
	for i := range testCases {
		fmt.Println(tests.IsEqual(twoSum(testCases[i].ar, testCases[i].target), testCases[i].output))
	}
}

// O(1) extra space
func twoSum(numbers []int, target int) []int {
	var l, r = 0, len(numbers) - 1

	for {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
}
