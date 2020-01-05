package top100

// 请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
// 输出: false
// 示例 2:
//
// 输入: 1->2->2->1
// 输出: true
// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

// 边遍历边反转

// 先找中点后反转
func isPalindromeFunc1(head *ListNode) bool {
	if head == nil {
		return true
	}
	preHead := &ListNode{Next: head}
	fast, slow := preHead, preHead

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	right, left := slow.Next, slow
	slow.Next = nil
	reverse(head)

	if fast == nil {
		left = left.Next
	}

	for right != nil && left != nil {
		if right.Val != left.Val {
			return false
		}

		right = right.Next
		left = left.Next
	}

	if right != nil || left != nil {
		return false
	}

	return true

}

func reverse(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}

	cur := reverse(root.Next)
	root.Next.Next = root
	root.Next = nil

	return cur
}

// 边遍历，边反转
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeFunc2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	pre := &ListNode{Next: head}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		temp := slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}

	head.Next = nil

	left, right := pre, slow

	if fast != nil {
		right = right.Next
	}

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}

		left = left.Next
		right = right.Next
	}

	if left != nil || right != nil {
		return false
	}

	return true
}
