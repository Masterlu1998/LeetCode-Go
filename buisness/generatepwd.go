package buisness

import "crypto/rand"

// 生成随机密码算法
// 思路：通过rand.Read()函数在byte切片中随机生成uint8的数字，然后用数字余除字母表的长度，生成下标
// ，从而在字母表中完全随机的选择字母。

var character = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

func GeneratePwd(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	out := make([]byte, length)
	for i := range out {
		index := uint8(b[i]) % uint8(len(character))
		out[i] = character[index]
	}
	return string(out)
}
