package top100

// 根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
// 你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var preIndex int

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	preIndex = 0

	return _buildTree(preorder, inorder)
}

func _buildTree(preorder, inorder []int) *TreeNode {
	rootVal := preorder[preIndex]
	preIndex++

	root := &TreeNode{
		Val: rootVal,
	}

	rootIndex := -1
	for index, val := range inorder {
		if val == rootVal {
			rootIndex = index
		}
	}

	if rootIndex <= 0 {
		root.Left = nil
	} else {
		root.Left = _buildTree(preorder, inorder[:rootIndex])
	}

	if rootIndex >= len(inorder)-1 {
		root.Right = nil
	} else {
		root.Right = _buildTree(preorder, inorder[rootIndex+1:])
	}

	return root
}
