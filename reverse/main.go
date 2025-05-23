package main

import (
	"fmt"
	"math"
)

/*
整数反转

给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0

提示：整数的范围 [-2147483648, 2147483647]
*/

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		// 防止溢出
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		d := x % 10
		x /= 10
		res = res*10 + d
	}
	return res
}
