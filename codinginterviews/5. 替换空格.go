package codinginterviews

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
// 示例 1：
//
// 输入：s = "We are happy."
// 输出："We%20are%20happy."

func replaceSpace(s string) string {
	count := 0
	for _, char := range s {
		if char == ' ' {
			count++
		}
	}

	result := make([]byte, count*2+len(s))

	i := len(result) - 1
	j := len(s) - 1

	for i >= 0 && j >= 0 {
		if s[j] == ' ' {
			result[i] = '0'
			result[i-1] = '2'
			result[i-2] = '%'
			i = i - 3
			j--
			continue
		}

		result[i] = s[j]
		i--
		j--
	}

	return string(result)
}
