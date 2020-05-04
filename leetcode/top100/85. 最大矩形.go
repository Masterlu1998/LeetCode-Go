package top100

// 给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
// 示例:
//
// 输入:
// [
//  ["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]
// ]
// 输出: 6

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	cache := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		last := 0
		if cache[i] == nil {
			cache[i] = make([]int, len(matrix[i]))
		}
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				last++
			} else {
				last = 0
			}

			cache[i][j] = last
		}
	}

	maxArea := 0
	for i := 0; i < len(cache[0]); i++ {
		for k := len(cache) - 1; k >= 0; k-- {
			last := int(^uint(0) >> 1)
			for j := k; j >= 0; j-- {
				last = min(cache[j][i], last)
				temp := last * (k + 1 - j)
				maxArea = max(maxArea, temp)
			}
		}

	}
	return maxArea
}
