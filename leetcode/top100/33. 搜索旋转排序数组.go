package top100

// 思路： 二分法解决问题
// 1. target > mid
//    一般来说都应该往右边找；除非target大于右边的全部数字，且此时mid在反转点右侧
// 2. target <= mid
//    一般来说应该往左边找；除非target小于右边的所有数字，且此时mid在反转点左侧

func searchFunc1(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			if target >= nums[left] && nums[mid] < nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] > target {
			if target <= nums[right] && nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			return mid
		}
	}

	return -1
}
