package top100

// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
// 说明：解集不能包含重复的子集。
//
// 示例:
//
// 输入: nums = [1,2,3]
// 输出:
// [
//  [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

// 思路：递归，将前一组的解全部算入结果，然后将自己加入每一个解，又是一种新解
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}

	last := subsets(nums[1:])
	result := make([][]int, len(last))
	copy(result, last)

	for i := 0; i < len(last); i++ {
		tmp := make([]int, len(last[i]))
		copy(tmp, last[i])
		tmp = append(tmp, nums[0])
		result = append(result, tmp)
	}

	return result
}
