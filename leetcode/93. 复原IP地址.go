package leetcode

import (
	"strconv"
	"strings"
)

// 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
// 示例:
//
// 输入: "25525511135"
// 输出: ["255.255.11.135", "255.255.111.35"]

// 思路1：回溯法，每一个点可以移动三次，如果结果不合法直接返回，合法继续下一个点
// 直到三个点被点出来为止
var restoreIpAddressesResult []string

func restoreIpAddresses(s string) []string {
	restoreIpAddressesResult = make([]string, 0)
	_restoreIpAddresses(s, 1, make([]string, 0))
	return restoreIpAddressesResult
}

func _restoreIpAddresses(s string, count int, tmp []string) {
	if count == 5 {
		if len(tmp) == 4 && s == "" {
			restoreIpAddressesResult = append(restoreIpAddressesResult, strings.Join(tmp, "."))
		}
		return
	}

	if s == "" {
		return
	}

	for i := 1; i <= 3; i++ {
		if len(s) < i {
			return
		}
		numstr := s[:i]
		num, _ := strconv.Atoi(numstr)
		if num > 255 || (len(numstr) > 1 && numstr[0] == '0') {
			return
		}
		tmp = append(tmp, numstr)
		_restoreIpAddresses(s[i:], count+1, tmp)
		tmp = tmp[:len(tmp)-1]

	}
}
