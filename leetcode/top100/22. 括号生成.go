package top100

// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
//
// 例如，给出 n = 3，生成结果为：
//
// [
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
// ]

var (
	result22 []string
	length   int
)

// 思路：回溯法加枝减，用左括号的个数和右括号的个数来枝剪不合法的结果
func generateParenthesis(n int) []string {
	result22 = make([]string, 0)
	length = 2 * n
	_generateParenthesis(n, n, "")
	return result22
}

func _generateParenthesis(left, right int, temp string) {
	if len(temp) == length {
		if left == 0 && right == 0 {
			result22 = append(result22, temp)
		}
		return
	}

	if left > 0 {
		_generateParenthesis(left-1, right, temp+"(")
	}
	if left < right && right > 0 {
		_generateParenthesis(left, right-1, temp+")")
	}
}
