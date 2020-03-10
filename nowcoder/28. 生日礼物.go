package nowcoder

import (
	"fmt"
	"strconv"
)

type letter struct {
	index int
	w, h  int
}

func main28() {
	for {
		total, w, h := 0, 0, 0
		p, err := fmt.Scan(&total, &w, &h)
		if p == 0 || err != nil {
			return
		}

		letters := make([]*letter, 0)
		for i := 0; i < total; i++ {
			width, height := 0, 0
			p, err := fmt.Scan(&width, &height)
			if p == 0 || err != nil {
				return
			}

			letters = append(letters, &letter{i + 1, width, height})
		}

		num, index := handle28(total, w, h, letters)
		if len(index) == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(num)
			letters := ""
			for _, i := range index {
				letters += strconv.Itoa(i.index) + " "
			}
			if len(letters) != 0 {
				fmt.Println(letters[:len(letters)-1])
			}
		}
	}
}

type element28 struct {
	rest    []*letter
	history []*letter
	w, h    int
}

type queue28 []*element28

func newQueue28() *queue28 {
	q := queue28(make([]*element28, 0))
	return &q
}

func (q *queue28) enQueue(i *element28) {
	*q = append(*q, i)
}

func (q *queue28) deQueue() {
	*q = (*q)[1:]
}

func (q *queue28) get() *element28 {
	return (*q)[0]
}

func (q *queue28) empty() bool {
	return len(*q) == 0
}

func (q *queue28) length() int {
	return len(*q)
}

func handle28(total int, w int, h int, letters []*letter) (int, []*letter) {
	q := newQueue28()
	quickSort28(letters, 0, len(letters)-1)
	q.enQueue(&element28{
		rest:    letters,
		history: make([]*letter, 0),
		w:       w,
		h:       h,
	})
	num := 0
	var lastElement *element28

	for !q.empty() {
		count := q.length()
		for i := 0; i < count; i++ {
			cur := q.get()
			lastElement = cur
			q.deQueue()
			for i, l := range cur.rest {
				if i >= 1 && cur.rest[i].h == cur.rest[i-1].h && cur.rest[i].w == cur.rest[i-1].w {
					continue
				}
				if l.w > cur.w && l.h > cur.h {
					restCp := make([]*letter, len(cur.rest))
					historyCp := make([]*letter, len(cur.history))
					copy(restCp, cur.rest)
					copy(historyCp, cur.history)
					q.enQueue(&element28{
						rest:    append(restCp[:i], restCp[i+1:]...),
						history: append(historyCp, l),
						w:       l.w,
						h:       l.h,
					})
				} else if l.w < cur.w && l.h < cur.h {
					break
				}
			}

		}
		num++
	}

	if lastElement == nil {
		return 0, nil
	}

	return len(lastElement.history), lastElement.history
}

func quickSort28(nums []*letter, left, right int) {
	if left < right {
		index := findIndex28(nums, left, right)
		quickSort28(nums, left, index-1)
		quickSort28(nums, index+1, right)
	}
}

func findIndex28(nums []*letter, left, right int) int {
	choosed := (left + right) / 2
	nums[left], nums[choosed] = nums[choosed], nums[left]
	pivot := left
	index := pivot + 1

	for i := left; i <= right; i++ {
		if nums[i].w < nums[pivot].w {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}

	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}
