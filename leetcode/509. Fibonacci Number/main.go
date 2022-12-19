package main

import (
    "fmt"
)

func main() {
    var a int
    _, _ = fmt.Scan(&a)
    fibRec(a)
}

//rec
func fibRec(n int) int {
    if n <= 1 {
        return n
    }
    return fibRec(n-1) + fibRec(n-2)
}

//non rec
func fibNoRec(n int) int {
    var f = make([]int, n+1, n+2)
    if n < 2 {
        f = f[0:2]
    }
    f[0], f[1] = 0, 1
    for i := 2; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[n]
}
