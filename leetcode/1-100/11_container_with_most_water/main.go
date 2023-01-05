package main

//enable for random structs
import _ "algorithms/base/core/random"

//Input: height = [1,8,6,2,5,4,8,3,7]
//Output: 49

//11_container_with_most_water
func main() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}

func maxArea(height []int) int {

	s, e := 0, len(height)-1

	var maxSquare int

	for s < e {
		maxSquare = max(maxSquare,
			(e-s)*
				min(height[e], height[s]))

		if height[e] < height[s] {
			e--
		} else {
			s++
		}
	}

	return maxSquare
}

func min(n, m int) int {
	if n > m {
		return m
	}
	return n
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
