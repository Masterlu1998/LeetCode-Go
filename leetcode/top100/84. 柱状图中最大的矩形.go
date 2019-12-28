package top100

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// 示例:
//
// 输入: [2,1,5,6,2,3]
// 输出: 10

// 思路：二分法，找出最小数的索引，然后用此数的左边和此数的右边以及以该数为高的矩形哪个
// 面积大

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	if len(heights) == 1 {
		return heights[0]
	}

	minIndex := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] < heights[minIndex] {
			minIndex = i
		}
	}

	return max(max(largestRectangleArea(heights[:minIndex]), largestRectangleArea(heights[minIndex+1:])), heights[minIndex]*len(heights))
}
