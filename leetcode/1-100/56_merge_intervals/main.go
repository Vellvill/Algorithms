package main

import (
    "fmt"
)

// Дано расписание, необхоидмо объединить его если оно совпадает
//{18, 28} {6, 10} {22, 25} {8, 15} {7 - 9} {30-36} --> {6,15} {18,28} {30,36}
//
//{2, 8} { 3, 9 } {12, 20} --> {2, 9}
//
//{2, 9} {12, 20} -> {2, 9} {12, 20}

func main() {
    ar := [][]int{{0, 1}, {2, 5}, {3, 7}}

    a := unionSchedule(ar)

    fmt.Println(a)

}

func unionSchedule(ar [][]int) [][]int {
    qs(ar)
    var last int
    for i := 1; i < len(ar); {
        if ar[i][1] <= ar[last][1] {
            if ar[i][0] < ar[last][0] {
                ar[last][0] = ar[i][0]
            }
            ar = append(ar[:i], ar[i+1:]...)
        } else if ar[i][0] >= ar[last][0] && ar[i][0] <= ar[last][1] {
            if ar[i][1] > ar[last][1] {
                ar[last][1] = ar[i][1]
            }
            ar = append(ar[:i], ar[i+1:]...)
        } else {
            last = i
            i++
        }
    }
    return ar
}

func qs(ar [][]int) [][]int {
    if len(ar) < 2 {
        return ar
    }

    l, r := 0, len(ar)-1

    p := (l + r) >> 1

    ar[p], ar[r] = ar[r], ar[p]

    for i := range ar {
        if ar[i][0] < ar[r][0] {
            ar[i], ar[l] = ar[l], ar[i]
            l++
        }
    }

    ar[l], ar[r] = ar[r], ar[l]

    qs(ar[:l])
    qs(ar[l+1:])

    return ar
}
