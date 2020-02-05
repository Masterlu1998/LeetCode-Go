package top100

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。
//
// 说明：
//
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:
//
// 输入: candidates = [2,3,6,7], target = 7,
// 所求解集为:
// [
//  [7],
//  [2,2,3]
// ]
// 示例 2:
//
// 输入: candidates = [2,3,5], target = 8,
// 所求解集为:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

// 递归+枝剪，每次从当前位置开始回溯能够避免重复
var result39 [][]int
var t int

func combinationSumFunc1(candidates []int, target int) [][]int {
	result39 = make([][]int, 0)
	t = target
	_combinationSum(candidates, []int{}, 0)
	return result39
}

func _combinationSum(candidates []int, temp []int, sum int) {
	if sum > t {
		return
	}

	if sum == t {
		clone := make([]int, len(temp))
		copy(clone, temp)
		result39 = append(result39, clone)
		return
	}

	for i, val := range candidates {
		// 从当前位置开始，避免回溯
		_combinationSum(candidates[i:], append(temp, val), sum+val)
	}
}
