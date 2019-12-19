package top100

// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：
//
// 给定的 n 保证是有效的。
//
// 进阶：
//
// 你能尝试使用一趟扫描实现吗？

// 思路：快慢指针，慢指针慢块指针n+1步，保证慢指针最终指在要删除的节点的前一个。同时为了处理
// 单节点的情况，在头节点前再添加一个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{
		Next: head,
	}

	fast, slow := preHead, preHead
	count := 0
	for fast != nil {
		fast = fast.Next
		count++
		if count > n+1 {
			slow = slow.Next
		}
	}

	slow.Next = slow.Next.Next
	return preHead.Next
}
