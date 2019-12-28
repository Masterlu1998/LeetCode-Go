package top100

// 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
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

type queue102 []*TreeNode

func newQueue102() *queue102 {
	q := queue102(make([]*TreeNode, 0))
	return &q
}

func (q *queue102) EnQueue(i *TreeNode) {
	*q = append(*q, i)
}

func (q *queue102) DeQueue() *TreeNode {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q *queue102) Empty() bool {
	return len(*q) == 0
}

func (q *queue102) Length() int {
	return len(*q)
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	tmp := make([]int, 0)
	q := newQueue102()
	cur := root
	q.EnQueue(cur)
	capacity := 1

	for !q.Empty() {
		cur = q.DeQueue()
		tmp = append(tmp, cur.Val)

		if cur.Left != nil {
			q.EnQueue(cur.Left)
		}

		if cur.Right != nil {
			q.EnQueue(cur.Right)
		}
		capacity--
		if capacity == 0 {
			result = append(result, tmp)
			tmp = make([]int, 0)
			capacity = q.Length()
		}
	}

	return result
}
