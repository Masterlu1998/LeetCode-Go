package top100

// 给定一个二叉树，返回它的中序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
// 输出: [1,3,2]

type stack94 []*TreeNode

func newStack94() *stack94 {
	st := stack94(make([]*TreeNode, 0))
	return &st
}

func (s *stack94) Push(i *TreeNode) {
	*s = append(*s, i)
}

func (s *stack94) Pop() *TreeNode {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack94) Empty() bool {
	return len(*s) == 0
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	st := newStack94()
	cur := root

	for cur != nil || !st.Empty() {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}

		cur = st.Pop()
		result = append(result, cur.Val)

		cur = cur.Right
	}
	return result
}
