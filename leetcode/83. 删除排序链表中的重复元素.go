package leetcode

// 83. 删除排序链表中的重复元素

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

// 示例 1:

// 输入: 1->1->2
// 输出: 1->2
// 示例 2:

// 输入: 1->1->2->3->3
// 输出: 1->2->3

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    tmp := head
    cur := head
    for cur != nil && cur.Next != nil {
        if cur.Next.Val == cur.Val {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return tmp
}