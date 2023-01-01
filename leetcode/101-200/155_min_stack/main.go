package main

//enable for random structs
import (
	_ "algorithms/base/core/random"
	"fmt"
)

func main() {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.Top())
	fmt.Println(m.GetMin())
}

type MinStack struct {
	stack   []int
	minVals []int
	minVal  int
}

func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		minVal: 1<<63 - 1,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if val < this.minVal {
		this.minVal = val
	}
}

func (this *MinStack) Pop() {
	l := len(this.stack)
	if l > 0 {
		if this.stack[l-1] == this.minVal && l > 1 {

			var minSum, currSum = 1<<63 - 1, 0
			var newMinIndx int

			for i := 0; i < l; i++ {
				currSum = this.minVal + this.stack[i]
				if currSum < minSum {
					minSum = currSum
					newMinIndx = i
				}
			}
			this.minVal = this.stack[newMinIndx]
		}

		this.stack = this.stack[:len(this.stack)-1]

	}

	if len(this.stack) == 0 {
		this.minVal = 1<<63 - 1
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal
}
