package nowcoder

import (
	"fmt"
	"strconv"
)

// 游戏网站提供若干升级补丁，每个补丁大小不一，玩家要升级到最新版，如何选择下载哪些补丁下载量最小。
//
//
//
// 输入
// 第一行输入              第一个数为用户版本  第二个数为最新版本，空格分开
//
// 接着输入N行补丁数据        第一个数补丁开始版本 第二个数为补丁结束版本 第三个数为补丁大小，空格分开
//
//
//
// 样例输入
// 1000 1050
//
// 1000 1020 50
//
// 1000 1030 70
//
// 1020 1030 15
//
// 1020 1040 30
//
// 1030 1050 40
//
// 1040 1050 20
//
//
//
// 输出
// 对于每个测试实例，输出一个升级路径以及最后实际升级的大小
//
// 样例输出
// 1000->1020->1040->1050(100)

func main25() {
	for {
		nowVersion, targetVersion := 0, 0
		p, err := fmt.Scan(&nowVersion, &targetVersion)
		if p == 0 || err != nil {
			return
		}

		relation := make(map[int][][]int)
		for {
			old, news, metrics := 0, 0, 0
			p, err := fmt.Scan(&old, &news, &metrics)
			if p == 0 || err != nil {
				break
			}

			relation[old] = append(relation[old], []int{news, metrics})
		}

		fmt.Println(handle25(relation, nowVersion, targetVersion))
	}

}

type element25 struct {
	metrics int
	val     int
	history []int
}

type queue25 []element25

func newQueue25() *queue25 {
	q := queue25(make([]element25, 0))
	return &q
}

func (q *queue25) enQueue(i element25) {
	*q = append(*q, i)
}

func (q *queue25) deQueue() {
	*q = (*q)[1:]
}

func (q *queue25) get() element25 {
	return (*q)[0]
}

func (q *queue25) empty() bool {
	return len(*q) == 0
}

func (q *queue25) length() int {
	return len(*q)
}

func handle25(relation map[int][][]int, nowVersion, targetVersion int) string {
	q := newQueue25()
	q.enQueue(element25{
		metrics: 0,
		val:     nowVersion,
		history: []int{nowVersion},
	})

	minMetrics := int(^uint(0) >> 1)
	var history []int
	for !q.empty() {
		length := q.length()
		for i := 0; i < length; i++ {
			now := q.get()
			q.deQueue()
			if now.val == targetVersion {
				if now.metrics < minMetrics {
					minMetrics = now.metrics
					history = now.history
				}
			}
			rel := relation[now.val]
			for _, next := range rel {
				q.enQueue(element25{
					metrics: next[1] + now.metrics,
					val:     next[0],
					history: append(now.history, next[0]),
				})
			}
		}
	}

	final := ""
	for _, version := range history {
		final += strconv.Itoa(version) + "->"
	}

	if len(final) == 0 {
		return ""
	}

	return final[:len(final)-2] + "(" + strconv.Itoa(minMetrics) + ")"
}
