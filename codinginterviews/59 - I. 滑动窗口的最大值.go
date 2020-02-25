package codinginterviews

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
//
// 示例:
//
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]
// 解释:
//
//  滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 提示：
//
// 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

type queue59 []int

func newQueue() *queue59 {
	q := queue59(make([]int, 0))
	return &q
}

func (q *queue59) GetHead() int {
	return (*q)[0]
}

func (q *queue59) GetTail() int {
	return (*q)[len(*q)-1]
}

func (q *queue59) PushTail(i int) {
	*q = append(*q, i)
}

func (q *queue59) PushHead(i int) {
	(*q) = append(*q, (*q)[len(*q)-1])
	copy((*q)[1:], (*q)[:len(*q)-1])
	(*q)[0] = i
}

func (q *queue59) PopHead() {
	(*q) = (*q)[1:]
}

func (q *queue59) PopTail() {
	(*q) = (*q)[:len(*q)-1]
}
func (q *queue59) Empty() bool {
	return len(*q) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	q := newQueue()
	result := make([]int, 0)

	for i, num := range nums {
		for !q.Empty() && num > q.GetTail() {
			q.PopTail()
		}

		q.PushTail(num)

		if i-(k-1) >= 0 {
			result = append(result, q.GetHead())
			if nums[i-(k-1)] == q.GetHead() {
				q.PopHead()
			}
		}
	}

	return result
}
