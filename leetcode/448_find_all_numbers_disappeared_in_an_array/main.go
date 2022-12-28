package main

//enable for random structs
import (
	_ "algorithms/base/random"
	"fmt"
)

//448_find_all_numbers_disappeared_in_an_array
func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 4, 3, 4, 4, 4, 4}))
}

func findDisappearedNumbers(a []int) []int {
	for start := 0; start < len(a); {
		if a[start] != a[a[start]-1] {
			a[start], a[a[start]-1] = a[a[start]-1], a[start]
			continue
		}
		start++
	}

	var res []int

	for i := range a {
		if a[i] != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
