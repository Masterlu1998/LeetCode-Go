package leetcode

import (
	"fmt"
)

// 组合总数
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 说明：
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。 
// 示例
// 输入: candidates = [2,3,6,7], target = 7,
// 所求解集为:
// [
//   [7],
//   [2,2,3]
// ]


// 思路1：回溯法，遍历所有的可能性，当结果超出给定目标的时候结束遍历。需要注意数组在函数间传递的拷贝问题，防止已经在
// 结果中的答案被修改值。自想，效率很低，查重部分做得不好。

var resultCom [][]int

func combinationSumFunc1(candidates []int, target int) [][]int {
    resultCom = make([][]int, 0)
    _combinationSumFunc1(candidates, []int{}, target)
    return resultCom
}

func _combinationSumFunc1(nums, currentNums []int, target int) {

    if sum(currentNums) == target {

        for _, val := range resultCom {
            if isSame(val, currentNums) {
                return
            }
        }
        resultCom = append(resultCom, currentNums)
        return
    } else if sum(currentNums) > target {
        return
    }
    
    for i := 0; i < len(nums); i++ {
        // currentNums = append(currentNums, nums[i])
        nextCurrentNums := make([]int, len(currentNums)+1)
        // copy(nextCurrentNums, currentNums)
        copy(nextCurrentNums, append(currentNums, nums[i]))
		// currentNums = currentNums[:len(currentNums)-1]
		fmt.Println(nextCurrentNums)
        _combinationSumFunc1(nums, nextCurrentNums, target)
    }
}

func isSame(nums1, nums2 []int) bool {
    cache1 := make(map[int]int)
    cache2 := make(map[int]int)
    
    for _, val := range nums1 {
        cache1[val]++
    } 
    
    for _, val := range nums2 {
        cache2[val]++
    }
    
    if len(cache1) != len(cache2) {
        return false
    }
    
    for key, val := range cache1 {
        if cache2[key] != val {
            return false
        }
    }
    return true
}

// 思路2：优化回溯方法，每次回溯从当前索引开始防止重复的情况发生，除去查重所带来的额外开销。

var resultCan [][]int

func combinationSumFunc2(candidates []int, target int) [][]int {
    resultCan = make([][]int, 0)
    _combinationSumFunc2(candidates, []int{}, target, 0)
    return resultCan
}

func _combinationSumFunc2(candidates, currentNums []int, target, start int) {  
    for i := start; i < len(candidates); i++ {
        if candidates[i] > target {
            continue
        }
        
        if candidates[i] == target {
            finalCurrentNums := make([]int, len(currentNums)+1)
            copy(finalCurrentNums, append(currentNums, candidates[i]))
            
            resultCan = append(resultCan, finalCurrentNums)
        }
        
        if candidates[i] < target {
            _combinationSumFunc2(candidates, append(currentNums, candidates[i]), target-candidates[i], i)
        }
    }
}