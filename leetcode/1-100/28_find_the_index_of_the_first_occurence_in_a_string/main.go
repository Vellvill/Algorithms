package main

import (
    "fmt"
)

func main() {
    fmt.Println(strStr("mississippi", "ippi3"))
}

func strStr(haystack string, needle string) int {
    if len(haystack) == 0 || len(needle) == 0 || haystack == needle {
        return 0
    }
    if len(needle) == 1 {
        for i := range haystack {
            if haystack[i] == needle[0] {
                return i
            }
        }
        return -1
    }
    for slow := 0; slow < len(haystack); slow++ {
        if haystack[slow] == needle[0] {
            fast := slow
            curr := 0
            for ; fast < len(haystack); fast, curr = fast+1, curr+1 {
                if curr == len(needle) {
                    return slow
                }
                if haystack[fast] == needle[curr] {
                    continue
                }
                break
            }
            if curr == len(needle) {
                return slow
            }
        }
    }
    return -1
}
