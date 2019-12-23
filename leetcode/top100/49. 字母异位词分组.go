package top100

import "sort"

// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
// 输出:
// [
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
// ]
// 说明：
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。

// 思路1：排序后的字符串作为map key值，key值相等的归为一类即可
func groupAnagrams(strs []string) [][]string {
	cache := make(map[string]*[]string)
	result := make([][]string, 0)

	for i := 0; i < len(strs); i++ {
		temp := []byte(strs[i])
		sort.Slice(temp, func(left, right int) bool {
			return temp[left] < temp[right]
		})

		if val, ok := cache[string(temp)]; ok {
			*val = append(*val, strs[i])
		} else {
			val := make([]string, 0)
			val = append(val, strs[i])
			cache[string(temp)] = &val
		}
	}

	for _, list := range cache {
		result = append(result, *list)
	}

	return result
}
