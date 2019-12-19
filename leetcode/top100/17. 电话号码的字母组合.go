package top100

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
// 示例:
//
// 输入："23"
// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 说明:
// 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// 倒序与正序算同一种

var cache = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

// 思路1：简单的递归，只要在上一次的结果上加上本次数字对应的字母就是答案
func letterCombinationsFunc1(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	if len(digits) == 1 {
		return cache[digits]
	}
	lastResult := letterCombinationsFunc1(digits[1:])
	var curResult []string
	for _, char := range cache[string(digits[0])] {
		for _, last := range lastResult {
			curResult = append(curResult, char+last)
		}
	}
	return curResult
}

// 思路2：标准回溯法起手式
var result []string

func letterCombinationsFunc2(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	result = make([]string, 0)
	_letterCombinations(digits, "")
	return result
}

func _letterCombinations(digits, temp string) {
	if len(digits) == 0 {
		result = append(result, temp)
		return
	}

	for _, char := range cache[string(digits[0])] {
		temp = temp + char
		_letterCombinations(digits[1:], temp)
		temp = temp[:len(temp)-1]
	}
}
