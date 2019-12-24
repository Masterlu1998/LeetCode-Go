package top100

// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
// 示例:
//
// 输入:
// [
//   [1,3,1],
//  [1,5,1],
//  [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。

// 思路：动态规划dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	cache := make([][]int, len(grid))

	tmp := 0
	for i := 0; i < len(grid); i++ {
		cache[i] = make([]int, len(grid[0]))
		tmp += grid[i][0]
		cache[i][0] = tmp
	}

	tmp = 0
	for j := 0; j < len(grid[0]); j++ {
		tmp += grid[0][j]
		cache[0][j] = tmp
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			cache[i][j] = min(cache[i-1][j], cache[i][j-1]) + grid[i][j]
		}
	}

	return cache[len(grid)-1][len(grid[0])-1]
}
