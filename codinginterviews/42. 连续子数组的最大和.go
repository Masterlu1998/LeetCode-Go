package codinginterviews

// 输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//
// 要求时间复杂度为O(n)。
//
//
//
// 示例1:
//
// 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

func maxSubArray(nums []int) int {
	result := ^int(^uint(0) >> 1)
	last := 0

	for i := 0; i < len(nums); i++ {
		last = max(last+nums[i], nums[i])
		result = max(last, result)
	}

	return result
}

func max(left, right int) int {
	if left > right {
		return left
	}

	return right
}
