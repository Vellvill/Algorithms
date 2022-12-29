package main

//enable for random structs
import (
	_ "algorithms/base/random"
	"fmt"
)

//5_longest_palindromic_substring
func main() {
	fmt.Println(longestPalindrome("acacad"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		return s
	}

	var longest string

	for i := range s {
		longest = maxString(longest, findLongest(s, i))
	}

	return longest
}

func findLongest(s string, st int) (res string) {
	var l, r = st - 1, st + 1

	var temp string

	switch st % 2 {
	case 0:
		if l >= 0 && s[st] == s[l] {
			l -= 1
		} else if r < len(s) && s[st] == s[r] {
			r += 1
		}
		temp = findOdd(s, l, r)
	default:
		l, r = st-1, st+1
		res = string(s[st])
		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] == s[r] {
				res = s[l : r+1]
				continue
			}
			break
		}
	}

	res = maxString(res, temp)

	return res
}

func findOdd(s string, l, r int) (res string) {
	for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
		if s[l] == s[r] {
			res = s[l : r+1]
		} else {
			break
		}
	}
	return res
}

func maxString(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
