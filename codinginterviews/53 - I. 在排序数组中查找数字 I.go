package codinginterviews

// 统计一个数字在排序数组中出现的次数。
//
//
//
// 示例 1:
//
// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: 2
// 示例 2:
//
// 输入: nums = [5,7,7,8,8,10], target = 6
// 输出: 0
//
//
// 限制：
//
// 0 <= 数组长度 <= 50000

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	index := -1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			index = mid
			break
		}
	}

	if index == -1 {
		return 0
	}

	result := 1
	// 向左找
	for i := index - 1; i >= 0; i-- {
		if nums[i] == target {
			result++
		}
	}

	// 向右找
	for i := index + 1; i < len(nums); i++ {
		if nums[i] == target {
			result++
		}
	}

	return result
}
