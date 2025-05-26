package main

import (
	"fmt"
	"slices"
)

/*
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同

示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

提示：

1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}, 0))
}

func fourSum(nums []int, target int) [][]int {
	// 排序
	slices.Sort(nums)
	ln := len(nums)
	res := make([][]int, 0)
	for i := range ln - 3 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < ln-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, ln-1
			for l < r {
				if l > j+1 && nums[l] == nums[l-1] {
					l++
					continue
				}
				if r < ln-1 && nums[r] == nums[r+1] {
					r--
					continue
				}
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
					continue
				}
				if sum > target {
					r--
					continue
				}
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				}
			}
		}
	}
	return res
}
