package top100

// 给定一个二叉树，检查它是否是镜像对称的。
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
// 3  4 4  3
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkNode(root, root)
}

func checkNode(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && checkNode(left.Left, right.Right) && checkNode(left.Right, right.Left)

}
