package main

import (
    "fmt"
)

func main() {
    fmt.Println(duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0, 0}))
}

func duplicateZeros(in []int) []int {
    var n, c = make([]int, len(in)), 0
    for i, j := 0, 0; i < len(in); i += 1 {
        if in[i] == 0 {
            c++
            continue
        }
        n[j] = in[i]
        j++
    }
    return n[:len(in)-c]
}
