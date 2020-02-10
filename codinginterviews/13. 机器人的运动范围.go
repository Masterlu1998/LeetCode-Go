package codinginterviews

import "fmt"

// 地上有一个m行n列的方格。机器人从坐标（0, 0）的格子开始移动，它可以每次向上下左右四个方向
// 移动一格，并且行坐标和列坐标之和不能大于k，求机器人可以到达的格子数量

func movingCount(thereshould, rows, cols int) int {
	// x -> cols y -> rows
	visited := make([][]bool, rows)

	count := 0

	var checkValid func(x, y int) bool
	checkValid = func(x, y int) bool {
		if x < 0 || y < 0 || x >= cols || y >= rows {
			return false
		}

		sum := getNumSum(x) + getNumSum(y)
		if sum > thereshould {
			return false
		}

		return true
	}

	var _movingCount func(x, y int, visited [][]bool)
	_movingCount = func(x, y int, visited [][]bool) {
		if !checkValid(x, y) {
			return
		}

		if len(visited[y]) == 0 {
			visited[y] = make([]bool, cols)
		}

		if !visited[y][x] {
			fmt.Println(x, y)
			visited[y][x] = true
			count++
			_movingCount(x-1, y, visited)
			_movingCount(x+1, y, visited)
			_movingCount(x, y-1, visited)
			_movingCount(x, y+1, visited)
		}
	}

	_movingCount(0, 0, visited)

	return count
}

func getNumSum(x int) int {
	result := 0
	for num := x; ; num /= 10 {
		temp := num % 10
		result += temp
		if num == 0 {
			break
		}
	}

	return result
}

func TestMovingCount() {
	fmt.Println(movingCount(2, 40, 40))
}
