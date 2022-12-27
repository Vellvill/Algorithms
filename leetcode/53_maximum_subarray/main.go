package main

import (
    "fmt"
)

func main() {
    fmt.Println(maxSubArray([]int{-2, -4}))
}

func maxSubArray(ar []int) int {
    if len(ar) == 0 {
        return 0
    } else if len(ar) == 1 {
        return ar[0]
    }

    var (
        maxSum  = -1 << 31
        currSum int //можно math.MinInt( -2 ^ 31)
    )

    //реализуется метод скользящего окна + DP
    //у нас есть диапозон значений, мы бежим по массиву и прибавляем к текущей сумме 0
    //и сравниваем с максимальной (изначально равна минимальному значению int32 для отрицательных чисел(по условию задачи программа
    //компелится на 32х))
    //если currSum > maxSum, то делаем присвоение, иначе идем дальше и смотрим currSum < 0 и зануляем if true (нет смысла вычитать
    //из последующей суммы предыдущую)

    for i := 0; i < len(ar); i++ {
        currSum += ar[i]

        if currSum > maxSum {
            maxSum = currSum
        }

        if currSum < 0 {
            currSum = 0
        }
    }

    return maxSum
}
