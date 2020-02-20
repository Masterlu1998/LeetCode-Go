package codinginterviews

import "sort"

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
//
//
//
// 参考以下这颗二叉搜索树：
//
//     5
//    / \
//   2   6
//  / \
// 1   3
// 示例 1：
//
// 输入: [1,6,3,2,5]
// 输出: false
// 示例 2：
//
// 输入: [1,3,2,6,5]
// 输出: true

// 当从后续遍历和中序遍历重组二叉数做
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}

	midorder := make([]int, len(postorder))
	copy(midorder, postorder)
	sort.Ints(midorder)
	return _verifyPostorder(postorder, midorder)
}

func _verifyPostorder(postorder, midorder []int) bool {
	if len(postorder) != len(midorder) {
		return false
	}

	if len(postorder) == 0 || len(midorder) == 0 {
		return true
	}

	rootVal := postorder[len(postorder)-1]

	index := -1

	for i, num := range midorder {
		if num == rootVal {
			index = i
			break
		}
	}

	if index == -1 {

		return false
	}

	return _verifyPostorder(postorder[index:len(postorder)-1], midorder[index+1:]) && _verifyPostorder(postorder[:index], midorder[:index])
}
