package codinginterviews

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
//  [20,9],
//  [15,7]
// ]
//
//
// 提示：
//
// 节点总数 <= 1000

func levelOrder32_3(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	lastLevel := make([]*TreeNode, 0)
	lastLevel = append(lastLevel, root)
	result := make([][]int, 0)

	for len(lastLevel) != 0 {
		temp := make([]*TreeNode, 0)
		level := make([]int, 0)

		for i := 0; i < len(lastLevel); i++ {
			var cur *TreeNode
			if len(result)&1 == 1 {
				// 从右往左
				cur = lastLevel[len(lastLevel)-1-i]
			} else {
				// 从左往右
				cur = lastLevel[i]
			}

			level = append(level, cur.Val)

			if lastLevel[i].Left != nil {
				temp = append(temp, lastLevel[i].Left)
			}

			if lastLevel[i].Right != nil {
				temp = append(temp, lastLevel[i].Right)
			}
		}

		lastLevel = temp
		result = append(result, level)
	}

	return result
}
