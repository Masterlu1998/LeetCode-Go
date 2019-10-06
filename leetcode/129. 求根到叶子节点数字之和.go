package leetcode

import "fmt"

// 给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
//
// 例如，从根到叶子节点路径 1->2->3 代表数字 123。
//
// 计算从根到叶子节点生成的所有数字之和。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例 1:
//
// 输入: [1,2,3]
//     1
//    / \
//   2   3
// 输出: 25
// 解释:
// 从根到叶子节点路径 1->2 代表数字 12.
// 从根到叶子节点路径 1->3 代表数字 13.
// 因此，数字总和 = 12 + 13 = 25.
// 示例 2:
//
// 输入: [4,9,0,5,1]
//     4
//    / \
//   9   0
//  / \
// 5   1
// 输出: 1026
// 解释:
// 从根到叶子节点路径 4->9->5 代表数字 495.
// 从根到叶子节点路径 4->9->1 代表数字 491.
// 从根到叶子节点路径 4->0 代表数字 40.
// 因此，数字总和 = 495 + 491 + 40 = 1026.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路1：用数组记录叶子节点的路径，到叶子节点的时候计算和并累加。
var s int

func sumNumbersFunc1(root *TreeNode) int {
	s = 0
	_sumNumbersFunc1(root, nil)
	return s
}

func _sumNumbersFunc1(node *TreeNode, path []*TreeNode) {
	if node == nil {
		return
	}

	path = append(path, node)
	if node.Left == nil && node.Right == nil {
		num := 0
		for _, n := range path {
			num = num*10 + n.Val
		}
		fmt.Println(num)
		s += num
		path = path[:len(path)-1]
		return
	}

	_sumNumbersFunc1(node.Left, path)

	_sumNumbersFunc1(node.Right, path)
}

// 思路二：将和作为参数传递下去，不再遍历数组
func sumNumbersFunc2(root *TreeNode) int {
	return _sumNumbersFunc2(root, 0)
}

func _sumNumbersFunc2(node *TreeNode, s int) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return s*10 + node.Val
	}

	return _sumNumbersFunc2(node.Left, s*10+node.Val) + _sumNumbersFunc2(node.Right, s*10+node.Val)
}
