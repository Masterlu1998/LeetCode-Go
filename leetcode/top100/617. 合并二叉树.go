package top100

// 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//
// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
//
// 示例 1:
//
// 输入:
//	Tree 1                     Tree 2
//          1                         2
//         / \                       / \
//        3   2                     1   3
//       /                           \   \
//      5                             4   7
// 输出:
// 合并后的树:
//	     3
//	    / \
//	   4   5
//	  / \   \
//	 5   4   7
// 注意: 合并必须从两个树的根节点开始。

// 遍历两个树，入参要么全是nil，要么都有默认0值
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var _mergeTrees func(t1, t2 *TreeNode) *TreeNode

	_mergeTrees = func(t1, t2 *TreeNode) *TreeNode {
		if t1 == nil && t2 == nil {
			return nil
		}

		t1.Val += t2.Val

		if t1.Left == nil && t2.Left != nil {
			t1.Left = &TreeNode{}
		}

		if t1.Left != nil && t2.Left == nil {
			t2.Left = &TreeNode{}
		}

		_mergeTrees(t1.Left, t2.Left)

		if t1.Right == nil && t2.Right != nil {
			t1.Right = &TreeNode{}
		}

		if t1.Right != nil && t2.Right == nil {
			t2.Right = &TreeNode{}
		}
		_mergeTrees(t1.Right, t2.Right)

		return t1
	}

	if t1 == nil && t2 != nil {
		t1 = &TreeNode{}
	}

	if t1 != nil && t2 == nil {
		t2 = &TreeNode{}
	}

	return _mergeTrees(t1, t2)
}
