package codinginterviews

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
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
// 返回：
//
// [3,9,20,15,7]
//
//
// 提示：
//
// 节点总数 <= 1000

type queue []*TreeNode

func NewQueue32_1() *queue {
	q := queue(make([]*TreeNode, 0))
	return &q
}

func (q *queue) Enqueue(i *TreeNode) {
	*q = append(*q, i)
}

func (q *queue) Dequeue() *TreeNode {
	result := (*q)[0]
	(*q) = (*q)[1:]
	return result
}

func (q *queue) Empty() bool {
	return len(*q) == 0
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := NewQueue32_1()
	result := make([]int, 0)
	q.Enqueue(root)
	for !q.Empty() {
		cur := q.Dequeue()
		result = append(result, cur.Val)
		if cur.Left != nil {
			q.Enqueue(cur.Left)
		}

		if cur.Right != nil {
			q.Enqueue(cur.Right)
		}
	}

	return result
}
