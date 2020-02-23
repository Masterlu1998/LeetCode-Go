package codinginterviews

// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
//
//
//
// 示例 1:
//
// 输入: [7,5,6,4]
// 输出: 5
//
//
// 限制：
//
// 0 <= 数组长度 <= 50000

var count = 0

func reversePairs(nums []int) int {
	count = 0
	if len(nums) == 0 {
		return 0
	}
	mergeSort(nums)
	return count
}

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	length := len(nums)
	mid := length / 2

	left := nums[:mid]
	right := nums[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	for len(left) != 0 && len(right) != 0 {
		if left[0] > right[0] {
			count += len(right)

			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
