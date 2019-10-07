package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路1：利用快慢指针找到链表需要反转的部分的头和尾，
// 然后进行链表反转，并规定头尾的指向
func reverseBetweenFunc1(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	if m == n {
		return head
	}

	newHead := &ListNode{0, head}
	fast := newHead
	slow := newHead
	length := n - m + 1

	for i := 0; i < n; i++ {
		if length <= 0 {
			slow = slow.Next
		}
		fast = fast.Next
		length--
	}

	reverseHead := slow.Next
	prev := reverseHead
	cur := reverseHead.Next
	next := cur.Next
	reverseHead.Next = fast.Next
	for cur != fast {
		cur.Next = prev
		prev = cur
		cur = next
		next = cur.Next
	}

	cur.Next = prev
	slow.Next = fast

	return newHead.Next
}
