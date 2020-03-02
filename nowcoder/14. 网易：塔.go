package nowcoder

// 题目描述
// 小易有一些立方体，每个立方体的边长为1，他用这些立方体搭了一些塔。
// 现在小易定义：这些塔的不稳定值为它们之中最高的塔与最低的塔的高度差。
// 小易想让这些塔尽量稳定，所以他进行了如下操作：每次从某座塔上取下一块立方体，并把它放到另一座塔上。
// 注意，小易不会把立方体放到它原本的那座塔上，因为他认为这样毫无意义。
// 现在小易想要知道，他进行了不超过k次操作之后，不稳定值最小是多少。
// 输入描述:
// 第一行两个数n,k (1 <= n <= 100, 0 <= k <= 1000)表示塔的数量以及最多操作的次数。
// 第二行n个数，ai(1 <= ai <= 104)表示第i座塔的初始高度。
// 输出描述:
// 第一行两个数s, m，表示最小的不稳定值和操作次数(m <= k)
// 接下来m行，每行两个数x,y表示从第x座塔上取下一块立方体放到第y座塔上。
// 示例1
// 输入
// 复制
// 3 2
// 5 8 5
// 输出
// 复制
// 0 2
// 2 1
// 2 3

func handle14(a []*element, k int) (int, int, [][]int) {
	res := int(^uint(0) >> 1)
	options := 0
	actions := make([][]int, 0)

	maxHeap := newHeap(a, func(left, right int) bool { return left > right })
	maxHeap.buildHeap()
	copya := make([]*element, len(a))
	copy(copya, a)
	minHeap := newHeap(copya, func(left, right int) bool { return left < right })
	minHeap.buildHeap()
	tpActions := make([][]int, 0)
	for i := 0; i < k; i++ {
		tpActions = append(tpActions, []int{maxHeap.data[0].index, minHeap.data[0].index})
		maxHeap.data[0].val = maxHeap.data[0].val - 1
		maxHeap.heapify(0)
		minHeap.data[0].val = minHeap.data[0].val + 1
		minHeap.heapify(0)
		tp := maxHeap.get() - minHeap.get()
		if tp < res {
			res = tp
			options = i + 1
			actions = tpActions
		}

	}

	return res, options, actions
}

func min(left, right int) int {
	if left < right {
		return left
	}

	return right
}

type element struct {
	val, index int
}

type heap struct {
	data    []*element
	compare func(left, right int) bool
}

func newHeap(data []*element, compare func(left, right int) bool) *heap {
	return &heap{data, compare}
}

func (h *heap) get() int {
	return h.data[0].val
}

func (h *heap) buildHeap() {
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *heap) heapify(cur int) {
	left, right := cur*2+1, cur*2+2
	unique := cur

	if left < len(h.data) && h.compare(h.data[left].val, h.data[unique].val) {
		unique = left
	}

	if right < len(h.data) && h.compare(h.data[right].val, h.data[unique].val) {
		unique = right
	}

	if unique != cur {
		h.data[unique], h.data[cur] = h.data[cur], h.data[unique]
		h.heapify(unique)
	}
}
