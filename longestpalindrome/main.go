package main

import "fmt"

/*
最长 回文 子串
给你一个字符串 s，找到 s 中最长的回文子串。
示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 也是有效的答案。
示例 2：
输入：s = "cbbd"
输出："bb"
*/

func main() {
	s := "babad"
	result := longestPalindrome(s)
	fmt.Println(result)
}

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < n; i++ {
		len1 := expandAroundCenter(s, i, i)   // 奇数长度回文
		len2 := expandAroundCenter(s, i, i+1) // 偶数长度回文
		// 取较长的回文长度
		maxlen := len1
		if len2 > len1 {
			maxlen = len2
		}
		if maxlen > end-start {
			start = i - (maxlen-1)/2
			end = i + maxlen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
