package top100

import "sort"

// 假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。
//
// 注意：
// 总人数少于1100人。
//
// 示例
//
// 输入:
// [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
// 输出:
// [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

// 先把数组按身高从高到低排序，然后根据索引填入数组，需要注意，这里身高相同时，需要按索引从小到大排
func reconstructQueue(people [][]int) [][]int {
	result := make([][]int, 0)
	if len(people) == 0 {
		return result
	}
	sort.Slice(people, func(left, right int) bool {
		if people[left][0] == people[right][0] {
			return people[left][1] < people[right][1]
		}
		return people[left][0] > people[right][0]
	})

	for _, p := range people {
		index := p[1]

		if index >= len(result) {
			result = append(result, p)
		} else {
			result = append(result, result[len(result)-1])
			copy(result[index+1:], result[index:])
			result[index] = p
		}
	}

	return result
}
