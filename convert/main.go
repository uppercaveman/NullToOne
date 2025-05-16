package main

import (
	"fmt"
	"strings"
)

/*
Z子形变换

将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下
P   A   H   N
A P L S I I G
Y   I   G
返回从左往右，逐行读取字符形成的新字符串
示例 1:
	输入：s = "PAYPALISHIRING", numRows = 3
	输出："PAHNAPLSIIGYIR"

示例 2:
	输入：s = "PAYPALISHIRING", numRows = 4
	输出："PINALSIGYAHRPI"
*/

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	result := convert(s, numRows)
	fmt.Println(result)
}

func convert(s string, numRows int) string {
	lenth := len(s)
	if numRows == 1 || numRows >= lenth {
		return s
	}
	rows := make([]string, numRows)
	curRow := 0
	goingDown := false

	for i := 0; i < lenth; i++ {
		rows[curRow] += string(s[i])
		if curRow == 0 {
			goingDown = true
		} else if curRow == numRows-1 {
			goingDown = false
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}
	return strings.Join(rows, "")
}
