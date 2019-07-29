package practice

import (
	"interstellar/pkg/algorithm/sort"
	"math"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
*/
//ThreeSum 三个数之和
func ThreeSum(nums []int) [][]int {
	sort.QSorting().Sorting(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		p1 := i + 1
		p2 := len(nums) - 1
		rep := map[int]bool{}
		for p1 < p2 {
			if nums[p1]+nums[p2] == -nums[i] {
				if _, ok := rep[nums[p1]]; ok {
					p1++
				} else {
					rep[nums[p1]] = true
					tmp := []int{
						nums[i],
						nums[p1],
						nums[p2],
					}
					res = append(res, tmp)
					p1++
					p2--
				}
			}
			if nums[p1]+nums[p2] > -nums[i] {
				p2--
			}
			if nums[p1]+nums[p2] < -nums[i] {
				p1++
			}
		}
	}
	return res
}

//ThreeSumClosest 三个数之和最接近
func ThreeSumClosest(nums []int, target int) int {
	sort.QSorting().Sorting(nums)
	tsum := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		p1 := i + 1
		p2 := len(nums) - 1
		for p1 < p2 {
			sum := nums[i] + nums[p1] + nums[p2]
			if math.Abs(float64(sum-target)) < math.Abs(float64(tsum-target)) {
				tsum = sum
			}
			if sum > target {
				p2--
			} else if sum < target {
				p1++
			} else {
				return tsum
			}
		}
	}
	return tsum
}
