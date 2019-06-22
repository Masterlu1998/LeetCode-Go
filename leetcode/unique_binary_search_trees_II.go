package leetcode

// 不同的二叉搜索树 II
// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
// 示例:
// 输入: 3
// 输出:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// 解释:
// 以上的输出对应以下 5 种不同结构的二叉搜索树：
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

// 思路1：自下而上返回左子树与右子树不同可能性的根节点，然后将根节点组合排列。

var mask []bool

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	mask = make([]bool, n+1)
	return _generateTrees(1, n)
}

func _generateTrees(left, right int) []*TreeNode {
	res := []*TreeNode{nil}

	for i := left; i <= right; i++ {
		if !mask[i] {
			mask[i] = true
			leftBranch := _generateTrees(left, i-1)
			rightBranch := _generateTrees(i+1, right)
			for _, leftNode := range leftBranch {
				for _, rightNode := range rightBranch {
					root := &TreeNode{i, leftNode, rightNode}
					res = append(res, root)
				}
			}
			mask[i] = false
		}
	}

	if len(res) > 1 {
		return res[1:]
	}
	return res
}
