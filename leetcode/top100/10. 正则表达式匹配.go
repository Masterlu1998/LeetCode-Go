package top100

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
// 说明:
//
// s 可能为空，且只包含从 a-z 的小写字母。
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
// 示例 1:
//
// 输入:
// s = "aa"
// p = "a"
// 输出: false
// 解释: "a" 无法匹配 "aa" 整个字符串。
// 示例 2:
//
// 输入:
// s = "aa"
// p = "a*"
// 输出: true
// 解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
// 示例 3:
//
// 输入:
// s = "ab"
// p = ".*"
// 输出: true
// 解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
// 示例 4:
//
// 输入:
// s = "aab"
// p = "c*a*b"
// 输出: true
// 解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
// 示例 5:
//
// 输入:
// s = "mississippi"
// p = "mis*is*p*."
// 输出: false

// 思路1：递归，先进行一次判断，当没有*号的时候s和p取下一个子串做比较，
// 有*号的时候分两种情况，第一种去除*号继续匹配，第二种用前一个字母的值继续匹配下去。
func isMatchFunc1(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	isFirstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isFirstMatch && isMatchFunc1(s[1:], p) || isMatchFunc1(s, p[2:])
	}

	return isFirstMatch && isMatchFunc1(s[1:], p[1:])
}

// 思路2：加入备忘录，加速递归速度
var memo [][]int

func isMatchFunc2(s string, p string) bool {
	memo = make([][]int, 0)
	for i := 0; i < len(p)+1; i++ {
		tmp := make([]int, len(s)+1)
		memo = append(memo, tmp)
	}

	return dp(0, 0, s, p)
}

func dp(i, j int, s, p string) bool {
	if memo[j][i] != 0 {
		return memo[j][i] == 1
	}
	var result bool
	if j == len(p) {
		result = i == len(s)

	} else {
		firstMatch := i != len(s) && (s[i] == p[j] || p[j] == '.')

		if len(p)-j >= 2 && p[j+1] == '*' {
			result = firstMatch && dp(i+1, j, s, p) || dp(i, j+2, s, p)
		} else {
			result = firstMatch && dp(i+1, j+1, s, p)
		}
	}

	if result {
		memo[j][i] = 1
	} else {
		memo[j][i] = -1
	}

	return result
}
