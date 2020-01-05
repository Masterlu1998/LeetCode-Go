package top100

// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
// 示例:
//
// 输入: [10,9,2,5,3,7,101,18]
// 输出: 4
// 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
// 说明:
//
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
// 你算法的时间复杂度应该为 O(n2) 。
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1

	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				temp = max(temp, dp[j])
			}
		}

		dp[i] = temp + 1
		result = max(result, dp[i])
	}

	return result
}
