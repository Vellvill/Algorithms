package main

//enable for random structs
import _ "algorithms/base/core/random"

//128_longest_consecutive_sequence
func main() {

}

type hash struct {
	m map[int]bool
}

func Constructor() *hash {
	return &hash{m: make(map[int]bool)}
}

func longestConsecutive(nums []int) int {
	m := Constructor()
	for i := range nums {
		m.m[nums[i]] = true
	}

	var answer = 0

	for i := range nums {
		if v, _ := m.m[nums[i]]; v == true {
			t := 1
			m.ways(nums[i], &t)
			answer = max(answer, t)
		}
	}
	return answer
}

func (h *hash) ways(v int, totalCount *int) {

	h.m[v] = false

	var i, j bool

	i, _ = h.m[v+1]
	j, _ = h.m[v-1]

	if i, _ = h.m[v+1]; i == true {
		*totalCount++
		h.ways(v+1, totalCount)
	}
	if j, _ = h.m[v-1]; j == true {
		*totalCount++
		h.ways(v-1, totalCount)
	}

	return
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
