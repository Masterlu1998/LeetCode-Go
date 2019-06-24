package leetcode

// 86. 分隔链表

// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

// 你应当保留两个分区中每个节点的初始相对位置。

// 示例:

// 输入: head = 1->4->3->2->5->2, x = 3
// 输出: 1->2->2->4->3->5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路1：遍历链表，发现符合要求的节点进行删除和插入操作，这种方法特殊情况较多，当prev与index相同的时候需要
// 手动定义prev与index都指向cur

func partitionFunc1(head *ListNode, x int) *ListNode {
	origin := new(ListNode)
	origin.Next = head
	index := origin
	prev := origin
	cur := head
	for cur != nil {
		if cur.Val < x {
			if prev == index {
				prev = cur
				index = cur
				cur = cur.Next
				continue
			}
			tmp := cur.Next
			prev.Next = tmp
			cur.Next = index.Next
			index.Next = cur
			index = cur
			cur = tmp

		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return origin.Next
}

// 思路2：做两个新链表，一个储存符合要求的节点，一个储存不符合要求的节点，然后两者连起来。

 func partitionFunc2(head *ListNode, x int) *ListNode {
    biggerHead, smallerHead := new(ListNode), new(ListNode)
    bigger, smaller := biggerHead, smallerHead
    
    for cur := head; cur != nil; cur = cur.Next {
        if cur.Val < x {
            smaller.Next = &ListNode{cur.Val, nil}
            smaller = smaller.Next
        } else {
            bigger.Next = &ListNode{cur.Val, nil}
            bigger = bigger.Next
        }
    }
    
    smaller.Next = biggerHead.Next
    return smallerHead.Next
}