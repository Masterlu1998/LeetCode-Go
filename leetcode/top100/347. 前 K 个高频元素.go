package top100

// 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
//
// 示例 1:
//
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:
//
// 输入: nums = [1], k = 1
// 输出: [1]
// 说明：
//
// 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
// 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。

type block struct {
	Num   int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	temp := make([]*block, 0)
	cache := make(map[int]*block)

	for i := 0; i < len(nums); i++ {
		if b, ok := cache[nums[i]]; ok {
			b.Count++
		} else {
			b := &block{nums[i], 1}
			cache[nums[i]] = b
		}
	}

	for _, b := range cache {
		temp = append(temp, b)
	}

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		buildMaxHeap(temp)
		result = append(result, temp[0].Num)
		temp = temp[1:]
	}

	return result
}

func buildMaxHeap(nums []*block) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify347(nums, i)
	}
}

func heapify347(nums []*block, node int) {
	left, right := node*2+1, node*2+2
	min := node

	if left < len(nums) && nums[left].Count > nums[min].Count {
		min = left
	}

	if right < len(nums) && nums[right].Count > nums[min].Count {
		min = right
	}

	if min != node {
		nums[min], nums[node] = nums[node], nums[min]
		heapify347(nums, min)
	}
}
