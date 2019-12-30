package top100

// 给定一个非空二叉树，返回其最大路径和。
//
// 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
// 示例 1:
//
// 输入: [1,2,3]
//
//       1
//      / \
//     2   3
//
// 输出: 6
// 示例 2:
//
// 输入: [-10,9,20,null,null,15,7]
//
//    -10
//    / \
//   9  20
//     /  \
//    15   7
//
// 输出: 42

var result124 int

func maxPathSum(root *TreeNode) int {
	result124 = ^int(^uint(0) >> 1)
	_maxPathSum(root)
	return result124
}

func _maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 增加与0比较的过程，可以直接省略自己，或者纯左值右值的比较，只有当左值或者右值大于0的时候
	// 才需要相加，小于0,就相当与放弃
	leftSum := max(_maxPathSum(root.Left), 0)
	rightSum := max(_maxPathSum(root.Right), 0)

	result124 = max(result124, leftSum+rightSum+root.Val)

	return root.Val + max(leftSum, rightSum)
}
