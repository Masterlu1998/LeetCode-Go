package queue

type queue []int

func newQueue() *queue {
	q := queue(make([]int, 0))
	return &q
}

func (q *queue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *queue) Dequeue() {
	*q = (*q)[1:]
}

func (q *queue) Get() int {
	return (*q)[0]
}

func (q *queue) Empty() bool {
	return len(*q) == 0
}

func (q *queue) Length() int {
	return len(*q)
}
