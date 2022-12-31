package main

//enable for random structs
import (
	_ "algorithms/base/random"
	"fmt"
)

var testCases = []string{
	"()",
	"()[]{}",
	"(]",
}

//20_valid_parentheses
func main() {
	for i := range testCases {
		fmt.Println(
			fmt.Sprintf(
				"case: %s, isvalid: %v",
				testCases[i],
				isValid(testCases[i])))
	}
}

func isValid(s string) bool {
	var stack []byte

	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for i := range s {
		if _, ok := m[s[i]]; ok {
			stack = append(stack, s[i])
			continue
		} else {
			if len(stack) == 0 {
				return false
			}
			if s[i] == m[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
