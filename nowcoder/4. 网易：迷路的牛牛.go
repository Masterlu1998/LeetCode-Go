package nowcoder

// 牛牛去犇犇老师家补课，出门的时候面向北方，但是现在他迷路了。虽然他手里有一张地图，但是他需要知道自己面向哪个方向，请你帮帮他。
// 输入描述:
// 每个输入包含一个测试用例。
// 每个测试用例的第一行包含一个正整数，表示转方向的次数N(N<=1000)。
// 接下来的一行包含一个长度为N的字符串，由L和R组成，L表示向左转，R表示向右转。
// 输出描述:
// 输出牛牛最后面向的方向，N表示北，S表示南，E表示东，W表示西。
// 示例1
// 输入
// 复制
// 3
// LRR
// 输出
// 复制
// E

func handle_4(actions string) string {
	cur := buildList()
	for i := 0; i < len(actions); i++ {
		if actions[i] == 'L' {
			cur = cur.prev
		} else {
			cur = cur.next
		}
	}

	return cur.val
}

type listNode struct {
	prev, next *listNode
	val        string
}

func buildList() *listNode {
	n, s, e, w := &listNode{val: "N"}, &listNode{val: "S"}, &listNode{val: "E"}, &listNode{val: "W"}
	n.next, n.prev = e, w
	s.next, s.prev = w, e
	e.next, e.prev = s, n
	w.next, w.prev = n, s
	return n
}
