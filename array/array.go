package array

import "fmt"

type Array struct {
	array []int
}

func MakeArray() *Array {
	return &Array{array: make([]int, 0)}
}

func (array *Array) Reverse() {
	for i, j := 0, len(array.array)-1; i <= j; i, j = i+1, j-1 {
		array.array[i], array.array[j] = array.array[j], array.array[i]
	}
}

func (array *Array) Append(n ...int) {
	for _, v := range n {
		array.array = append(array.array, v)
	}
}

func (a *Array) RemoveDuplicates() {

	m := make(map[int]int)

	for i, v := range a.array {
		if _, ok := m[v]; ok {
			if i == len(a.array) {
				a.array = a.array[:len(a.array)-1]
				break
			} else {
				a.array = append(a.array[:i], a.array[i+1:]...)
				i++
			}
		}

		m[v] = i
	}
}

// O(n)
func Interception(a, b []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, v := range a {
		m[v] = 1
	}
	for _, v := range b {
		if j, ok := m[v]; ok {
			if j == 1 {
				res = append(res, v)
				m[v]++
			}
		}
	}
	return res
}

func (a *Array) BinarySearch(t int) bool {
	first, last := 0, len(a.array)-1
	for first <= last {
		mid := (first + last) / 2
		if a.array[mid] == t {
			return true
		} else if t > mid {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return false
}

func Interception0N2(a, b []int) ([]int, error) {

	if len(a) == 0 || len(b) == 0 {
		return nil, fmt.Errorf("Nil slices %v, %v", a, b)
	}

	res := make([]int, 0)
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				res = append(res, aa)
			}
		}
	}
	return res, nil
}

func FindSum(a []int, b int) ([]int, error) {

	if len(a) < 2 {
		return nil, fmt.Errorf("len(%v) < 2 ", a)
	}

	m := make(map[int]int)
	res := make([]int, 0)
	for i, v := range a {
		if j, ok := m[v]; ok {
			res = append(res, j, i)
			return res, nil
		}

		m[b-v] = i
	}
	return nil, nil
}
