package main

//enable for random structs
import (
	_ "algorithms/base/core/random"
	"fmt"
)

//15_3sum
func main() {
	fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}))
}

func threeSum(nums []int) (res [][]int) {

	nums = mergeSort(nums)
	lenght := len(nums)

	for i := 0; i < lenght; i++ {
		if i > 0 && nums[i] == nums[i-1] { //убрать дубли
			continue
		}

		for l, r := i+1, lenght-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]

			if sum > 0 { //сдвинуть правую если больше нуля
				r--
			} else if sum < 0 { //сдвинуть левую если меньше нуля
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]}) //пушим в итоговый массив резалт
				l++                                                 //сдвигаем левую
				for nums[l] == nums[l-1] && l < r {                 //до тех пор пока не будет равна прыдыщему l (избавится от дублей)
					l++
				}
			}
		}

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

func merge(ar1, ar2 []int) []int {
	i, j := 0, 0
	res := make([]int, 0, len(ar1)+len(ar2))
	for i < len(ar1) && j < len(ar2) {
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
