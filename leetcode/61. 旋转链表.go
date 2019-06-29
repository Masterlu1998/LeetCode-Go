package leetcode

// 61. 旋转链表

// 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

// 示例 1:

// 输入: 1->2->3->4->5->NULL, k = 2
// 输出: 4->5->1->2->3->NULL
// 解释:
// 向右旋转 1 步: 5->1->2->3->4->NULL
// 向右旋转 2 步: 4->5->1->2->3->NULL
// 示例 2:

// 输入: 0->1->2->NULL, k = 4
// 输出: 2->0->1->NULL
// 解释:
// 向右旋转 1 步: 2->0->1->NULL
// 向右旋转 2 步: 1->2->0->NULL
// 向右旋转 3 步: 0->1->2->NULL
// 向右旋转 4 步: 2->0->1->NULL

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路1：向左旋转k等于向右旋转length-k，同时向左旋转k等于向左旋转 n*length+k 。基于以上推论，将向左旋转改为向
// 右旋转。

func rotateRightFunc1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := head
	count := 1
	for tail.Next != nil {
		tail = tail.Next
		count++
	}

	realK := count - k%count

	newHead := head
	for i := 0; i < realK; i++ {
		tmp := newHead.Next
		newHead.Next = nil
		tail.Next = newHead
		tail = newHead
		newHead = tmp
	}
	return newHead
}

// 思路2：成环。遍历到尾部的时候记住链表length，然后再遍历 length - k % length 处，断开链表即可。

func rotateRightFunc2(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    cur := head
    count := 1
    for {
        count++
        cur = cur.Next
        if cur.Next == nil {
            cur.Next = head
            for i := 0; i < count - k % count; i++ {
                cur = cur.Next
            }
            newHead := cur.Next
            cur.Next = nil
            return newHead
        }
    }
}