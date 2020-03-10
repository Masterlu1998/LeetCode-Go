package nowcoder

// 题目描述
//
//
// 小东和其他小朋友正在玩一个关于选举的游戏。选举是通过投票的方式进行的，得票最多的人将获胜。
//
//
//
// 小东是编号为1的候选者，此外还有其他的候选者参加选举。根据初步的调查情况，所有准备投票的小朋友都有一定的投票倾向性，小东如果要获得胜利，必须争取部分准备为其他候选人投票的小朋友。由于小东的资源较为有限，她希望用最小的代价赢得胜利，请你帮忙计算她最少需要争取的选票数。
//
//
//
//
//
// 输入
// 输入有若干组，每组包含两行，第一行为一个正整数n（2<=n<=100），表示候选者的数量，第二行为每个候选人预期得到的选票数，以空格分开，每人的预期得票数在1到1000之间（包含1和1000）。
//
// 经过小东的争取后，可能出现候选人得票数为0或超过1000的情况。
//
//
//
// 样例输入
// 5
//
// 5 1 11 2 8
//
// 4
//
// 1 8 8 8
//
// 2
//
// 7 6
//
//
//
// 输出
// 对每组测试数据，单独输出一行，内容为小东最少需要争取的选票数。
//
// 样例输出
// 4
//
// 6
//
// 0

import (
	"fmt"
)

type candidate struct {
	index      int
	getTickets int
}

func main27() {
	for {
		total := 0
		p, err := fmt.Scan(&total)
		if p == 0 || err != nil {
			return
		}

		candidates := make([]*candidate, 0)
		for i := 0; i < total; i++ {
			tickets := 0
			p, err := fmt.Scan(&tickets)
			if p == 0 || err != nil {
				return
			}
			candidates = append(candidates, &candidate{i, tickets})
		}

		fmt.Println(handle27(candidates))
	}
}

func handle27(c []*candidate) int {
	metrics := 0
	origin := c[0]

	rest := c[1:]
	quickSort(rest, 0, len(rest)-1)

	for origin.getTickets <= rest[0].getTickets {
		metrics++
		origin.getTickets += 1
		rest[0].getTickets -= 1
		quickSort(rest, 0, len(rest)-1)
	}

	return metrics
}

func quickSort(nums []*candidate, left, right int) {
	if left < right {
		index := findIndex(nums, left, right)
		quickSort(nums, left, index-1)
		quickSort(nums, index+1, right)
	}
}

func findIndex(nums []*candidate, left, right int) int {
	choosed := (left + right) / 2
	nums[left], nums[choosed] = nums[choosed], nums[left]
	pivot := left
	index := pivot + 1

	for i := left; i <= right; i++ {
		if nums[i].getTickets > nums[pivot].getTickets {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}

	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}
