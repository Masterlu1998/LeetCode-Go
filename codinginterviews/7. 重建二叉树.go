package codinginterviews

// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
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
// 限制：
//
// 0 <= 节点个数 <= 5000

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}

	index := -1
	for i, num := range inorder {
		if num == rootVal {
			index = i
		}
	}

	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}
