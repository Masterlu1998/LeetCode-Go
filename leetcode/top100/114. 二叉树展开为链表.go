package top100

// 给定一个二叉树，原地将它展开为链表。
//
// 例如，给定二叉树
//
//    1
//   / \
//  2   5
// / \   \
// 3   4   6
// 将其展开为：
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type stack114 []*TreeNode

func NewStack114() *stack114 {
	st := stack114(make([]*TreeNode, 0))
	return &st
}

func (s *stack114) Push(i *TreeNode) {
	*s = append(*s, i)
}

func (s *stack114) Pop() *TreeNode {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack114) Empty() bool {
	return len(*s) == 0
}

func flattenFunc1(root *TreeNode) {
	tail, cur := root, root
	st := NewStack114()

	for !st.Empty() || cur != nil {
		for cur != nil {
			tail = cur
			st.Push(cur)
			cur = cur.Left
		}

		cur = st.Pop()

		if cur.Right != nil {
			tail.Left = cur.Right
			cur.Right = nil
			cur = tail.Left
		} else {
			cur = nil
		}
	}

	cur = root
	for cur != nil {
		cur.Right = cur.Left
		cur.Left = nil
		cur = cur.Right
	}

}

func flatten(root *TreeNode) {
	cur := root
	st := NewStack114()

	for !st.Empty() || cur != nil {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}

		if !st.Empty() {
			cur = st.Pop()
			tmp := cur.Right
			cur.Right = cur.Left
			cur.Left = nil
			for cur.Right != nil {
				cur = cur.Right
			}
			cur.Right = tmp
			cur = tmp
		}

	}

}
