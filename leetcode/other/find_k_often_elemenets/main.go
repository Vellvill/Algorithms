package main

import (
    "fmt"
)

func main() {
    fmt.Println(merge([]int{1, 2, 3, 5, 7}, []int{1, 3, 5, 7, 9, 10}))
    fmt.Println(mergeSort([]int{3, 3, 5, 2, 4, 3, 6, 4}))
}

func merge(ar1, ar2 []int) []int {
    res, i, j := make([]int, 0, len(ar1)+len(ar2)), 0, 0

    for ; i < len(ar1) && j < len(ar2); {
        if ar1[i] < ar2[j] {
            res = append(res, ar1[i])
            i++
        } else {
            res = append(res, ar2[j])
            j++
        }
    }

    for ; i < len(ar1); i++ {
        res = append(res, ar1[i])
    }

    for ; j < len(ar2); j++ {
        res = append(res, ar2[j])
    }

    return res
}

func mergeSort(ar []int) []int {
    if len(ar) < 2 {
        return ar
    }

    l := mergeSort(ar[:len(ar)/2])
    r := mergeSort(ar[len(ar)/2:])

    return merge(l, r)
}
