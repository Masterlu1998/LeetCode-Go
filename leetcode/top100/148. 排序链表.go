package top100

// 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
// 示例 1:
//
// 输入: 4->2->1->3
// 输出: 1->2->3->4
// 示例 2:
//
// 输入: -1->5->3->4->0
// 输出: -1->0->3->4->5

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}
	left, right := split(head)
	return merge(sortList(left), sortList(right))
}

func split(head *ListNode) (*ListNode, *ListNode) {

	fast, slow, pre := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}

	pre.Next = nil
	return head, slow
}

func merge(left, right *ListNode) *ListNode {
	preHead := &ListNode{}
	cur := preHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	for left != nil {
		cur.Next = left
		left = left.Next
		cur = cur.Next
	}

	for right != nil {
		cur.Next = right
		right = right.Next
		cur = cur.Next
	}

	return preHead.Next
}
