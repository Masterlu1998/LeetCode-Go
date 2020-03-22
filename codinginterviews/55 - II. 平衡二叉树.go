package codinginterviews

// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
//
//
//
// 示例 1:
//
// 给定二叉树 [3,9,20,null,null,15,7]
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 返回 true 。
//
// 示例 2:
//
// 给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
// 返回 false 。
//
//
//
// 限制：
//
// 1 <= 树的结点个数 <= 10000

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	tp := depth(root.Left) - depth(root.Right)
	if tp >= -1 && tp <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := depth(root.Left), depth(root.Right)
	return max(left, right) + 1
}
