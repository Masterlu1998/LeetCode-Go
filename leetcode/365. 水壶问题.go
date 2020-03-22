package leetcode

// 有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
//
// 如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
//
// 你允许：
//
// 装满任意一个水壶
// 清空任意一个水壶
// 从一个水壶向另外一个水壶倒水，直到装满或者倒空
// 示例 1: (From the famous "Die Hard" example)
//
// 输入: x = 3, y = 5, z = 4
// 输出: True
// 示例 2:
//
// 输入: x = 2, y = 6, z = 5
// 输出: False

// bfs
type queue []element

type element struct {
	x, y int
}

func newQueue() *queue {
	q := queue(make([]element, 0))
	return &q
}

func (q *queue) Enqueue(i element) {
	*q = append(*q, i)
}

func (q *queue) Get() element {
	return (*q)[0]
}

func (q *queue) Dequeue() {
	*q = (*q)[1:]
}

func (q *queue) Empty() bool {
	return len(*q) == 0
}

func canMeasureWater(x int, y int, z int) bool {
	q := newQueue()
	cache := make(map[element]bool)

	start := element{0, 0}
	q.Enqueue(start)

	for !q.Empty() {
		cur := q.Get()
		q.Dequeue()

		if cur.x+cur.y == z || cur.x == z || cur.y == z {
			return true
		}

		if cache[cur] {
			continue
		}

		cache[cur] = true

		q.Enqueue(element{x, cur.y})
		q.Enqueue(element{cur.x, y})
		q.Enqueue(element{cur.x, 0})
		q.Enqueue(element{0, cur.y})
		q.Enqueue(element{cur.x - min(cur.x, y-cur.y), cur.y + min(cur.x, y-cur.y)})
		q.Enqueue(element{cur.x + min(cur.y, x-cur.x), cur.y - min(cur.y, x-cur.x)})

	}

	return false
}

func min(left, right int) int {
	if left < right {
		return left
	}

	return right
}
