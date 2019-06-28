package leetcode

import (
	"sort"
)

// 47. 全排列 II

// 给定一个可包含重复数字的序列，返回所有不重复的全排列。

// 示例:

// 输入: [1,1,2]
// 输出:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

var resultPer2 [][]int
func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    resultPer2 = make([][]int, 0)
    _permuteUnique(nums, nums, []int{})
    return resultPer2
}

func _permuteUnique(nums, restnums, tmp []int) {
    if len(tmp) == len(nums) {
        input := make([]int, len(tmp))
        copy(input, tmp)
        resultPer2 = append(resultPer2, input)
        return
    }
    
    last := int(^uint(0) >> 1)
    for i := 0; i < len(restnums); i++ {
        if restnums[i] == last {
            continue
        }
        restcopy := make([]int, len(restnums))
        copy(restcopy, restnums)
        _permuteUnique(nums, append(restcopy[:i], restcopy[i+1:]...), append(tmp, restnums[i]))
        last = restnums[i]
    }
}