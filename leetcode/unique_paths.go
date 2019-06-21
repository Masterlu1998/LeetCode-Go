package leetcode


// 不同路径
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 问总共有多少条不同的路径？
// 说明：m 和 n 的值均不超过 100。
// 示例 1:
// 输入: m = 3, n = 2
// 输出: 3
// 解释:
// 从左上角开始，总共有 3 条路径可以到达右下角。
// 1. 向右 -> 向右 -> 向下
// 2. 向右 -> 向下 -> 向右
// 3. 向下 -> 向右 -> 向右


// 思路1：暴力回溯法
var resultUni int

func uniquePathsFunc1(m int, n int) int {
    resultUni = 0
    _uniquePathsFunc1(m, n, 1, 1)
    return resultUni
}

func _uniquePathsFunc1(m, n, nowM, nowN int) {
    if nowM == m && nowN == n {
        resultUni++
        return
    }
    if nowM < m {
        _uniquePathsFunc1(m, n, nowM+1, nowN)
    }
    
    if nowN < n {
        _uniquePathsFunc1(m, n, nowM, nowN+1)
    }
}

// 思路2：动态规划，状态转移方程f(m, n) = f(m-1, n) + f(m, n-1)，当 m = 0 或 n = 0 的时候为边界答案为1
// 以下是递归实现和非递归实现

var cache map[[2]int]int

func uniquePathsFunc2(m int, n int) int {
    cache = make(map[[2]int]int)
    return _uniquePathsFunc2(m, n)
}

func _uniquePathsFunc2(m, n int) int {
    if m == 1 || n == 1 {
        return 1
    }
    if val, ok := cache[[2]int{m, n}]; ok {
        return val
    }
    result := _uniquePathsFunc2(m-1, n) + _uniquePathsFunc2(m, n-1)
    cache[[2]int{m, n}] = result
    return result
}

// 思路2：动态规划非递归实现

func uniquePathsFunc3(m int, n int) int {
    var matrix [100][100]int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 || j == 0 {
                matrix[i][j] = 1
            } else {
                matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
            }
        }
    }
    return matrix[m-1][n-1]
}
