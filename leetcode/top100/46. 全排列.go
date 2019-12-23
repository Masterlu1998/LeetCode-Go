package top100

// 给定一个没有重复数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
// 输出:
// [
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
// ]

// 回溯
var result46 [][]int
var length46 int

func permute(nums []int) [][]int {
	result46 = make([][]int, 0)
	length46 = len(nums)
	_permute(nums, []int{})
	return result46
}

func _permute(nums []int, temp []int) {
	if len(temp) == length46 {
		result46 = append(result46, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		_permute(append(newNums[:i], newNums[i+1:]...), append(temp, nums[i]))
	}
}
