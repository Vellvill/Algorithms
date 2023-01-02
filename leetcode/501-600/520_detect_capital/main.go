package main

//enable for random structs
import (
	. "algorithms/base/core/random"
	"fmt"
	"strings"
)

//520_detect_capital

//We define the usage of capitals in a word to be right when one of the following cases holds:
//
//All letters in this word are capitals, like "USA".
//All letters in this word are not capitals, like "leetcode".
//Only the first letter in this word is capital, like "Google".
//Given a string word, return true if the usage of capitals in it is right.

func main() {
	for i := 0; i < 5; i++ {
		w := GetWord()
		fmt.Println(w)
		fmt.Println(detectCapitalUse(w))
	}
}

func detectCapitalUse(word string) bool {
	return word == strings.ToUpper(word) ||
		word == string(append([]byte(strings.ToUpper(string(word[0]))), strings.ToLower(word[1:])...)) ||
		word == strings.ToLower(word)

}
