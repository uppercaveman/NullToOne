package main

import (
	"fmt"
	"math"
)

/*
字符串转换整数 (atoi)
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数。

函数 myAtoi(string s) 的算法如下：

空格：读入字符串并丢弃无用的前导空格（" "）
符号：检查下一个字符（假设还未到字符末尾）为 '-' 还是 '+'。如果两者都不存在，则假定结果为正。
转换：通过跳过前置零来读取该整数，直到遇到非数字字符或到达字符串的结尾。如果没有读取数字，则结果为0。
舍入：如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被舍入为 −231 ，大于 231 − 1 的整数应该被舍入为 231 − 1 。
返回整数作为最终结果。



示例 1：

输入：s = "42"

输出：42

解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。

带下划线线的字符是所读的内容，插入符号是当前读入位置。
第 1 步："42"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："42"（读入 "42"）
           ^
示例 2：

输入：s = " -042"

输出：-42

解释：

第 1 步："   -042"（读入前导空格，但忽视掉）
            ^
第 2 步："   -042"（读入 '-' 字符，所以结果应该是负数）
             ^
第 3 步："   -042"（读入 "042"，在结果中忽略前导零）
               ^
示例 3：

输入：s = "1337c0d3"

输出：1337

解释：

第 1 步："1337c0d3"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："1337c0d3"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："1337c0d3"（读入 "1337"；由于下一个字符不是一个数字，所以读入停止）
             ^
示例 4：

输入：s = "0-1"

输出：0

解释：

第 1 步："0-1" (当前没有读入字符，因为没有前导空格)
         ^
第 2 步："0-1" (当前没有读入字符，因为这里不存在 '-' 或者 '+')
         ^
第 3 步："0-1" (读入 "0"；由于下一个字符不是一个数字，所以读入停止)
          ^
示例 5：

输入：s = "words and 987"

输出：0

解释：

读取在第一个非数字字符“w”处停止。

提示：

0 <= s.length <= 200
s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成
*/

// 自动机
type State int

const (
	START State = iota
	SIGN
	NUMBER
	END
)

func main() {
	s := "0-1" // "-91283472332" // "-42"
	fmt.Println(myAtoi(s))
	fmt.Println(myAtoi1(s))
	fmt.Println(myAtoi2(s))
}

// 自动机解法
func myAtoi(s string) int {
	// 定义状态转移表
	stateMap := make(map[State][]State, 0)
	stateMap[START] = []State{START, SIGN, NUMBER, END}
	stateMap[SIGN] = []State{END, END, NUMBER, END}
	stateMap[NUMBER] = []State{END, END, NUMBER, END}
	stateMap[END] = []State{END, END, END, END}

	// 符号
	sign := 1 // 1 表示正数，-1 表示负数
	// 当前状态
	state := START
	// 结果
	result := 0
	for _, v := range s {
		state = stateMap[state][getState(v)]
		switch state {
		case SIGN:
			if v == '-' {
				sign = -1
			}
		case NUMBER:
			tmp := result * sign
			if tmp > math.MaxInt32/10 || (tmp == math.MaxInt32/10 && int(v-'0') > 7) {
				return math.MaxInt32
			}
			if tmp < math.MinInt32/10 || (tmp == math.MinInt32/10 && int(v-'0') > 8) {
				return math.MinInt32
			}
			result = result*10 + int(v-'0')
		case END:
			return result * sign
		}
	}
	return result * sign
}

func getState(c rune) State {
	switch c {
	case ' ':
		return START
	case '+', '-':
		return SIGN
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER
	default:
		return END
	}
}

// 模拟法
func myAtoi1(s string) int {
	// 去除前导空格
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}
	if len(s) == 0 {
		return 0
	}
	// 判断符号
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	// 转换数字
	result := 0
	for i := 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		tmp := result * sign
		if tmp > math.MaxInt32/10 || (tmp == math.MaxInt32/10 && int(s[i]-'0') > 7) {
			return math.MaxInt32
		}
		if tmp < math.MinInt32/10 || (tmp == math.MinInt32/10 && int(s[i]-'0') > 8) {
			return math.MinInt32
		}
		result = result*10 + int(s[i]-'0')
	}
	return result * sign
}

// AI实现
func myAtoi2(s string) int {
	// 1. 去除前导空格
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// 2. 判断符号
	sign := 1
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// 3. 转换数字
	result := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		if result > (1<<31-1)/10 || (result == (1<<31-1)/10 && digit > 7) {
			if sign == 1 {
				return (1 << 31) - 1
			} else {
				return -(1 << 31)
			}
		}
		result = result*10 + digit
		i++
	}

	return result * sign
}
