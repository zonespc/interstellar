package search

//BinarySearch bs
func BinarySearch(a []int, key int) (ok bool, index int) {
	if a == nil {
		return false, -1
	}
	low := 0
	high := len(a) - 1
	middle := low + (high-low)/2
	for low <= high && a[middle] != key {
		if a[middle] < key {
			low = middle + 1
		} else {
			high = middle - 1
		}
		middle = low + (high-low)/2
	}
	if low <= high {
		return true, middle
	}
	return false, -1
}
