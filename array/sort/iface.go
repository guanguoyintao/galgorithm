package sort

type Array interface {
	// Len 获取数据集合元素个数
	Len() int
	// Less 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
	Less(i, j int) bool
	// Swap 交换 i 和 j 索引的两个元素的位置
	Swap(i, j int)
}
