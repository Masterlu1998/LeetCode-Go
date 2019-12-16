package top100

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 示例 2：
//
// 输入: "cbbd"
// 输出: "bb"

// 思路1：围绕字符串的字符或者间隔展开，遍历所有的回文字符串，找到其中最长的那个回文字符串
func longestPalindromeFunc1(s string) string {
	if len(s) == 1 {
		return s
	}

	final := ""
	for i := 0; i <= 2*len(s)-1; i++ {
		cur := check(i, s)
		if len(final) < len(cur) {
			final = cur
		}
	}

	return final
}

func check(index int, s string) string {
	if index == 0 {
		return ""
	}

	if index%2 == 1 {
		result := ""
		// 间隔为中心
		for i := 0; ; i = i + 2 {
			left, right := (index-1-i)/2, (index+1+i)/2
			if left < 0 || right > len(s)-1 || s[left] != s[right] {
				break
			}
			result = string(s[left]) + result + string(s[right])
		}
		return result
	}

	// 字符为中心
	result := string(s[index/2])
	for i := 2; ; i = i + 2 {
		left, right := (index-i)/2, (index+i)/2
		if left < 0 || right > len(s)-1 || s[left] != s[right] {
			break
		}
		result = string(s[left]) + result + string(s[right])
	}
	return result
}
