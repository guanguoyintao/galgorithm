package string_matching

// kmp 的经典之处
// 生成 next 数组
func generateNext(p string) []int {
	m := len(p)
	next := make([]int, m, m)
	next[0] = -1
	next[1] = 0
	i, j := 0, 1 // 前缀子串、后缀子串起始位置
	// 因为是通过最长可匹配前缀子串计算，所以 j 的值最大不会超过 m-1
	for j < m-1 {
		if i == -1 || p[i] == p[j] {
			i++
			j++
			// 设置当前最长可匹配前缀子串结尾字符下标
			next[j] = i
		} else {
			// 如果 p[i] != p[j]，回到上一个最长可匹配前缀子串结尾字符下标
			i = next[i]
		}
	}
	return next
}

// KMPSearch KMP 算法实现函数
func KMPSearch(base, pattern string) (int, bool) {
	n := len(base)                // 主串长度
	m := len(pattern)             // 模式串长度
	next := generateNext(pattern) // 生成 next 数组
	i, j := 0, 0
	for i < n && j < m {
		// 如果主串字符和模式串字符不相等，
		// 更新模式串坏字符下标位置为好前缀最长可匹配前缀子串尾字符下标
		// 然后从这个位置重新开始与主串匹配
		// 相当于前面提到的把模式串向后移动 j - k 位
		if j == -1 || base[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		// 完全匹配，返回下标位置
		return i - j, true
	} else {
		return -1, false
	}
}
