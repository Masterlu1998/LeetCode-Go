package top100

// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回滑动窗口中的最大值。
//
//
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
//
//
//
// 进阶：
//
// 你能在线性时间复杂度内解决此题吗？

// 维护一个堆
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, 0)

	for i := 0; i+k-1 < len(nums); i++ {
		temp := make([]int, len(nums))
		copy(temp, nums)
		buildMaxHeap239(temp[i : i+k])
		result = append(result, temp[i])
	}

	return result
}

func buildMaxHeap239(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify239(nums, i)
	}
}

func heapify239(nums []int, node int) {
	left, right := 2*node+1, 2*node+2
	max := node
	if left < len(nums) && nums[left] > nums[max] {
		max = left
	}

	if right < len(nums) && nums[right] > nums[max] {
		max = right
	}

	if max != node {
		nums[max], nums[node] = nums[node], nums[max]
		heapify239(nums, max)
	}
}

// TODO：用双向链表做
