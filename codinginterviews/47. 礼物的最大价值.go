package codinginterviews

// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
//
//
//
// 示例 1:
//
// 输入:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// 输出: 12
// 解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

func maxValue(grid [][]int) int {
	width, height := len(grid[0]), len(grid)
	dp := make([][]int, height)

	sum := 0
	for i := 0; i < height; i++ {
		if len(dp[i]) == 0 {
			dp[i] = make([]int, width)
		}
		sum += grid[i][0]
		dp[i][0] = sum
	}

	sum = 0
	for i := 0; i < width; i++ {
		sum += grid[0][i]
		dp[0][i] = sum
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[height-1][width-1]
}
