package top100

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
// 输出: true
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//      / \
//     3   6
// 输出: false
// 解释: 输入为: [5,1,4,null,null,3,6]。
//      根节点的值为 5 ，但是其右子节点值为 4 。

// 搜索二叉数，中序遍历一定是一个从小到大的有序数组
type stack98 []*TreeNode

func newStack98() *stack98 {
	st := stack98(make([]*TreeNode, 0))
	return &st
}

func (s *stack98) Push(i *TreeNode) {
	*s = append(*s, i)
}

func (s *stack98) Pop() *TreeNode {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack98) Empty() bool {
	return len(*s) == 0
}

func isValidBST(root *TreeNode) bool {
	last := ^int(^uint(0) >> 1)

	st := newStack98()
	cur := root

	for cur != nil || !st.Empty() {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}

		cur = st.Pop()
		if last >= cur.Val {
			return false
		}
		last = cur.Val
		cur = cur.Right
	}

	return true
}
