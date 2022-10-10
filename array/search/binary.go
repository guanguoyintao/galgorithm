package search

func BinarySearch(sortedArray []int, lookingFor int) int {
	var low int = 0
	var high int = len(sortedArray) - 1
	for low <= high {
		var mid int = low + (high-low)/2
		var midValue int = sortedArray[mid]
		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
