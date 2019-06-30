package leetcode

// 79. 单词搜索

// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 示例:

// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]

// 给定 word = "ABCCED", 返回 true.
// 给定 word = "SEE", 返回 true.
// 给定 word = "ABCB", 返回 false.

// 思路1：先遍历整个矩阵，找到与头字母匹配的位置，然后在该位置上下左右进行试探匹配接下来的字母。需要查重判断，
// 这里使用了一个map来做重复判断。

var (
    width, height int
    cacheExi map[[2]int]bool
) 

func existFunc2(board [][]byte, word string) bool {
    title := word[0]
    cacheExi = make(map[[2]int]bool)
    width = len(board[0])
    height = len(board)
    result := false
    
    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            if board[i][j] == title {
                result = result || _existFunc2(board, word, 0, j, i)
                delete(cacheExi, [2]int{j, i})
            }
        }
    }
    return result
}

func _existFunc2(board [][]byte, word string, index, w, h int) bool {
    cacheExi[[2]int{w, h}] = true
    if board[h][w] != word[index] {
        return false
    }
    
    if index == len(word) - 1 {
        if board[h][w] != word[index] {
            return false
        }
        return true
    }
    
    result := false
    
    if w - 1 >= 0 && !cacheExi[[2]int{w-1, h}] {
        result = result || _existFunc2(board, word, index+1, w-1, h)
        delete(cacheExi, [2]int{w-1, h})
    }
    
    if w + 1 < width && !cacheExi[[2]int{w+1, h}] {
        result = result || _existFunc2(board, word, index+1, w+1, h)
        delete(cacheExi, [2]int{w+1, h})
    }
    
    if h - 1 >= 0 && !cacheExi[[2]int{w, h-1}] {
        result = result || _existFunc2(board, word, index+1, w, h-1)
        delete(cacheExi, [2]int{w, h-1})
    }
    
    if h + 1 < height && !cacheExi[[2]int{w, h+1}] {
        result = result || _existFunc2(board, word, index+1, w, h+1)
        delete(cacheExi, [2]int{w, h+1})
    }
    
    return result
}

// 思路2：原地修改数组来标记已经遍历过的位置，节省创建map的开销

func existFunc1(board [][]byte, word string) bool {
    title := word[0]
    width = len(board[0])
    height = len(board)
    result := false
    
    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            if board[i][j] == title {
                result = result || _existFunc1(board, word, 0, j, i)
                board[i][j] = word[0]
            }
        }
    }
    return result
}

func _existFunc1(board [][]byte, word string, index, w, h int) bool {
    if board[h][w] != word[index] {
        return false
    }
    
    if index == len(word) - 1 {
        return true
    }
    
    result := false
    tmp := board[h][w]
    board[h][w] = '#'
    if w - 1 >= 0 {
        result = result || _existFunc1(board, word, index+1, w-1, h)
    }
    
    if w + 1 < width {
        result = result || _existFunc1(board, word, index+1, w+1, h)
    }
    
    if h - 1 >= 0 {
        result = result || _existFunc1(board, word, index+1, w, h-1)
    }
    
    if h + 1 < height {
        result = result || _existFunc1(board, word, index+1, w, h+1)
    }
    
    board[h][w] = tmp
    return result
}