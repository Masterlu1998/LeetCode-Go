package codinginterviews

// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
// 示例1：
//
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
// 限制：
//
// 0 <= 链表长度 <= 1000

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		cur = cur.Next
	}

	for l1 != nil {
		cur.Next = &ListNode{Val: l1.Val}
		l1 = l1.Next
		cur = cur.Next
	}

	for l2 != nil {
		cur.Next = &ListNode{Val: l2.Val}
		l2 = l2.Next
		cur = cur.Next
	}

	return head.Next
}
