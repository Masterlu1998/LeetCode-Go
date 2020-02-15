package codinginterviews

// 找出数组中重复的数字。
//
//
// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//
// 示例 1：
//
// 输入：
// [2, 3, 1, 0, 2, 5, 3]
// 输出：2 或 3

// 众星归位
func findRepeatNumber(nums []int) int {
	for i, num := range nums {
		if i == num {
			continue
		}

		if num == nums[num] {
			return num
		}

		nums[num], nums[i] = nums[i], nums[num]
	}

	return -1
}
