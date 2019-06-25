package leetcode

// 19. 删除链表的倒数第N个节点

// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

// 示例：

// 给定一个链表: 1->2->3->4->5, 和 n = 2.

// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：

// 给定的 n 保证是有效的。

// 进阶：

// 你能尝试使用一趟扫描实现吗？

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路1：用一个数组存储所有的节点，相当于把链表变成了数组，实际上是作弊行为。

 func removeNthFromEndFunc1(head *ListNode, n int) *ListNode {
    list := make([]*ListNode, 1)
    list[0] = head
    currentNode := head
    for currentNode.Next != nil {
        currentNode = currentNode.Next
        list = append(list, currentNode)
    }
    
    length := len(list)
    
    if length == 1 {
        return nil
    }
    
    if n == 1 {
        list[length - 2].Next = nil
    } else if length == n {
        return list[1]
    } else {
        deleteIndex := len(list) - n
        list[deleteIndex-1].Next = list[deleteIndex+1]
    }
    return list[0]
}


// 思路2：快慢指针。快指针比满指针快n，当快指针为nil的时候，慢指针刚好在倒数第n+1个节点处，删除第n个节点。为了防止
// 特殊情况，这里在头节点之前又加了一个节点，防止需要删除的是头节点但是头结点之前没有节点的问题。

func removeNthFromEndFunc2(head *ListNode, n int) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	fast := newHead
	for i := 0; i <= n; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next

	}

	slow := newHead
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return newHead.Next
}
