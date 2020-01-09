package top100

// 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
//
// 例如：
//
// 输入: 二叉搜索树:
//              5
//            /   \
//           2     13
//
// 输出: 转换为累加树:
//             18
//            /   \
//          20     13

// 中序倒着遍历一遍，用当前累计和就可以了
type stack538 []*TreeNode

func newStack538() *stack538 {
	st := stack538(make([]*TreeNode, 0))
	return &st
}

func (s *stack538) Push(i *TreeNode) {
	*s = append((*s), i)
}

func (s *stack538) Pop() *TreeNode {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack538) Empty() bool {
	return len(*s) == 0
}

func convertBST(root *TreeNode) *TreeNode {
	cur := root
	st := newStack538()
	temp := 0

	for cur != nil || !st.Empty() {
		for cur != nil {
			st.Push(cur)
			cur = cur.Right
		}

		cur = st.Pop()
		temp += cur.Val
		cur.Val = temp
		cur = cur.Left
	}

	return root
}
