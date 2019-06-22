package leetcode

// 螺旋矩阵

// 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

// 示例 1:

// 输入:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// 输出: [1,2,3,6,9,8,7,4,5]
// 示例 2:

// 输入:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// 输出: [1,2,3,4,8,12,11,10,9,5,6,7]

// 思路1：计算矩阵的长和宽，每次走过一次就减掉一个长或宽，以走完四个方向作为一次循环，遍历整个矩阵。

func spiralOrder(matrix [][]int) []int {
    result := make([]int, 0)
    if len(matrix) == 0{
        return result
    }
    width := len(matrix[0])
    height := len(matrix)
    x, y := -1, 0
    
    add := func(part int) bool {
        result = append(result, part)
        
        if len(result) == len(matrix[0]) * len(matrix) {
            return true
        }
        return false
    }
    
    for {
        for i := 0; i < width; i++ {
            x++
            if add(matrix[y][x]) {
                return result
            }
        }
        height--

        for i := 0; i < height; i++ {
            y++
            
            if add(matrix[y][x]) {
                return result
            }
        }
        
        width--
        
        for i := 0; i < width; i++ {
            x--
            
            if add(matrix[y][x]) {
                return result
            }
        }
        
        height--

        for i := 0; i < height; i++ {
            y--
            
            if add(matrix[y][x]) {
                return result
            }
        }
        
        width--
    }
}