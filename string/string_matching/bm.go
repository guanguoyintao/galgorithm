package string_matching

// BM BM算法，即Boyer-Moore算法。这是一种目前常用的字符串匹配算法。horspool算法是其简化版。该
// 算法除了使用了horspool算法里用到的根据字符最后出现位置挪动模式串的原理外，还利用已经匹配的
// 后缀序列：如果该序列在模式串出现多次，则挪动模式串对齐倒数第二次出现的位置；如果后缀在模式串
// 中只出现末尾那一次，则使用模式串的前缀去匹配该后缀序列的后缀，取其最长匹配序列对齐；如果两者
// 没有非空的匹配序列，则模式串与当前文本串匹配右侧末端位置后面一个字符对齐。BM算法利用这两个规
// 则，选择移动距离最长的来挪动模式串，从而提升匹配速度。改进算法将horspool对齐规则替换为更好
// 的sunday规则（horspool和sunday规则基本一致，只是对齐指示字符不同）。

// BadCharList 产生坏字符的失配移动表，j处失配，p串整体右移j-bc[T[j]]位
type badMatchTable struct {
	table   [256]int
	pattern string
}

func newBadMatchTable(pattern string) *badMatchTable {
	b := badMatchTable{
		pattern: pattern,
	}

	b.table = [256]int{}
	b.table = b.generateTable()

	return &b
}

func (b *badMatchTable) generateTable() [256]int {

	for i := 0; i < 256; i++ {
		b.table[i] = len(b.pattern)
	}

	lastPatternByte := len(b.pattern) - 1

	for i := 0; i < lastPatternByte; i++ {
		b.table[int(b.pattern[i])] = lastPatternByte - i
	}

	return b.table
}

func BMSearch(text string, pattern string) (int, bool) {
	byteText := []byte(text)
	bytePattern := []byte(pattern)

	textLength := len(byteText)
	patternLength := len(bytePattern)

	if textLength == 0 || patternLength == 0 || patternLength > textLength {
		return -1, false
	}

	lastPatternByte := patternLength - 1

	mt := newBadMatchTable(pattern)
	index := 0
	for index <= (textLength - patternLength) {
		for i := lastPatternByte; byteText[index+i] == bytePattern[i]; i-- {
			if i == 0 {
				return index, true
			}
		}
		index += mt.table[byteText[index+lastPatternByte]]
	}

	return -1, false
}
