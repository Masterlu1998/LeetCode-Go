package leetcode

import "fmt"

// K 个一组翻转链表

// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 示例 :
// 给定这个链表：1->2->3->4->5
// 当 k = 2 时，应当返回: 2->1->4->3->5
// 当 k = 3 时，应当返回: 3->2->1->4->5
// 说明 :
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路1：链表按k分开，并翻转每一个子链表，然后把链表连接在一起

func reverseKGroupFunc1(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	cur := head
	newHead := head
	for i := 1; i < k; i++ {
		newHead = newHead.Next
		if newHead == nil {
			return head
		}
	}

	for cur != nil {
		tail := next(cur, k)
		if tail == nil {
			break
		}
		nextHead := tail.Next
		partHead, partTail := reversePart(cur, k)
		fmt.Println(partHead.Val)
		if prev != nil {
			prev.Next = partHead
		}
		partTail.Next = nextHead

		cur = nextHead
		prev = partTail
	}
	return newHead
}

func reversePart(head *ListNode, k int) (*ListNode, *ListNode) {
	cur := head.Next
	prev := head
	head.Next = nil
	for i := 1; i < k; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev, head
}

func next(node *ListNode, k int) *ListNode {
	count := 0
	for ; node != nil; node = node.Next {
		count++
		if count == k {
			return node
		}
	}
	return nil
}

// 思路2：递归实现，需要判断剩下的链表长度能不能进行一次转换

func reverseKGroupFunc2(head *ListNode, k int) *ListNode {
	return reversePartFunc2(head, k)
}

func reversePartFunc2(node *ListNode, k int) *ListNode {
	if node == nil {
		return nil
	}
	if valid(node, k) {
		cur := node.Next
		prev := node
		prev.Next = nil
		tail := node
		for i := 1; i < k; i++ {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		tail.Next = reversePartFunc2(cur, k)
		return prev
	}
	return node
}

func valid(node *ListNode, k int) bool {
	for i := 1; i < k; i++ {
		node = node.Next
		if node == nil {
			return false
		}
	}
	return true
}
