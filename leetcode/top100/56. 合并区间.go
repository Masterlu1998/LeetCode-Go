package top100

import "sort"

// 给出一个区间的集合，请合并所有重叠的区间。
//
// 示例 1:
//
// 输入: [[1,3],[2,6],[8,10],[15,18]]
// 输出: [[1,6],[8,10],[15,18]]
// 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2:
//
// 输入: [[1,4],[4,5]]
// 输出: [[1,5]]
// 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

// 思路：先排序，将所有的区间连通，然后利用栈进行合并
type stack56 [][]int

func newStack56() *stack56 {
	s := stack56(make([][]int, 0))
	return &s
}

func (s *stack56) Push(i []int) {
	*s = append(*s, i)
}

func (s *stack56) Pop() []int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *stack56) Get() []int {
	return (*s)[len(*s)-1]
}

func (s *stack56) GetLength() int {
	return len(*s)
}

func mergeFunc1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(left, right int) bool {
		return intervals[left][0] < intervals[right][0]
	})
	st := newStack56()

	for i := 0; i < len(intervals); i++ {
		newItem := intervals[i]
		if st.GetLength() > 0 {
			oldItem := st.Get()
			if newItem[1] < oldItem[0] || newItem[0] > oldItem[1] {
				st.Push(newItem)
			} else {
				newItem = []int{min(newItem[0], oldItem[0]), max(newItem[1], oldItem[1])}
			}
		}

		st.Push(newItem)
	}

	return [][]int(*st)
}
