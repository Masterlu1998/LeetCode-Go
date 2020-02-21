package codinginterviews

// 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
//
// 例如，
//
// [2,3,4] 的中位数是 3
//
// [2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
// 设计一个支持以下两种操作的数据结构：
//
// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。
// 示例 1：
//
// 输入：
// ["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
// [[],[1],[2],[],[3],[]]
// 输出：[null,null,null,1.50000,null,2.00000]
// 示例 2：
//
// 输入：
// ["MedianFinder","addNum","findMedian","addNum","findMedian"]
// [[],[2],[],[3],[]]
// 输出：[null,null,2.00000,null,2.50000]

type heap41 struct {
	data    []int
	compare func(left, right int) bool
}

func newHeap(c func(left, right int) bool) *heap41 {
	return &heap41{data: make([]int, 0), compare: c}
}

func (h *heap41) Add(i int) {
	h.data = append(h.data, i)
	h.buildHeap()
}

func (h *heap41) buildHeap() {
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *heap41) heapify(node int) {
	unique := node
	left, right := 2*node+1, 2*node+2

	if left < len(h.data) && h.compare(h.data[left], h.data[unique]) {
		unique = left
	}

	if right < len(h.data) && h.compare(h.data[right], h.data[unique]) {
		unique = right
	}

	if unique != node {
		h.data[unique], h.data[node] = h.data[node], h.data[unique]
		h.heapify(unique)
	}
}

func (h *heap41) GetTop() int {
	return h.data[0]
}

func (h *heap41) PopTop() {
	h.data = h.data[1:]
	h.buildHeap()
}

func (h *heap41) Empty() bool {
	return len(h.data) == 0
}

type MedianFinder struct {
	maxHeap, minHeap *heap41
	length           int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	minCompare := func(left, right int) bool {
		return left < right
	}

	maxCompare := func(left, right int) bool {
		return left > right
	}

	return MedianFinder{
		minHeap: newHeap(minCompare),
		maxHeap: newHeap(maxCompare),
		length:  0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.length&1 == 1 {
		// 奇数 -> minHeap

		if !this.maxHeap.Empty() && this.maxHeap.GetTop() > num {
			temp := this.maxHeap.GetTop()
			this.maxHeap.PopTop()
			this.maxHeap.Add(num)
			this.minHeap.Add(temp)
		} else {
			this.minHeap.Add(num)
		}
	} else {
		// 偶数 -> maxHeap
		if !this.minHeap.Empty() && this.minHeap.GetTop() < num {
			temp := this.minHeap.GetTop()
			this.minHeap.PopTop()
			this.minHeap.Add(num)
			this.maxHeap.Add(temp)
		} else {
			this.maxHeap.Add(num)
		}
	}
	this.length++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.length&1 == 1 {
		return float64(this.maxHeap.GetTop())
	}

	return float64(this.minHeap.GetTop()+this.maxHeap.GetTop()) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
