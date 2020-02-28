package main

import "fmt"

func main() {
	total := 0
	n, err := fmt.Scan(&total)
	if n == 0 || err != nil {
		return
	}
	for i := 0; i < total; i++ {
		length := 0
		n, err := fmt.Scan(&length)
		if n == 0 || err != nil {
			return
		}

		var road string
		n, err = fmt.Scan(&road)
		if n == 0 || err != nil {
			return
		}
		fmt.Println(road)
		fmt.Println(handle(road))
	}
}

func handle(road string) int {
	if len(road) == 0 {
		return 0
	}

	count := 0
	cache := make([]bool, 0)
	for cur := 0; cur < len(road); cur++ {
		if road[cur] == '.' {
			if cur-1 >= 0 && cache[cur-1] {
				continue
			}

			if cur+2 < len(road) && road[cur+2] == '.' {
				cache[cur+1] = true
			} else {
				cache[cur] = true
			}
			count++
		}
	}

	return count
}
