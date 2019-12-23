package top100

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 示例:
//
// 输入: [-2,1,-3,4,-1,2,1,-5,4],
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

// 贪心算法，要么带上当前值，要么舍弃前面的和从当前值开始，保证每一步都是最大的
func maxSubArrayFunc1(nums []int) int {
	maxResult, curSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = max(nums[i], curSum+nums[i])
		maxResult = max(curSum, maxResult)
	}
	return maxResult
}

// 分治法，分三种情况，要不全在左边，要不全在右边，要不穿心而过，之后左边和右边又可以分成自己的左边和右边
// 分到1为止返回。穿心而过单独计算最大值。每一次选择其中最大的晋级。
func maxSubArray(nums []int) int {
	return leftOrRightSumFunc2(nums, 0, len(nums)-1)
}

func leftOrRightSumFunc2(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	pivot := (left + right) / 2

	// pivot必须归于左边，因为pivot在上一步会被舍取小数偏小，如果再减1,就有可能比left还要小导致异常
	leftSum := leftOrRightSumFunc2(nums, left, pivot)
	rightSum := leftOrRightSumFunc2(nums, pivot+1, right)
	midSum := crossSum(nums, left, right)

	return max(max(leftSum, rightSum), midSum)
}

func crossSum(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	pivot := (left + right) / 2

	leftMaxSum := ^int(^uint(0) >> 1)
	curSum := 0
	for i := pivot; i >= left; i-- {
		curSum += nums[i]
		leftMaxSum = max(curSum, leftMaxSum)
	}

	curSum = 0
	rightMaxSum := ^int(0)
	for i := pivot + 1; i <= right; i++ {
		curSum += nums[i]
		rightMaxSum = max(curSum, rightMaxSum)
	}

	return leftMaxSum + rightMaxSum
}

func max(left, right int) int {
	if left > right {
		return left
	}

	return right
}
