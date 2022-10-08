package sort

func buildHeep(nums []int, len int) {
	// 找到最后一个节点的父节点
	parent := len/2 - 1
	for parent >= 0 {
		heapify(nums, parent, len)
		parent--
	}
}

func heapify(nums []int, parent, len int) {
	// 判断两个子节点是否比父节点大，如果是的话替换
	top := parent
	leftChild := parent*2 + 1
	rightChild := parent*2 + 2
	if rightChild < len && nums[leftChild] > nums[top] {
		// 左节点是否大于父节点
		top = leftChild
	}
	if rightChild < len && nums[rightChild] > nums[top] {
		// 右节点是否大于父节点
		top = rightChild
	}
	if parent != top {
		nums[top], nums[parent] = nums[parent], nums[top]
		heapify(nums, top, len)
	}
}

func HeapSort(nums []int) []int {
	// 堆排序,只能确认第一次个数是最大或最小的
	// 调换第一个元素和最后一个元素位置、从0倒数第二个继续堆排序
	i := len(nums)
	for i > 1 {
		buildHeep(nums, i)
		nums[0], nums[i-1] = nums[i-1], nums[0]
		i--
	}

	return nums
}
