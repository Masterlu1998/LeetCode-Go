package leetcode

// 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。你不能只是单纯的改变节点内部的值，而是需要实际的
// 进行节点交换。
// 示例
// 给定 1->2->3->4, 你应该返回 2->1->4->3.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路1：维护三个指针，分别为当前结点c，前一个结点p，后一个结点n，然后对三个节点进行操作：
// 1：p的下一个结点指向n的下一个结点
// 2：c的下一个结点指向p
// 3：c改为n的下一个结点，p改为n，n改为c的一下个结点
// 4：重复上述步骤，直到不再有n
// 同时当n不再有下一个结点的时候将p的下一个节点指向n，并将c的下一个节点指向p直接返回结果

func swapPairsFunc1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	c := head.Next
	result := c

	for c.Next != nil {
		n := c.Next
		if n.Next != nil {
			p.Next = n.Next
			c.Next = p
			c = n.Next
			p = n
		} else {
			p.Next = n
			c.Next = p
			return result
		}
	}
	p.Next = nil
	c.Next = p
	return result
}

// 思路2：同样使用三个指针p，c，n，但是在开始操作之前增加一个指向原来头节点的空节点，并用p指向它，c指向头节点，
// n指向c节点的下一个操作，然后做如下操作：
// 1：将c节点的下一个节点指向n的下一个节点。
// 2：将n节点的下一个节点指向c节点。
// 3：将p节点的下一个节点指向n节点。
// 4：p节点改为c节点。
// 5：c节点改为c节点的下一个节点。
// 6：重复上述步骤直到c为空或者c的下一个为空时结束循环

func swapPairsFunc2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    newHead := new(ListNode)
    newHead.Next = head
    
    c := newHead.Next
    p := newHead
    for c != nil && c.Next != nil {
        n := c.Next
        c.Next = n.Next
        n.Next = c
        p.Next = n
        
        p = c
        c = c.Next
    }
    
    return newHead.Next
}