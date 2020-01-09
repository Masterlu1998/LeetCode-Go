package top100

import "sort"

// 给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
// 你找到的子数组应是最短的，请输出它的长度。
//
// 示例 1:
//
// 输入: [2, 6, 4, 8, 10, 9, 15]
// 输出: 5
// 解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
// 说明 :
//
// 输入的数组长度范围在 [1, 10,000]。
// 输入的数组可能包含重复元素 ，所以升序的意思是<=。

// 排序数组，然后找左边界和右边界
func findUnsortedSubarray(nums []int) int {
	origin := make([]int, len(nums))
	copy(origin, nums)
	sort.Ints(nums)

	start := 0

	for start < len(nums) {
		if nums[start] != origin[start] {
			break
		}
		start++
	}

	if start == len(nums) {
		return 0
	}

	end := len(nums) - 1
	for end >= 0 {
		if nums[end] != origin[end] {
			break
		}
		end--
	}

	return end - start + 1
}
