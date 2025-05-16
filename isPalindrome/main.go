package main

import "fmt"

/*
回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
示例 1：
输入：x = 121
输出：true
示例 2：
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
提示：
-231 <= x <= 231 - 1
*/

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println("-----------------------")
	fmt.Println(isPalindrome1(121))
	fmt.Println(isPalindrome1(-121))
	fmt.Println(isPalindrome1(10))
}

// 解法一：将整数反转，然后判断反转后的整数是否与原整数相等
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	m := 0
	n := x
	for n > 0 {
		m = m*10 + n%10
		n /= 10
	}
	return m == x
}

// 解法二：将整数转换为字符串，然后使用双指针判断是否为回文数
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	s := fmt.Sprintf("%d", x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
