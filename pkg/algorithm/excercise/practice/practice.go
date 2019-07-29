package practice

//MaxArea 最大面积
/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
*/
func MaxArea(height []int) int {
	lenIdx := len(height)
	maxCap := 0
	for i := 0; i < lenIdx; i++ {
		minTaller := 0
		fTaller := true
		maxLeft := 0
		for j := 0; j <= i; j++ {
			if fTaller && height[j] >= height[i] {
				minTaller = j
				fTaller = false
			}
			if height[j] <= height[i] && height[j]*(i-j) > maxLeft {
				maxLeft = height[j] * (i - j)
			}
		}
		if maxCap < (i-minTaller)*height[i] {
			maxCap = (i - minTaller) * height[i]
		}
		if maxCap < maxLeft {
			maxCap = maxLeft
		}
	}
	return maxCap
}

//MaxArea1 1
func MaxArea1(height []int) int {
	max := 0
	for p := 0; p < len(height); p++ {
		i := 0
		j := len(height) - 1
		for i <= p {
			if height[i] >= height[p] {
				break
			}
			i++
		}
		for j >= p {
			if height[j] >= height[p] {
				break
			}
			j--
		}
		maxP := (p - i) * height[p]
		if (j - p) > (p - i) {
			maxP = (j - p) * height[p]
		}
		if maxP > max {
			max = maxP
		}
	}
	return max
}

//MaxAreaByDualPointer 双指针
func MaxAreaByDualPointer(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		maxTmp := height[l] * (r - l)
		if height[l] > height[r] {
			maxTmp = height[r] * (r - l)
			r--
		} else {
			l++
		}
		if max < maxTmp {
			max = maxTmp
		}
	}
	return max
}
