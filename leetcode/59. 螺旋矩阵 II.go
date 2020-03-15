package leetcode

// 给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
//
// 示例:
//
// 输入: 3
// 输出:
// [
// [ 1, 2, 3 ],
// [ 8, 9, 4 ],
// [ 7, 6, 5 ]
// ]

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	steprow, stepcol := n, n

	cur := 1
	row, col := 0, -1
	for steprow > 0 && stepcol > 0 {
		for i := 0; i < stepcol; i++ {
			col++
			matrix[row][col] = cur
			cur++
		}
		steprow--

		for i := 0; i < steprow; i++ {
			row++
			matrix[row][col] = cur
			cur++
		}
		stepcol--

		for i := 0; i < stepcol; i++ {
			col--
			matrix[row][col] = cur
			cur++
		}
		steprow--

		for i := 0; i < steprow; i++ {
			row--
			matrix[row][col] = cur
			cur++
		}
		stepcol--

	}

	return matrix
}
