package leetcode

import "strconv"

// 给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

// 示例:
//
// s = "3[a]2[bc]", 返回 "aaabcbc".
// s = "3[a2[c]]", 返回 "accaccacc".
// s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

// 思路： 遍历字符并压栈，遇到"]"，弹出栈中与第一个"["中间的字符作为内容，然后弹出栈中的数字作为重复的次数，然后将
// 结果压栈。重复上述步骤，然后遍历栈连接字符串。需要注意的是字符串的顺序问题，栈中的字符串保证从左到右顺序正确，需
// 要将栈取出的字符串放在最前面而不能简单的+=。

func decodeString(s string) string {
	stack := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, string(s[i]))
		} else {
			content := ""
			for len(stack) > 0 {
				last := len(stack) - 1
				if stack[last] == "[" {
					stack = stack[:last]
					break
				} else {
					content = stack[last] + content
					stack = stack[:last]
				}
			}

			numStr := ""
			for len(stack) > 0 {
				last := len(stack) - 1
				if stack[last] >= "0" && stack[last] <= "9" {
					numStr = stack[last] + numStr
					stack = stack[:last]
				} else {
					break
				}
			}

			result := ""
			num, _ := strconv.Atoi(numStr)

			for i := 0; i < num; i++ {
				result = content + result
			}

			stack = append(stack, result)
		}
	}

	final := ""
	for len(stack) != 0 {
		last := len(stack) - 1
		final = stack[last] + final
		stack = stack[:last]
	}

	return final
}
