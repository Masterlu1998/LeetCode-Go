package top100

import "algorithm-go/leetcode"

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

// 思路1：依次循环处理两个指针即可；需要注意几点：1. 考虑进位。2. 考虑链表长短不一。
func addTwoNumbers(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {
	lastNode := &leetcode.ListNode{}
	head := lastNode
	node1, node2 := l1, l2
	lastAddNum := 0

	for node1 != nil && node2 != nil {
		currentAddNum := 0
		val := node1.Val + node2.Val + lastAddNum
		if val >= 10 {
			currentAddNum = val / 10
			val = val % 10
		}
		newNode := &leetcode.ListNode{
			Val: val,
		}

		lastAddNum = currentAddNum
		lastNode.Next = newNode
		lastNode = lastNode.Next
		node1 = node1.Next
		node2 = node2.Next
	}

	for node1 != nil {
		currentAddNum := 0
		val := node1.Val + lastAddNum
		if val >= 10 {
			currentAddNum = val / 10
			val = val % 10
		}
		newNode := &leetcode.ListNode{
			Val: val,
		}
		lastAddNum = currentAddNum
		lastNode.Next = newNode
		lastNode = lastNode.Next
		node1 = node1.Next
	}

	for node2 != nil {
		currentAddNum := 0
		val := node2.Val + lastAddNum
		if val >= 10 {
			currentAddNum = val / 10
			val = val % 10
		}

		newNode := &leetcode.ListNode{
			Val: val,
		}

		lastAddNum = currentAddNum
		lastNode.Next = newNode
		lastNode = lastNode.Next
		node2 = node2.Next
	}

	if lastAddNum != 0 {
		newNode := &leetcode.ListNode{
			Val: lastAddNum,
		}
		lastNode.Next = newNode
	}

	return head.Next
}
