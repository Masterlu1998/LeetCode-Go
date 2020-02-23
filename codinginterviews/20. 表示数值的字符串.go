package codinginterviews

import "fmt"

// 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"及"-1E-16"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。

func isNumber(s string) bool {

	// s = strings.Trim(s, " ")
	cur := 0
	result, cur := isInt(cur, s)
	fmt.Println(result, cur)

	var temp bool
	if cur < len(s) && s[cur] == '.' {
		cur++
		temp, cur = isUnsignedInt(cur, s)
		result = result || temp
	}

	if cur < len(s) && (s[cur] == 'e' || s[cur] == 'E') {
		cur++
		temp, cur = isInt(cur, s)

		result = result && temp
	}

	return result
}

func isInt(index int, s string) (bool, int) {
	if index >= len(s) {
		return false, index
	}

	if s[index] == '+' || s[index] == '-' {
		index++
	}

	return isUnsignedInt(index, s)
}

func isUnsignedInt(index int, s string) (bool, int) {
	old := index
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		index++
	}
	return old < index, index
}
