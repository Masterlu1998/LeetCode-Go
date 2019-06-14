package leetcode

// 全排列
// 给定一个没有重复数字的序列，返回其所有可能的全排列。
// 示例
// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

// 思路1：回溯法，每次从剩下的数字中选取一个数字加入结果集中，然后将结果集与剩下的数字传入下一次递归，然后换一个
// 数字加入结果集，并修改剩下的数字，传入下一次递归。循环上述过程，直到结果集的长度与给出的数组长度相同。（自想）

var resultPer [][]int

func permute(nums []int) [][]int {
	resultPer = make([][]int, 0)
	_permute(nums, nums, []int{})
	return resultPer
}

func _permute(nums, restNums, tempResult []int) {
	if len(tempResult) == len(nums) {
		resultPer = append(resultPer, tempResult)
		return
	}

	for i := 0; i < len(restNums); i++ {
		nextTempResult := append(tempResult, restNums[i])
		// 语言问题，需要拷贝一份切片进行回溯操作
		nextRestNums := make([]int, len(restNums))
		copy(nextRestNums, restNums)
		nextRestNums = append(nextRestNums[:i], nextRestNums[i+1:]...)
		_permute(nums, nextRestNums, nextTempResult)
	}
}
