package leetcode

// 494. 目标和

// 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//
// 示例 1:
//
// 输入: nums: [1, 1, 1, 1, 1], S: 3
// 输出: 5
// 解释:
//
// -1+1+1+1+1 = 3
// +1-1+1+1+1 = 3
// +1+1-1+1+1 = 3
// +1+1+1-1+1 = 3
// +1+1+1+1-1 = 3
//
// 一共有5种方法让最终目标和为3。
//
// 注意:
//
// 数组的长度不会超过20，并且数组中的值全为正数。
// 初始的数组的和不会超过1000。
// 保证返回的最终结果为32位整数。

var globalNums []int
var count int

// 思路1：暴力递归回溯
func findTargetSumWays(nums []int, S int) int {
	var tmp int
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
	}

	if tmp < S || -tmp > S || (tmp+S)%2 == 1 {
		// 数组和绝对值小于目标值或目标与数组和非偶数说明不可能达到目标值，直接返回0
		return 0
	}

	globalNums, count = nums, 0
	_findTargetSumWays(0, 0, S)
	return count
}

func _findTargetSumWays(nowOp, now, target int) {

	if nowOp == len(globalNums) {
		// 达到要求
		if now == target {
			count++
		}
		return
	}

	_findTargetSumWays(nowOp+1, now+globalNums[nowOp], target)
	_findTargetSumWays(nowOp+1, now-globalNums[nowOp], target)
}
