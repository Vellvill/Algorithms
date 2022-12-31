package main

import (
    "strconv"
    "fmt"
    "strings"
)

func main() {
    fmt.Println(decodeString("3[a]2[bc2[a]]"))

}

func decodeString(s string) string {
    if len(s) == 0 {
        return ""
    }
    res, _ := decodeRec(s, 0)
    return res
}

func decodeRec(s string, idx int) (res string, end int) {
    var increaser int
    for ; idx < len(s); idx++ {
        if dig, err := strconv.Atoi(string(s[idx])); err == nil {
            increaser = increaser*10 + dig
        } else if s[idx] == '[' {
            dec, end := decodeRec(s, idx+1)
            res += strings.Repeat(dec, increaser)
            idx = end
            increaser = 0
        } else if s[idx] == ']' {
            return res, idx
        } else {
            res += string(s[idx])
        }
    }
    return res, end
}
