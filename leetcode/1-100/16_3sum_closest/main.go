package main

//enable for random structs
import (
	_ "algorithms/base/core/random"
	"fmt"
	"math"
	"math/rand"
)

//16_3sum_closest
func main() {
	a := []int{0, 0, 0}
	fmt.Println(threeSumClosest(a, 1))
}

func threeSumClosest(nums []int, target int) int {
	qs(nums)

	var maxDiff = 1<<32 - 1
	var res = -1

	for i := 0; i < len(nums); i++ {

		l, r := i+1, len(nums)-1

		for l < r {
			var sum = nums[i] + nums[l] + nums[r]

			if diff := int(math.Abs(float64(sum - target))); diff < maxDiff {
				maxDiff = diff
				res = sum
			}

			if sum == target {
				return target
			} else if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func qs(ar []int) {
	if len(ar) < 2 {
		return
	}

	l, r := 0, len(ar)-1

	p := rand.Int() % len(ar)

	ar[r], ar[p] = ar[p], ar[r]

	for i := range ar {
		if ar[i] < ar[r] {
			ar[i], ar[l] = ar[l], ar[i]
			l++
		}
	}

	ar[r], ar[l] = ar[l], ar[r]

	qs(ar[:l])
	qs(ar[l+1:])
}
