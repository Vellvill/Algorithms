package array

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
