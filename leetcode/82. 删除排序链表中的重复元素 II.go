package leetcode

// 82. 删除排序链表中的重复元素 II

// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//
// 示例 1:
//
// 输入: 1->2->3->3->4->4->5
// 输出: 1->2->5
// 示例 2:
//
// 输入: 1->1->1->2->3
// 输出: 2->3

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路1：使用三个指针，分别指向当前节点，当前节点前一个节点，当前节点前前一个节点，当遇到重复的节点的时候
// 前前一个节点停下，再次遇到不相同的节点的时候，前前一个节点指向当前节点，当前节点作为前一个节点，下一个
// 节点作为当前节点。结尾需要做一次判断，开头需要新增两个节点来满足算法的需要。

func deleteDuplicates2Func1(head *ListNode) *ListNode {
	firstNew := new(ListNode)
	firstNew.Val = int(^uint(0) >> 1)
	secondNew := new(ListNode)
	secondNew.Val = int(^uint(0) >> 1)
	firstNew.Next = secondNew
	secondNew.Next = head

	start := firstNew
	end := secondNew
	cur := head
	flag := -1
	for cur != nil {
		if cur.Val == end.Val {
			flag = 1
			cur = cur.Next
			end = end.Next
		} else {
			if flag == 1 {
				flag = -1
				start.Next = cur
				end = cur
				cur = cur.Next
			} else {
				cur = cur.Next
				start = start.Next
				end = end.Next
			}
		}
	}
	if start.Next != end && start.Next.Val == end.Val {
		start.Next = cur
	}

	return secondNew.Next
}

// 思路2：快慢指针，快指针用来跳过所有相同的节点，慢节点用来遍历所有的节点，当快节点的下一个节点不等于慢
// 节点的时候，就把慢节点的下一个节点给快节点的下一个节点，然后让快指针指向快节点的下一个节点。

func deleteDuplicates2Func2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{int(^uint(0) >> 1), head}
	fast := head
	slow := newHead
	for fast != nil {

		for fast != nil && fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}

		if fast != slow.Next {
			slow.Next = fast.Next
			fast = fast.Next
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return newHead.Next
}
