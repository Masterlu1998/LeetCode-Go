package nowcoder

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func convertToTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	index := -1
	for i, num := range inorder {
		if num == rootVal {
			index = i
			break
		}
	}

	root.Left = convertToTree(preorder[1:index+1], inorder[:index])
	root.Right = convertToTree(preorder[index+1:], inorder[index+1:])
	return root
}

func convertToSumTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	temp := node.Val
	node.Val = convertToSumTree(node.Left) + convertToSumTree(node.Right)
	return node.Val + temp
}

func view(node *TreeNode) {
	if node == nil {
		return
	}
	view(node.Left)
	fmt.Print(node.Val, " ")
	view(node.Right)
}
