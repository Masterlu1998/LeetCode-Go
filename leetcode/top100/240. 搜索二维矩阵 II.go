package top100

// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// 示例:
//
// 现有矩阵 matrix 如下：
//
// [
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
// ]
// 给定 target = 5，返回 true。
//
// 给定 target = 20，返回 false。

// 遍历
func searchMatrixFunc1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	for i := len(matrix[0]) - 1; i >= 0; i-- {
		flag := false
		for j := 0; j < len(matrix); j++ {
			if target == matrix[j][i] {
				return true
			}
			if target < matrix[j][i] {
				flag = true
				break
			}
		}

		if !flag {
			return false
		}
	}

	return false
}

// 指针法，从左下角开始，大于目标值向上，小于目标值向右
func searchMatrixFunc2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	col, row := 0, len(matrix)-1

	for col < len(matrix[0]) && row >= 0 {
		if target > matrix[row][col] {
			col++

		} else if target < matrix[row][col] {
			row--

		} else {
			return true
		}
	}

	return false
}
