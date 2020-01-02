package top100

// 翻转一棵二叉树。
//
// 示例：
//
// 输入：
//
//     4
//   /   \
//  2     7
// / \   / \
// 1   3 6   9
// 输出：
//
//     4
//   /   \
//  7     2
// / \   / \
// 9   6 3   1

// 迭代
type stack226 []*TreeNode

func newStack226() *stack226 {
	st := stack226(make([]*TreeNode, 0))
	return &st
}

func (s *stack226) Push(i *TreeNode) {
	*s = append(*s, i)
}

func (s *stack226) Pop() *TreeNode {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack226) Empty() bool {
	return len(*s) == 0
}

func invertTreeFunc1(root *TreeNode) *TreeNode {
	cur := root
	st := newStack226()
	out := newStack226()

	for cur != nil || !st.Empty() {
		for cur != nil {
			st.Push(cur)
			out.Push(cur)
			cur = cur.Left
		}

		cur = st.Pop()
		cur = cur.Right
	}

	for !out.Empty() {
		node := out.Pop()
		node.Left, node.Right = node.Right, node.Left
	}

	return root
}

// 递归
func invertTreeFunc2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeFunc1(root.Left)
	invertTreeFunc1(root.Right)
	return root
}
