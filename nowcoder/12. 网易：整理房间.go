package nowcoder

import "sort"

// 又到了周末，小易的房间乱得一团糟。
// 他希望将地上的杂物稍微整理下，使每团杂物看起来都紧凑一些，没有那么乱。
// 地上一共有n团杂物，每团杂物都包含4个物品。第i物品的坐标用(ai,bi)表示，小易每次都可以将它绕着(xi,yi)逆时针旋转90^ \circ90
// ∘
// ，这将消耗他的一次移动次数。如果一团杂物的4个点构成了一个面积不为0的正方形，我们说它是紧凑的。
// 因为小易很懒，所以他希望你帮助他计算一下每团杂物最少需要多少步移动能使它变得紧凑。
// 输入描述:
// 第一行一个数n(1 <= n <= 100)，表示杂物的团数。
// 接下来4n行，每4行表示一团杂物，每行4个数ai, bi，xi, yi, (-104 <= xi, yi, ai, bi <= 104)，表示第i个物品旋转的它本身的坐标和中心点坐标。
// 输出描述:
// n行，每行1个数，表示最少移动次数。
// 示例1
// 输入
// 复制
// 4
// 1 1 0 0
// -1 1 0 0
// -1 1 0 0
// 1 -1 0 0
// 1 1 0 0
// -2 1 0 0
// -1 1 0 0
// 1 -1 0 0
// 1 1 0 0
// -1 1 0 0
// -1 1 0 0
// -1 1 0 0
// 2 2 0 1
// -1 0 0 -2
// 3 0 0 -2
// -1 1 -2 0
// 输出
// 复制
// 1
// -1
// 3
// 3
// 说明
// 对于第一团杂物，我们可以旋转第二个或者第三个物品1次。

func handle12(points [][]int) int {
	res := -1
	temp := make([][]int, 0)
	var _handle func(cur int, count int)
	_handle = func(cur int, count int) {
		if len(temp) == 4 {
			if check(temp) {

				if res == -1 || count < res {
					res = count
				}
			}

			return
		}
		temp = append(temp, []int{points[cur][0], points[cur][1]})
		_handle(cur+1, count)
		temp = temp[:len(temp)-1]

		// 旋转三次
		a, b := -1*(points[cur][1]-points[cur][3])+points[cur][2], points[cur][0]-points[cur][2]+points[cur][3]
		temp = append(temp, []int{a, b})
		_handle(cur+1, count+1)
		temp = temp[:len(temp)-1]

		a, b = -1*(points[cur][0]-points[cur][2])+points[cur][2], -1*(points[cur][1]-points[cur][3])+points[cur][3]
		temp = append(temp, []int{a, b})
		_handle(cur+1, count+2)
		temp = temp[:len(temp)-1]

		a, b = (points[cur][1]-points[cur][3])+points[cur][2], -points[cur][0]+points[cur][2]+points[cur][3]
		temp = append(temp, []int{a, b})
		_handle(cur+1, count+3)
		temp = temp[:len(temp)-1]
	}

	_handle(0, 0)
	return res
}

func check(tp [][]int) bool {
	temp := make([][]int, len(tp))
	copy(temp, tp)
	sort.Slice(temp, func(left, right int) bool {
		if temp[left][0] == temp[right][0] {
			return temp[left][1] < temp[right][1]
		}

		return temp[left][0] < temp[right][0]
	})

	if temp[0][0] == temp[1][0] && temp[0][1] == temp[1][1] {
		return false
	}

	isRec := temp[0][0]+temp[3][0] == temp[2][0]+temp[1][0] && temp[0][1]+temp[3][1] == temp[2][1]+temp[1][1]

	if temp[0][0]-temp[3][0] == 0 {
		return temp[2][1]-temp[1][1] == 0 && isRec
	} else if temp[2][0]-temp[1][0] == 0 {
		return temp[0][1]-temp[3][1] == 0 && isRec
	} else {
		return (temp[0][1]-temp[3][1])/(temp[0][0]-temp[3][0])*(temp[2][1]-temp[1][1])/(temp[2][0]-temp[1][0]) == -1 && isRec
	}

	return false
}
