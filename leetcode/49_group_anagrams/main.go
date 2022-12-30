package main

//enable for random structs
import _ "algorithms/base/random"

//49_group_anagrams
func main() {

}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	m := make(map[string]*[]string)
	for i := range strs {
		sorted := string(sortBytes([]byte(strs[i])))
		if ar, ok := m[sorted]; ok {
			*ar = append(*ar, strs[i])
		} else {
			newAr := []string{strs[i]}
			m[sorted] = &newAr
		}
	}

	res := make([][]string, len(m))
	var i int
	for _, v := range m {
		res[i] = *v
		i++
	}

	return res
}

func sortBytes(s []byte) []byte {
	if len(s) < 2 {
		return s
	}

	l, r := 0, len(s)-1

	p := (l + r) >> 1

	s[r], s[p] = s[p], s[r]

	for i := range s {
		if s[i] < s[r] {
			s[i], s[l] = s[l], s[i]
			l++
		}
	}

	s[r], s[l] = s[l], s[r]

	sortBytes(s[:l])
	sortBytes(s[l+1:])

	return s
}
