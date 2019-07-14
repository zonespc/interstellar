package excercise

//qSortPartition
func qSortPartition(a []int, low int, high int) int {
	if low >= high {
		return low
	}
	key := a[low]
	i := low + 1
	j := high
	for {
		for a[i] < key {
			if i >= high {
				break
			}
			i++
		}
		for a[j] > key {
			if j <= low {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	a[j], a[low] = a[low], a[j]
	return j
}

//FindKthMax 找到无序数组中第K大数字
func FindKthMax(a []int, K int) int {
	low := 0
	high := len(a) - 1
	mid := qSortPartition(a, low, high)
	for mid != K-1 {
		if mid > K-1 {
			mid = qSortPartition(a, low, mid-1)
		} else {
			mid = qSortPartition(a, mid+1, high)
		}
	}
	return a[mid]
}
