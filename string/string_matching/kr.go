package string_matching

const primeRK = 16777619

// hash散列方法， 返回字符串hash以及 primeRK的k-1（len(sep)-1）次方
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i]) // 循环得到字符串hash
	}

	// 位运算巧妙的获取 primeRK 的 len(sqp)-1 次方
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func KRSearch(base, pattern string) (int, bool) {
	// Rabin-Karp search
	hashss, pow := hashStr(pattern)
	n := len(pattern)
	var h uint32
	// 计算目标字符串前n位hash并与待匹配字符串hash进行对比
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(base[i])
	}
	// hash相同并且字符串相等则返回当前位置下标
	if h == hashss && base[:n] == pattern {
		return 0, true
	}

	// Rabin-Karp 算法的精华所在，相面详细介绍
	for i := n; i < len(base); {
		h *= primeRK
		h += uint32(base[i])
		h -= pow * uint32(base[i-n])
		i++
		if h == hashss && base[i-n:i] == pattern {
			return i - n, true
		}
	}
	return -1, true
}
