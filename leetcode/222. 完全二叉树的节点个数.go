package leetcode

import "math"

// 给出一个完全二叉树，求出该树的节点个数。
//
// 说明：
//
// 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。
//
// 示例:
//
// 输入:
//     1
//    / \
//   2   3
//  / \  /
// 4  5 6
//
// 输出: 6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路1：暴力遍历，不用完全二叉树的特性。
func countNodesFunc1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return count
}

// 思路2：如果左子树的高度等于右子树的高度，则可以直接用2**n-1公式
func countNodesFunc2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := 0, 0
	leftNode, rightNode := root, root
	for leftNode != nil {
		leftNode = leftNode.Left
		leftHeight++
	}

	for rightNode != nil {
		rightNode = rightNode.Right
		rightHeight++
	}

	if rightHeight == leftHeight {
		return intPow(2, rightHeight) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func intPow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
