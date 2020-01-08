package top100

// 给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。
//
// 然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
//
// 你需要计算完成所有任务所需要的最短时间。
//
// 示例 1：
//
// 输入: tasks = ["A","A","A","B","B","B"], n = 2
// 输出: 8
// 执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
// 注：
//
// 任务的总个数为 [1, 10000]。
// n 的取值范围为 [0, 100]。

type element struct {
	types byte
	show  int
}

// 以n+1一轮安排任务，保证每个任务都有充足的冷却时间，每次安排出现次数前n+1的任务
func leastInterval(tasks []byte, n int) int {
	cache := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		cache[tasks[i]]++
	}

	heap := make([]*element, 0)
	for types, count := range cache {
		e := &element{
			types: types,
			show:  count,
		}
		heap = append(heap, e)
	}

	buildMaxHeap621(heap)
	result := make([]byte, 0)

	for len(heap) != 0 {
		temp := make([]*element, 0)
		for i := 0; i <= n; i++ {
			if len(heap) != 0 {
				e := heap[0]
				result = append(result, e.types)
				e.show--
				if e.show != 0 {
					temp = append(temp, e)
				}
				heap = heap[1:]
				buildMaxHeap621(heap)
			} else {
				if len(temp) == 0 {
					break
				}
				result = append(result, 0)
			}
		}
		heap = append(heap, temp...)
		buildMaxHeap621(heap)
	}

	return len(result)
}

func buildMaxHeap621(nums []*element) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify621(nums, i)
	}
}

func heapify621(nums []*element, node int) {
	left, right := node*2+1, node*2+2
	max := node
	if left < len(nums) && nums[max].show < nums[left].show {
		max = left
	}

	if right < len(nums) && nums[max].show < nums[right].show {
		max = right
	}

	if max != node {
		nums[max], nums[node] = nums[node], nums[max]
		heapify621(nums, max)
	}
}
