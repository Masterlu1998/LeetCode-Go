package codinginterviews

// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//
//
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1:
//
// 输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
// 输出: 2

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	mid := len(nums) / 2

	for left <= right {
		index := findIndex(left, right, nums)
		if index > mid {
			right = index - 1
		} else if index < mid {
			left = index + 1

		} else {
			return nums[index]
		}
	}

	return 0
}

func findIndex(left, right int, nums []int) int {
	rand := (left + right) / 2
	nums[rand], nums[left] = nums[left], nums[rand]
	pivot := left
	index := pivot + 1
	for i := left; i <= right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}

	nums[index-1], nums[pivot] = nums[pivot], nums[index-1]
	return index - 1
}
