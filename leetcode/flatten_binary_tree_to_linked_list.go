package leetcode

// 二叉树展开为链表
// 给定一个二叉树，原地将它展开为链表。
// 例如，给定二叉树
//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// 将其展开为：
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路1：非递归遍历二叉树，先遍历左树，直到不再有左子树的节点，然后将该节点的右子树暂存，将左子树搬运到右子树上
// 然后遍历右子树，将最后一个没有右子树的节点作为下一个操作的节点。

func flattenFunc1(root *TreeNode) {
	node := root
	stack := make([]*TreeNode, 0)

	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) != 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp := node.Right
			node.Right = node.Left
			node.Left = nil
			for node.Right != nil {
				node = node.Right
			}
			node.Right = tmp
			node = tmp
		}
	}
}

// 思路2：递归遍历二叉树，设置全局变量为遍历的最后一个节点，每次在这个节点右边增加子树

var lastNode *TreeNode
func flattenFunc2(root *TreeNode) {
	if root == nil {
		return 
	}
	_flattenFunc2(root)
}

func _flattenFunc2(node *TreeNode) {
	if node == nil {
		return
	}

	if lastNode != nil {
		lastNode.Right = node
		lastNode.Left = nil
	}
	lastNode = node
	originRight := node.Right
	_flattenFunc2(node.Left)
	_flattenFunc2(originRight)
}