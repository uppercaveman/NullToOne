package main

import (
	"fmt"
)

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

*/

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

// 滑动窗口法
func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLen := 0
	m := [128]int{}
	for k, v := range s {
		m[v]++
		for m[v] > 1 {
			m[s[left]]--
			left++
		}
		if k-left+1 > maxLen {
			maxLen = k - left + 1
		}
	}
	return maxLen
}
