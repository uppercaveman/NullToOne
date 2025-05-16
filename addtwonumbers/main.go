package main

import (
	"fmt"
)

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807
示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/

func main() {
	l1 := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2}}}
	l2 := &ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 5}}}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10 // 进位
		// 创建新节点，存储当前位的和
		// 取余数作为当前位的值
		// 例如：sum = 12，当前位的值为 2，进位为 1
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}
