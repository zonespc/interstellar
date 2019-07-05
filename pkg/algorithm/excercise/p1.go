package excercise

//MinSubarrayLen 返回数组中子数组和大于等于k的最小长度
func MinSubarrayLen(a []int, k int) int {
	Max := int(^uint(0) >> 1)
	ans := Max
	i := 0
	j := 0
	sum := 0
	for i < len(a) {
		if sum+a[i] < k {
			sum += a[i]
			i++
		} else {
			if i-j+1 < ans {
				ans = i - j + 1
			}
			sum -= a[j]
			j++
		}
	}
	if ans != Max {
		return ans
	}
	return 0
}
