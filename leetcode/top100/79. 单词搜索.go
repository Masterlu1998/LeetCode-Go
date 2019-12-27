package top100

// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
// 示例:
//
// board =
// [
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
// ]
//
// 给定 word = "ABCCED", 返回 true.
// 给定 word = "SEE", 返回 true.
// 给定 word = "ABCB", 返回 false.

// 回溯
var target string
var cache79 [][]bool

func existFunc1(board [][]byte, word string) bool {
	cache79 = make([][]bool, len(board))
	target = word
	result := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				result = result || _exist(board, j, i, 0)
			}
		}
	}
	return result
}

func _exist(board [][]byte, curx, cury, cur int) bool {
	if cur == len(target) {

		return true
	}

	if curx == len(board[0]) || curx == -1 || cury == -1 || cury == len(board) {
		return false
	}

	if cache79[cury] == nil {
		cache79[cury] = make([]bool, len(board[0]))
	}

	if board[cury][curx] != target[cur] || cache79[cury][curx] == true {
		return false
	}

	cache79[cury][curx] = true

	if _exist(board, curx+1, cury, cur+1) || _exist(board, curx, cury+1, cur+1) || _exist(board, curx-1, cury, cur+1) || _exist(board, curx, cury-1, cur+1) {
		return true
	}

	cache79[cury][curx] = false

	return false
}

// 更好的写法
func existFunc2(board [][]byte, word string) bool {
	cache := make([][]bool, len(board))

	var dfs func(curx, cury, index int) bool
	dfs = func(curx, cury, index int) bool {
		if index == len(word) {
			return true
		}

		if curx == -1 || cury == -1 || curx == len(board[0]) || cury == len(board) {
			return false
		}

		if cache[cury] == nil {
			cache[cury] = make([]bool, len(board[cury]))
		}

		if board[cury][curx] != word[index] || cache[cury][curx] == true {
			return false
		}

		cache[cury][curx] = true

		if dfs(curx+1, cury, index+1) || dfs(curx, cury+1, index+1) || dfs(curx-1, cury, index+1) || dfs(curx, cury-1, index+1) {
			return true
		}

		cache[cury][curx] = false

		return false
	}

	result := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			result = result || dfs(j, i, 0)
		}
	}

	return result
}
