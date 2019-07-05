package excercise

import "math"

//MaxIdenticalLen 返回两个数组中子数组相同的最大长度
func MaxIdenticalLen(A []int, B []int) int {
	var dp [][]int
	n := len(A)
	m := len(B)
	for i := 0; i < n+1; i++ {
		tmp := make([]int, m+1)
		dp = append(dp, tmp)
	}
	length := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i] != B[j] {
				dp[i+1][j+1] = 0
			} else {
				dp[i+1][j+1] = dp[i][j] + 1
				length = int(math.Max(float64(length), float64(MaxArray2Value(dp))))
			}
		}
	}
	return length
}

//MaxArray2Value 二维数组最大值
func MaxArray2Value(d2 [][]int) int {
	max := 0
	for i := 0; i < len(d2); i++ {
		rmax := MaxArrayValue(d2[i])
		if max < rmax {
			max = rmax
		}
	}
	return max
}

//MaxArrayValue 一维数组最大值
func MaxArrayValue(d1 []int) int {
	max := 0
	for i := 0; i < len(d1); i++ {
		if max < d1[i] {
			max = d1[i]
		}
	}
	return max
}
