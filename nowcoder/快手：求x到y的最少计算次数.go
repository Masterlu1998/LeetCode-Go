package nowcoder

type queue []int

func NewQueue() *queue {
	q := queue(make([]int, 0))
	return &q
}

func (q *queue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *queue) Dequeue() int {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *queue) Empty() bool {
	return len(*q) == 0
}

func (q *queue) Length() int {
	return len(*q)
}

func handle_100(x, y int) int {
	res := 0
	q := NewQueue()
	q.Enqueue(x)
	var bfs func(x int)
	bfs = func(x int) {
		for !q.Empty() {
			length := q.Length()
			for i := 0; i < length; i++ {
				tp := q.Dequeue()
				if tp == y {
					return
				}
				q.Enqueue(tp + 1)
				q.Enqueue(tp - 1)
				q.Enqueue(tp * 2)
			}
			res++
		}
	}
	bfs(x)
	return res
}
