package codinginterviews

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的时间复杂度都是O(1)。
//
// 若队列为空，pop_front 和 max_value 需要返回 -1
//
// 示例 1：
//
// 输入:
// ["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
// [[],[1],[2],[],[],[]]
// 输出: [null,null,null,2,1,2]
// 示例 2：
//
// 输入:
// ["MaxQueue","pop_front","max_value"]
// [[],[],[]]
// 输出: [null,-1,-1]

type spQueue []int

func NewQueue() *spQueue {
	q := spQueue(make([]int, 0))
	return &q
}

func (q *spQueue) GetHead() int {
	return (*q)[0]
}

func (q *spQueue) GetTail() int {
	return (*q)[len(*q)-1]
}

func (q *spQueue) PushHead(i int) {
	(*q) = append(*q, (*q)[len(*q)-1])
	copy((*q)[1:], (*q)[:len(*q)-1])
	(*q)[0] = i
}

func (q *spQueue) PushTail(i int) {
	(*q) = append(*q, i)
}

func (q *spQueue) PopTail() {
	*q = (*q)[:len(*q)-1]
}

func (q *spQueue) PopHead() {
	*q = (*q)[1:]
}

func (q *spQueue) Empty() bool {
	return len(*q) == 0
}

type MaxQueue struct {
	queue1 *spQueue
	queue2 []int
}

func Constructor59_2() MaxQueue {
	return MaxQueue{NewQueue(), make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue2) == 0 {
		return -1
	}

	return this.queue1.GetHead()
}

func (this *MaxQueue) Push_back(value int) {
	this.queue2 = append(this.queue2, value)

	for !this.queue1.Empty() && value > this.queue1.GetTail() {
		this.queue1.PopTail()
	}

	this.queue1.PushTail(value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue2) == 0 {
		return -1
	}
	val := this.queue2[0]
	this.queue2 = this.queue2[1:]

	if this.queue1.GetHead() == val {
		this.queue1.PopHead()
	}

	return val
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor59_2();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
