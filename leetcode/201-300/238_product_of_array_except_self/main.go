package main

//enable for random structs
import (
	_ "algorithms/base/random"
	"fmt"
)

//идея решения в том, чтобы получить префиксные произведения до i и постфиксные произведения после i
//записать их результаты в массивы pref и post и для каждого индекса записать их в нужное значение output

//238_product_of_array_except_self
func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

// [ 1 , 2 , 3 , 4] --> [  24       12         8        6  ]
//                      (2*3*4),  (1*3*4)   (1*2*4), (1*2*3)

//оптимизировать работу
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return nums
	} else if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}

	pref, post := make([]int, len(nums)), make([]int, len(nums))

	pref[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		pref[i] = nums[i] * pref[i-1]
	}

	post[len(nums)-1] = nums[len(nums)-1]

	for i := len(post) - 2; i >= 0; i-- {
		post[i] = nums[i] * post[i+1]
	}

	res := make([]int, len(nums))
	res[0] = post[1]

	var i, j = 0, 2
	var indx = 1

	for ; i < len(nums) && j < len(nums); i, j = i+1, j+1 {
		res[indx] = pref[i] * post[j]
		indx++
	}

	res[len(res)-1] = pref[len(res)-2]

	return res
}
