package leetcode

import (
	"sort"
)

// 组合总和 II
// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 说明：
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。 
// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 所求解集为:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// 示例 2:
// 输入: candidates = [2,5,2,1,2], target = 5,
// 所求解集为:
// [
//   [1,2,2],
//   [5]
// ]

var resultCom2 [][]int

func combinationSum2(candidates []int, target int) [][]int {
    resultCom2 = make([][]int, 0)
    sort.Ints(candidates)
    _combinationSum2(candidates, []int{}, target, 0)
    return resultCom2
}

func _combinationSum2(cadidates, temp []int, target, index int) {
    if sum(temp) > target {
        return
    }
    if sum(temp) == target {
        final := make([]int, len(temp))
        copy(final, temp)
        resultCom2 = append(resultCom2, final)
        return
    }
    
    for i := index; i < len(cadidates); i++ {
        if i > index && cadidates[i] == cadidates[i-1] {
            continue
        }
        _combinationSum2(cadidates, append(temp, cadidates[i]), target, i+1)
    }
}


