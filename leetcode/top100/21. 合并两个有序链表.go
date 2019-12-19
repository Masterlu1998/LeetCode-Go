package top100

// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例：
//
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

// 思路：没啥好说的，该干嘛干嘛
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	cur := preHead

	node1, node2 := l1, l2

	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			cur.Next = node1
			node1 = node1.Next
		} else if node1.Val > node2.Val {
			cur.Next = node2
			node2 = node2.Next
		}

		cur = cur.Next
	}

	for node1 != nil {
		cur.Next = node1
		node1 = node1.Next
		cur = cur.Next
	}

	for node2 != nil {
		cur.Next = node2
		node2 = node2.Next
		cur = cur.Next
	}

	return preHead.Next
}
