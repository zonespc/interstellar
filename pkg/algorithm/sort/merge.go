package sort

//MergeSort 归并排序
func MergeSort(a []int) {
	MergeGroup(a, 0, len(a)-1)
}

//MergeGroup 分组
func MergeGroup(a []int, start int, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	MergeGroup(a, start, mid)
	MergeGroup(a, mid+1, end)
	MergeTwoSort(a, start, mid, end)
}

//MergeTwoSort 合并排好序的序列
func MergeTwoSort(a []int, start int, mid int, end int) {
	a1 := make([]int, mid-start+1)
	copy(a1, a[start:mid+1])
	a2 := make([]int, end-mid)
	copy(a2, a[mid+1:end+1])
	i, j := 0, 0
	cnt := start
	for i < len(a1) || j < len(a2) {
		if i >= len(a1) {
			for j < len(a2) {
				a[cnt] = a2[j]
				cnt++
				j++
			}
			break
		}
		if j >= len(a2) {
			for i < len(a1) {
				a[cnt] = a1[i]
				cnt++
				i++
			}
			break
		}
		if a1[i] < a2[j] {
			a[cnt] = a1[i]
			i++
		} else {
			a[cnt] = a2[j]
			j++
		}
		cnt++
	}
}
