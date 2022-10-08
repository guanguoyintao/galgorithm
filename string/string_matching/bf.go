package string_matching

func BFSearch(base, pattern string) (int, bool) {
	for i := range base {
		for j := range pattern {
			if i+j == len(base) {
				return -1, false
			}
			if base[i+j] != pattern[j] {
				break
			}
			if j == len(pattern)-1 {
				return i, true
			}
		}
	}
	return -1, false
}
