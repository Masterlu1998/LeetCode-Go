package nowcoder

import (
	"strconv"
	"strings"
)

// 题目描述
// 如果version1 > version2 返回1，如果 version1 < version2 返回-1，不然返回0.
//
// 输入的version字符串非空，只包含数字和字符.。.字符不代表通常意义上的小数点，只是用来区分数字序列。例如字符串2.5并不代表二点五，只是代表版本是第一级版本号是2，第二级版本号是5.
//
// 输入描述:
// 两个字符串，用空格分割。
// 每个字符串为一个version字符串，非空，只包含数字和字符.
// 输出描述:
// 只能输出1, -1，或0
// 示例1
// 输入
// 0.1 1.1
// 输出
// -1
// 备注:
// version1和version2的长度不超过1000，由小数点'.'分隔的每个数字不超过256。

func handle20(version1, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")

	for len(v1) != 0 && len(v2) != 0 {
		left, _ := strconv.Atoi(v1[0])
		right, _ := strconv.Atoi(v2[0])
		v1 = v1[1:]
		v2 = v2[1:]

		if left > right {
			return 1
		}

		if left < right {
			return -1
		}
	}

	if len(v1) > 0 {
		return 1
	}

	if len(v2) > 0 {
		return -1
	}

	return 0
}
