package codinginterviews

// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
// B是A的子结构， 即 A中有出现和B相同的结构和节点值。
//
// 例如:
// 给定的树 A:
//
//      3
//     / \
//    4   5
//   / \
//  1   2
// 给定的树 B：
//
//    4
//   /
//  1
// 返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
//
// 示例 1：
//
// 输入：A = [1,2,3], B = [3,1]
// 输出：false
// 示例 2：
//
// 输入：A = [3,4,5,1,2], B = [4,1]
// 输出：true
// 限制：
//
// 0 <= 节点个数 <= 10000

type stack26 []*TreeNode

func NewStack() *stack26 {
	st := stack26(make([]*TreeNode, 0))
	return &st
}

func (st *stack26) Push(i *TreeNode) {
	*st = append(*st, i)
}

func (st *stack26) Pop() *TreeNode {
	result := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return result
}

func (st *stack26) Empty() bool {
	return len(*st) == 0
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}

	cur := A
	st := NewStack()

	for cur != nil || !st.Empty() {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}

		cur = st.Pop()
		if cur.Val == B.Val && check(cur, B) {
			return true
		}
		cur = cur.Right

	}

	return false
}

func check(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}

	if A == nil && B != nil {
		return false
	}

	if A != nil && B == nil {
		return true
	}

	return A.Val == B.Val && check(A.Left, B.Left) && check(A.Right, B.Right)
}
