package top100

// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

// 维护一个最小堆即可，迭代之后，堆顶元素就是所需要的元素
func findKthLargest(nums []int, k int) int {
	h := make([]int, k)
	copy(h, nums)

	buildMinHead(h)

	for i := k; i < len(nums); i++ {
		if nums[i] > h[0] {
			h[0] = nums[i]
			buildMinHead(h)
		}
	}

	return h[0]
}

func buildMinHead(s []int) {
	for i := len(s) - 1; i >= 0; i-- {
		heapify(s, i)
	}
}

func heapify(s []int, node int) {
	left := 2*node + 1
	right := 2*node + 2
	min := node
	if left < len(s) && s[left] < s[min] {
		min = left
	}

	if right < len(s) && s[right] < s[min] {
		min = right
	}

	if min != node {
		s[node], s[min] = s[min], s[node]
		heapify(s, min)
	}
}
