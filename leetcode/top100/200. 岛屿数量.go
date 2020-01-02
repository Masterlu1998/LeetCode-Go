package top100

// 给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
//
// 示例 1:
//
// 输入:
// 11110
// 11010
// 11000
// 00000
//
// 输出: 1
// 示例 2:
//
// 输入:
// 11000
// 11000
// 00100
// 00011
//
// 输出: 3

// 深度遍历
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs200(j, i, grid)
			}
		}
	}

	return count
}

func dfs200(x, y int, grid [][]byte) {
	if x >= len(grid[0]) || y >= len(grid) || x < 0 || y < 0 {
		return
	}

	if grid[y][x] == '0' {
		return
	}

	grid[y][x] = '0'
	dfs200(x+1, y, grid)
	dfs200(x, y+1, grid)
	dfs200(x-1, y, grid)
	dfs200(x, y-1, grid)
}
