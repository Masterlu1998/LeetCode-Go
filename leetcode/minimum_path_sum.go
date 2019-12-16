package leetcode

// 最小路径和
// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 示例:
// 输入:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。

// 思路1：动态规划，状态转移方程：f(x, y) = grid[y][x] + Min(f(x-1, y), f(x, y-1))，当x=0或y=0的时候
// 为边界，值为x=0或y=0时，到y或x的所有数之和。有递归和非递归实现。

var cacheMin map[[2]int]int

func minPathSumFunc1(grid [][]int) int {
	cacheMin = make(map[[2]int]int)
	return _minPathSumFunc1(grid, len(grid[0])-1, len(grid)-1)
}

func _minPathSumFunc1(grid [][]int, x, y int) int {
	if val, ok := cacheMin[[2]int{x, y}]; ok {
		return val
	}

	if x == 0 {
		result := 0
		for i := 0; i <= y; i++ {
			result += grid[i][0]
		}
		return result
	}

	if y == 0 {
		result := 0
		for i := 0; i <= x; i++ {
			result += grid[0][i]
		}
		return result
	}
	result := grid[y][x] + Min(_minPathSumFunc1(grid, x-1, y), _minPathSumFunc1(grid, x, y-1))
	cacheMin[[2]int{x, y}] = result
	return result
}

// 非递归实现

func minPathSumFunc2(grid [][]int) int {
	width := len(grid[0])
	height := len(grid)
	matrix := make([][]int, height)
	// 初始化边界值
	temp := 0
	for i := 0; i < width; i++ {
		temp += grid[0][i]
		matrix[0] = append(matrix[0], temp)
	}

	temp = grid[0][0]
	for j := 1; j < height; j++ {
		temp += grid[j][0]
		matrix[j] = make([]int, width)
		matrix[j][0] = temp
	}

	// 动态规划
	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			matrix[i][j] = grid[i][j] + Min(matrix[i-1][j], matrix[i][j-1])
		}
	}

	return matrix[height-1][width-1]
}
