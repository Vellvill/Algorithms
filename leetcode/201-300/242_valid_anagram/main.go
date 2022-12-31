package main

//enable for random structs
import _ "algorithms/base/random"

//242_valid_anagram
func main() {

}

func isAnagram(s string, t string) bool {
	m := make(map[byte]int)

	for i := range s {
		m[s[i]]++
	}

	for i := range t {
		if j, ok := m[t[i]]; ok && j == 1 {
			delete(m, t[i])
		} else if ok && j >= 1 {
			m[t[i]]--
		} else {
			return false
		}
	}
	return len(m) == 0
}
