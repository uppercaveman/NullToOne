package main

import (
	"fmt"
)

/*
寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log (m+n))。
你可以假设 nums1 和 nums2 不会同时为空。
示例 1:
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2 是位于中间的数
示例 2:
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 2.5 是位于中间的两个数的平均值
*/

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := mergeSortedArrays(nums1, nums2)
	n := len(merged)
	if n%2 == 1 {
		return float64(merged[n/2])
	}
	return float64(merged[n/2-1]+merged[n/2]) / 2
}

func mergeSortedArrays(nums1 []int, nums2 []int) []int {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	merged = append(merged, nums1[i:]...)
	merged = append(merged, nums2[j:]...)
	return merged
}
