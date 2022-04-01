package main

import "fmt"

func main() {
	a := []int{3, 3, 6, 7, 3, 4, 5, 7, 8, 3, 4}
	hash := make(map[int]int)
	for _, v := range a {
		if _, ok := hash[v]; ok {
			hash[v]++
		} else {
			hash[v] = 1
		}
	}
	fmt.Println(hash) //will commited?
}
