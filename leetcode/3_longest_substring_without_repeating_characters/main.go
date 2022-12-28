package main

import (
    "fmt"
)

func main() {
    fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    } else if len(s) == 1 {
        return 1
    }

    m := make(map[byte]int)
    var sum int

    for left, right := 0, 0; right < len(s); right++ {
        if index, ok := m[s[right]]; ok {
            left = max(left, index)
        }
        sum = max(sum, right-left+1)
        m[s[right]] = right + 1
    }

    return sum
}

func max(n, m int) int {
    if n > m {
        return n
    }
    return m
}
