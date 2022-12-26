package main

import (
    "fmt"
)

func main() {
    a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
    rotate(a, 3)
    fmt.Println(a)
}

func rotate(ar []int, k int) {
    copy(ar, append(ar[len(ar)-k%len(ar):len(ar)], ar[:len(ar)-k%len(ar)]...))
}
