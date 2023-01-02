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
	mapa := make(map[byte]int)

	type freq struct {
		b byte
		f int
	}

	maxF := freq{}

	for i := 0; i < len(s); i++ {
		if _, ok := mapa[s[i]]; ok {
			mapa[s[i]]++
			if mapa[s[i]] > maxF.f {
				maxF = freq{
					b: s[i],
					f: mapa[s[i]],
				}
			}
		} else {
			mapa[s[i]] = 1
		}
	}

	left, right := 0, len(s)

	for i := 0; k > 0; i++ {
		if s[i] != maxF.b {
			k--
			left++
		}
	}

	fmt.Println(string(maxF.b), maxF.f)

	return right - left + 1
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
