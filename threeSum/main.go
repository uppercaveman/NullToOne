package main

import (
	"fmt"
	"slices"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。

示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。

示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。

提示：

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	ln := len(nums)
	for i := range ln - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, ln-1
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < ln-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			}
			if sum > 0 {
				r--
			}
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return res
}
