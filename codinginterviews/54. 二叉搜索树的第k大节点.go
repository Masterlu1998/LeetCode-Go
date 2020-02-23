package codinginterviews

// 给定一棵二叉搜索树，请找出其中第k大的节点。
//
//
//
// 示例 1:
//
// 输入: root = [3,1,4,null,2], k = 1
//   3
//  / \
// 1   4
//  \
//    2
// 输出: 4
// 示例 2:
//
// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//       5
//      / \
//     3   6
//    / \
//   2   4
//  /
// 1
// 输出: 4
//
//
// 限制：
//
// 1 ≤ k ≤ 二叉搜索树元素个数

type stack54 []*TreeNode

func NewStack54() *stack54 {
	st := stack54(make([]*TreeNode, 0))
	return &st
}

func (st *stack54) Push(i *TreeNode) {
	*st = append(*st, i)
}

func (st *stack54) Pop() *TreeNode {
	result := (*st)[len(*st)-1]
	(*st) = (*st)[:len(*st)-1]
	return result
}

func (st *stack54) Empty() bool {
	return len(*st) == 0
}

func kthLargest(root *TreeNode, k int) int {
	cur := root
	st := NewStack54()

	for cur != nil || !st.Empty() {
		for cur != nil {
			st.Push(cur)
			cur = cur.Right
		}

		cur = st.Pop()
		k--
		if k == 0 {
			return cur.Val
		}

		cur = cur.Left
	}

	return 0
}
