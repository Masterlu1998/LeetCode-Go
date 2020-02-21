package codinginterviews

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。
//
// 示例:
//
// s = "abaccdeff"
// 返回 "b"
//
// s = ""
// 返回 " "

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}

	cache := make(map[byte]int)

	for i := range s {
		cache[s[i]]++
	}

	for i := range s {
		if count, ok := cache[s[i]]; ok && count == 1 {
			return s[i]
		}
	}

	return ' '
}
