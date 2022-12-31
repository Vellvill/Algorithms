package main

//enable for random structs
import (
	"fmt"
)

//739_daily_temperatures

//Input: temperatures = [73,74,75,71,69,72,76,73]
//Output: [1,1,4,2,1,1,0,0]
//
//Input: temperatures = [30,40,50,60]
//Output: [1,1,1,0]
//
//Input: temperatures = [30,60,90]
//Output: [1,1,0]

func main() {
	fmt.Println(dailyTemperaturesON2([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
	fmt.Println(dailyTemperaturesON([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
}

type Pair struct {
	t     int
	index int
}

func dailyTemperaturesON2(temperatures []int) []int {
	res := make([]int, len(temperatures))

	for i := range temperatures {
		for j := i; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}

	return res
}

func dailyTemperaturesON(temperatures []int) []int {
	l := len(temperatures)

	stack := make([]Pair, 0, len(temperatures))
	res := make([]int, len(temperatures))

	for i := l - 1; i >= 0; i-- {
		if len(stack) != 0 {
			for currPair := stack[len(stack)-1]; len(stack) > 0; {
				currPair = stack[len(stack)-1]
				if currPair.t > temperatures[i] {
					res[i] = currPair.index - i
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		stack = append(stack, Pair{
			t:     temperatures[i],
			index: i,
		})
	}

	return res
}
