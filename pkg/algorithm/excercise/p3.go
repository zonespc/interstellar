package excercise

//FindDisorderArrayMid 查找无序数组的中位数
func FindDisorderArrayMid(a []int) float64 {
	low := 0
	high := len(a) - 1
	mid := high / 2
	div := findDisorderArrayMid(a, low, high)
	for div != mid {
		if div > mid {
			div = findDisorderArrayMid(a, low, div-1)
		} else {
			div = findDisorderArrayMid(a, div+1, high)
		}
	}
	if len(a)%2 == 0 {
		return float64((a[div] + a[div+1])) / 2.0
	}
	return float64(a[div])
}

func findDisorderArrayMid(a []int, low int, high int) int {
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
	a[low], a[j] = a[j], a[low]
	return j
}
