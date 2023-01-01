package main

//enable for random structs
import (
	_ "algorithms/base/core/random"
	"fmt"
)

//424_longest_repeating_character_replacement
func main() {
	type testCase struct {
		str string
		k   int
	}
	testCases := []testCase{
		{
			str: "ABAB",
			k:   2,
		},
		{
			str: "AABABBA",
			k:   1,
		},
	}
	for i := range testCases {
		fmt.Println(characterReplacement(testCases[i].str, testCases[i].k))
	}
}

func characterReplacement(s string, k int) int {
	hash := make(map[byte]int)
	lenghth := len(s)
	l, r := 0, 0
	var longest int

	for l < lenghth {
		currL := 0
		temp := k
		currChar := s[l]
		var firstOccusion byte
		for ; r < lenghth && temp > 0; r++ {
			if s[r] != currChar {
				temp--
				if temp == k-1 {
					firstOccusion = s[r]
					hash[s[r]] = r
				}
			}
			currL = max(currL, r-l+1)
		}
		longest = max(longest, currL)

		l = hash[firstOccusion]

	}

	return longest
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
