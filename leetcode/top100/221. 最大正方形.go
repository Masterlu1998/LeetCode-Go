package top100

// 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
//
// 示例:
//
// 输入:
//
// 1 0 1 0 0
// 1 0 1 1 1
// 1 1 1 1 1
// 1 0 0 1 0
//
// 输出: 4

// 动态规划，当你周围的一圈人都是一个数的时候，说明就差你凑一个正方形了
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix))
	maxLength := 0

	for i := 0; i < len(dp); i++ {
		if dp[i] == nil {
			dp[i] = make([]int, len(matrix[0]))
		}
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxLength = 1
		}
	}

	for i := 0; i < len(dp[0]); i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			maxLength = 1
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
				}
			}
		}
	}

	return maxLength * maxLength
}
