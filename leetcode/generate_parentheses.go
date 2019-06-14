package leetcode

// 思路1：有条件的回溯法，添加左括号与右括号的时候加上条件判断。添加左括号的时候判断左括号的数量有没有超过括号总数
// 添加右括号的时候，判断左括号的数量有没有超过对应的右括号。当结果长度超过2*n的时候完成整个回溯。

var result []string
func generateParenthesis(n int) []string {
    result = make([]string, 0)
    _generateParenthesis(n, 0, 0, "")
    return result
}

func _generateParenthesis(n, leftNum, rightNum int, temp string) {
    if len(temp) == 2 * n {
        result = append(result, temp)
        return
    }
    
    if leftNum < n {
        temp += "("
        _generateParenthesis(n, leftNum+1, rightNum, temp)
        temp = temp[:len(temp)-1]
        if leftNum - rightNum > 0 {
            temp += ")"
            _generateParenthesis(n, leftNum, rightNum+1, temp)
        }
    } else {
        if leftNum - rightNum > 0 {
            temp += ")"
            _generateParenthesis(n, leftNum, rightNum+1, temp)
        }
    }

}