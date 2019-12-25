package top100

// 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
//
// 示例：
//
// 输入: S = "ADOBECODEBANC", T = "ABC"
// 输出: "BANC"
// 说明：
//
// 如果 S 中不存这样的子串，则返回空字符串 ""。
// 如果 S 中存在这样的子串，我们保证它是唯一的答案。

// 滑动窗口思想，右指针往前推进，直到满足条件，左指针往前推进，直到不符合条件，然后重复
func minWindow(s string, t string) string {
	cache := make(map[byte]int)
	mark := make(map[byte]int)
	match := 0

	for i := 0; i < len(t); i++ {
		cache[t[i]]++
	}

	left, right := 0, 0
	min := ""
	for right < len(s) {
		if count, ok := cache[s[right]]; ok {
			mark[s[right]]++
			if count == mark[s[right]] {
				match++
			}
		}
		right++
		for right <= len(s) && match == len(cache) {
			temp := s[left:right]

			if min == "" || len(min) > len(temp) {
				min = temp
			}

			if count, ok := cache[s[left]]; ok {
				mark[s[left]]--
				if count > mark[s[left]] {
					match--
				}
			}
			left++
		}
	}

	return min

}
