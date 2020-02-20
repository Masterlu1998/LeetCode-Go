package codinginterviews

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
//
//
//
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 返回其层次遍历结果：
//
// [
//  [3],
//  [9,20],
//  [15,7]
// ]
//
//
// 提示：
//
// 节点总数 <= 1000

func levelOrder32_2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	temp := make([]*TreeNode, 0)
	temp = append(temp, root)
	result := make([][]int, 0)

	for len(temp) != 0 {
		level := make([]int, 0)
		mid := make([]*TreeNode, 0)
		for len(temp) != 0 {
			cur := temp[0]
			temp = temp[1:]

			level = append(level, cur.Val)

			if cur.Left != nil {
				mid = append(mid, cur.Left)
			}

			if cur.Right != nil {
				mid = append(mid, cur.Right)
			}

		}
		result = append(result, level)
		temp = mid
	}

	return result
}
