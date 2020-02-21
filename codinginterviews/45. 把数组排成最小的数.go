package codinginterviews

import "strconv"

// 输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
//
//
//
// 示例 1:
//
// 输入: [10,2]
// 输出: "102"
// 示例 2:
//
// 输入: [3,30,34,5,9]
// 输出: "3033459"
//
//
// 提示:
//
// 0 < nums.length <= 100

func minNumber(nums []int) string {
	r := quickSort(0, len(nums)-1, nums)

	result := ""
	for _, num := range r {
		result += strconv.Itoa(num)
	}

	return result
}

func compare(left, right int) bool {
	leftStr := strconv.Itoa(left)
	rightStr := strconv.Itoa(right)
	lr := leftStr + rightStr
	rl := rightStr + leftStr
	return lr < rl
}

func quickSort(left, right int, nums []int) []int {
	if left < right {
		index := findIndex45(left, right, nums)
		quickSort(left, index-1, nums)
		quickSort(index+1, right, nums)
	}

	return nums
}

func findIndex45(left, right int, nums []int) int {
	choose := (left + right) / 2
	nums[left], nums[choose] = nums[choose], nums[left]
	pivot := left
	index := pivot + 1

	for i := left; i <= right; i++ {
		if compare(nums[i], nums[pivot]) {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}

	nums[index-1], nums[pivot] = nums[pivot], nums[index-1]
	return index - 1
}
