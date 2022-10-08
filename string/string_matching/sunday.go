package string_matching

func isin(elem rune, target string) (bool, int) {
	for i, i2 := range target {
		if i2 == elem {
			return true, i
		}
	}
	return false, 0
}

func SundaySearch(base, pattern string) (int, bool) {
	// 首先排除特殊情况
	if len(base) == 0 {
		return -1, false
	}
	// base = "hello"  pattern = "" return true
	if len(pattern) == 0 {
		return -1, false
	}
	// 其他的情况
	index := 0

	length := len(pattern)

	for index <= len(base)-length {
		pindex := 0
		offset := 0
		for pindex < length && pattern[pindex] == base[pindex+index] {
			// 如果找到一个不相同的 则直接将index 右移 length + 1
			pindex += 1
		}
		if pindex == length {
			return index, true
		}

		if pindex < length {
			if ok, value := isin(rune(base[index+length]), pattern); ok {
				offset = length - value
			} else {
				offset = length + 1
			}
		}

		index += offset
	}

	return index, false
}
