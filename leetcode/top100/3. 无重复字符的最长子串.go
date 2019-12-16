package top100

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//

// 思路1：双指针，用一个map记录指针中间的所有字符和索引位置，如果发现重复，就把后指针放到重复的字母的后一位上
// 移动过程中随时计算长度，如果超过最大值就替换
func lengthOfLongestSubstringFunc1(s string) int {
	cache := make(map[byte]int)
	max, i := 0, 0
	for j := 0; j < len(s); j++ {
		if index, ok := cache[s[j]]; ok {
			// 不用删除以前的记录，只要判断i是否比index大就可以了
			if i <= index {
				i = index + 1
			}
		}
		cache[s[j]] = j
		length := j - i + 1
		if length > max {
			max = length
		}
	}
	return max
}

// 思路2：因为一个byte实际上大于0小于128,可以用数组来代替map
func lengthOfLongestSubstringFunc2(s string) int {
	result, slide := 0, make([]int, 128)

	for i, j := 0, 0; j < len(s); j++ {
		i = max(i, slide[s[j]])
		result = max(result, j-i+1)
		slide[s[j]] = j + 1
	}

	return result
}

func max(left, right int) int {
	if left > right {
		return left
	}

	return right
}
