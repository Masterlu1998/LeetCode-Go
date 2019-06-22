package leetcode

// 恢复二叉搜索树
// 二叉搜索树中的两个节点被错误地交换。
// 请在不改变其结构的情况下，恢复这棵树。
// 示例 1:
// 输入: [1,3,null,null,2]
//    1
//   /
//  3
//   \
//    2
// 输出: [3,1,null,null,2]
//    3
//   /
//  1
//   \
//    2
// 示例 2:
// 输入: [3,1,4,null,null,2]
//   3
//  / \
// 1   4
//    /
//   2
// 输出: [2,1,4,null,null,3]
//   2
//  / \
// 1   4
//    /
//   3

// 思路1：中序遍历搜索二叉树，并用数组记录节点。然后将数组根据节点值的大小进行排序，最后遍历排序前和排序后的数组，
// 找到错误的索引并进行交换即可。

func recoverTreeFunc1(root *TreeNode) {
	result := make([]*TreeNode, 0)
	node := root
	stack := make([]*TreeNode, 0)

	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node)

		node = node.Right
	}
	origin := make([]*TreeNode, len(result))
	copy(origin, result)
	sortNode(result)
	wrongIndex := -1
	for i := 0; i < len(result); i++ {
		if origin[i].Val != result[i].Val {
			if wrongIndex == -1 {
				wrongIndex = i
			} else {
				result[i].Val, result[wrongIndex].Val = result[wrongIndex].Val, result[i].Val
				break
			}
		}
	}

}

func sortNode(list []*TreeNode) {
	_sortNode(list, 0, len(list)-1)
}

func _sortNode(list []*TreeNode, left, right int) []*TreeNode {
	if left < right {
		index := findIndex(list, left, right)
		_sortNode(list, left, index-1)
		_sortNode(list, index+1, right)
	}
	return list
}

func findIndex(list []*TreeNode, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := left; i <= right; i++ {
		if list[i].Val < list[pivot].Val {
			list[i], list[index] = list[index], list[i]
			index++
		}
	}

	list[index-1], list[pivot] = list[pivot], list[index-1]
	return index - 1
}

// 思路2：中序遍历二叉树，找到两个有问题的节点并进行交换。需要注意的是，需要遍历整个二叉树，并通过记住逆序的开始
// 和结尾找到最终需要调换的两个节点，而并不是出现两个不符合要求的节点就进行交换，也不是局部节点满足升序就可以的。

func recoverTreeFunc2(root *TreeNode) {
	node := root
	stack := make([]*TreeNode, 0)
	var a, b *TreeNode

	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if b == nil {
			if a == nil || node.Val > a.Val {
				a = node
			} else {
				b = node
			}
		} else {
			if node.Val < b.Val {
				b = node
			}
		}
		node = node.Right
	}
	if a != nil && b != nil {
		a.Val, b.Val = b.Val, a.Val
	}
}
