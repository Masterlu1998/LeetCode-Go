package codinginterviews

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
// 示例 1：
//
// 输入：arr = [3,2,1], k = 2
// 输出：[1,2] 或者 [2,1]
// 示例 2：
//
// 输入：arr = [0,1,2,1], k = 1
// 输出：[0]

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	heap := make([]int, k)

	for i := 0; i < k; i++ {
		heap[i] = arr[i]
	}
	buildMaxHeap(heap)

	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			heap[0] = arr[i]
			heapify(0, heap)
		}
	}

	return heap
}

func buildMaxHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(i, nums)
	}
}

func heapify(node int, nums []int) {
	largest := node
	left, right := 2*node+1, 2*node+2
	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}

	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}

	if largest != node {
		nums[largest], nums[node] = nums[node], nums[largest]
		heapify(largest, nums)
	}
}
