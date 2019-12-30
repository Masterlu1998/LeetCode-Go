package top100

// 给定一个未排序的整数数组，找出最长连续序列的长度。
//
// 要求算法的时间复杂度为 O(n)。
//
// 示例:
//
// 输入: [100, 4, 200, 1, 3, 2]
// 输出: 4
// 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

func longestConsecutive(nums []int) int {
	cache := make(map[int]bool)
	result := 0

	for _, val := range nums {
		cache[val] = true
	}

	for num, _ := range cache {
		length := 1

		if !cache[num-1] {
			cur := num
			for cache[cur+1] {
				cur += 1
				length += 1
			}

			result = max(result, length)
		}
	}
	return result
}
