package sort

func partitionWithDig(list []int, low, high int) int {
	pivot := list[low] //导致 low 位置值为空
	for low < high {
		//high指针值 >= pivot high指针👈移
		for low < high && pivot <= list[high] {
			high--
		}
		//填补low位置空值
		//high指针值 < pivot high值 移到low位置
		//high 位置值空
		list[low] = list[high]
		//low指针值 <= pivot low指针👉移
		for low < high && pivot >= list[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		list[high] = list[low]
	}
	//pivot 填补 low位置的空值
	list[low] = pivot

	return low
}

func partitionWithSwap(list []int, low, high int) int {
	//每次取中间的值做为基准值.
	mid := list[(low+high)/2]
	//低位点与高位点必须重合就结束一轮的排序.
	for low < high {
		//如果高位点大于基准值,则一直向左偏移
		for list[high] > mid {
			high--
		}
		//如果低位点小于基准值,则一直向右偏移
		for list[low] < mid {
			low++
		}
		//低位与高位停止偏移,表示需要交换数据位置,并高低位都偏移一次.否则会出现死循环.
		if low < high { //code:(1)必须保证是low <= high可以重合, 如果low,high偏移交叉了不进行偏移.否则会出错.
			list[low], list[high] = list[high], list[low]
			low++
			high--
		}
	}

	return low
}

// QuickSortByDig 挖坑法
func QuickSortByDig(list []int, low, high int) []int {
	result := list
	if low < high {
		//位置划分
		pivot := partitionWithDig(result, low, high)
		//左边部分排序
		QuickSortByDig(result, low, pivot-1)
		//右边排序
		QuickSortByDig(result, pivot+1, high)
	}

	return result
}

// QuickSortBySwap 交换法
func QuickSortBySwap(list []int, low, high int) []int {
	result := list
	if low < high {
		//位置划分
		pivot := partitionWithSwap(result, low, high)
		//左边部分排序
		QuickSortBySwap(result, low, pivot-1)
		//右边排序
		QuickSortBySwap(result, pivot+1, high)
	}

	return result
}
