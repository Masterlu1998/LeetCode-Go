package top100

// 删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
//
// 说明: 输入可能包含了除 ( 和 ) 以外的字符。
//
// 示例 1:
//
// 输入: "()())()"
// 输出: ["()()()", "(())()"]
// 示例 2:
//
// 输入: "(a)())()"
// 输出: ["(a)()()", "(a())()"]
// 示例 3:
//
// 输入: ")("
// 输出: [""]

// 暴力回溯
func removeInvalidParentheses(s string) []string {

	result := make([]string, 0)
	cache := make(map[string]bool)
	minChange := int(^uint(0) >> 1)
	var _removeInvalidParentheses func(s, temp string, left, right, change int)
	_removeInvalidParentheses = func(s, temp string, left, right, change int) {
		if len(s) == 0 {
			if left == right {
				if change == minChange {
					cache[temp] = true
				} else if change < minChange {
					cache = map[string]bool{
						temp: true,
					}
					minChange = change
				}
			}
			return
		}

		if right > left {
			return
		}

		if s[0] == '(' {
			_removeInvalidParentheses(s[1:], temp+"(", left+1, right, change)
			_removeInvalidParentheses(s[1:], temp, left, right, change+1)
		} else if s[0] == ')' {
			_removeInvalidParentheses(s[1:], temp+")", left, right+1, change)
			_removeInvalidParentheses(s[1:], temp, left, right, change+1)
		} else {
			_removeInvalidParentheses(s[1:], temp+string(s[0]), left, right, change)
		}
	}

	_removeInvalidParentheses(s, "", 0, 0, 0)

	for key, _ := range cache {
		result = append(result, key)
	}
	return result
}
