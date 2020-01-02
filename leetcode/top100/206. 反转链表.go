package top100

// 反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

// 迭代
func reverseListFuc1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{Next: head}
	pre, cur, next := preHead, head, head.Next

	for next != nil {
		cur.Next = pre
		pre = cur
		cur = next
		next = next.Next
	}
	preHead.Next.Next = nil
	cur.Next = pre

	return cur
}

// 递归
func reverseListFunc2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseListFunc2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return cur
}
