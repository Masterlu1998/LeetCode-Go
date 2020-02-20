package codinginterviews

import (
	"strconv"
	"strings"
)

// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
//
//
//
// 示例:
// 给定如下二叉树，以及目标和 sum = 22，
//
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
// 返回:
//
// [
//   [5,4,11,2],
//   [5,8,4,5]
// ]
//
//
// 提示：
//
// 节点总数 <= 10000

func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	var dfs func(node *TreeNode, s []int, temp int)
	dfs = func(node *TreeNode, s []int, temp int) {
		if node == nil {
			return
		}

		s = append(s, node.Val)
		temp += node.Val
		strings.Join(strconv.Itoa())
		if temp == sum && node.Left == nil && node.Right == nil {
			copys := make([]int, len(s))
			copy(copys, s)
			result = append(result, copys)
			return
		}

		dfs(node.Left, s, temp)
		dfs(node.Right, s, temp)
		s = s[:len(s)-1]
	}

	dfs(root, []int{}, 0)

	return result
}
