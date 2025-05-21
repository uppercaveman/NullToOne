package main

import (
	"slices"
)

/*
最近的三数之和
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0
解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。

提示：

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104

*/

func main() {
	println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	println(threeSumClosest([]int{0, 0, 0}, 1))
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	slices.Sort(nums)
	ln := len(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := range ln - 2 {
		l, r := i+1, ln-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			// 这里的「最接近」即为差值的绝对值最小
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				l++
			}
			if sum > target {
				r--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
