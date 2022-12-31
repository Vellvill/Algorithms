package main

//enable for random structs
import (
	_ "algorithms/base/random"
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

}

type Pair struct {
	t     int
	index int
}

//with built-in stack (with generics)
//func dailyTemperatures(temperatures []int) []int {
//	l := len(temperatures)
//
//	stack := make([]Pair, 0, len(temperatures))
//
//	for i := l - 1; i >= 0; i-- {
//		if len(stack) != 0 &&
//		stack = append(stack, Pair{
//			t:     temperatures[i],
//			index: i,
//		})
//	}
//
//	for i := range
//
//	return stack
//}
