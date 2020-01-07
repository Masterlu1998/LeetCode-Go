package top100

// 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//
// 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//
// 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//
// 示例:
//
// 输入:
// [4,3,2,7,8,2,3,1]
//
// 输出:
// [5,6]

// 根据下标标记遍历过的数字为负数，剩下不是负数的数字，说明下标代表的数没有遍历过
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		number := nums[i]
		if number < 0 {
			number = -1 * number
		}

		if nums[number-1] > 0 {
			nums[number-1] = -nums[number-1]
		}
	}

	result := make([]int, 0)
	for index, val := range nums {
		if val > 0 {
			result = append(result, index+1)
		}
	}

	return result
}
