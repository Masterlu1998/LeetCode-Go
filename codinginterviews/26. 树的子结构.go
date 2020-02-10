package codinginterviews

// 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。
// s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

type stack26 []*TreeNode

func NewStack() *stack26 {
	st := stack26(make([]*TreeNode, 0))
	return &st
}

func (st *stack26) Push(i *TreeNode) {
	*st = append(*st, i)
}

func (st *stack26) Pop() *TreeNode {
	result := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return result
}

func (st *stack26) Empty() bool {
	return len(*st) == 0
}
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}

	node := s
	st := NewStack()

	for node != nil || !st.Empty() {
		for node != nil {
			st.Push(node)
			node = node.Left
		}

		node = st.Pop()

		if node.Val == t.Val {
			isSame := compare(node, t)
			if isSame {
				return true
			}
		}
		node = node.Right
	}

	return false
}

func compare(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val {
		return compare(s.Left, t.Left) && compare(s.Right, t.Right)
	}

	return false
}
