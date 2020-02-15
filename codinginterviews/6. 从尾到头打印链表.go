package codinginterviews

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
// 示例 1：
//
// 输入：head = [1,3,2]
// 输出：[2,3,1]

type stack []*ListNode

func NewStack6() *stack {
	st := stack(make([]*ListNode, 0))
	return &st
}

func (s *stack) Push(i *ListNode) {
	*s = append(*s, i)
}

func (s *stack) Pop() *ListNode {
	result := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return result
}

func (s *stack) Empty() bool {
	return len(*s) == 0
}

func reversePrint(head *ListNode) []int {
	st := NewStack6()

	cur := head

	for cur != nil {
		st.Push(cur)
		cur = cur.Next
	}

	result := make([]int, 0)
	for !st.Empty() {
		e := st.Pop()
		result = append(result, e.Val)
	}

	return result
}
