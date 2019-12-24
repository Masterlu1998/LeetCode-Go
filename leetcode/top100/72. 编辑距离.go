package top100

// 给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 示例 1:
//
// 输入: word1 = "horse", word2 = "ros"
// 输出: 3
// 解释:
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// 示例 2:
//
// 输入: word1 = "intention", word2 = "execution"
// 输出: 5
// 解释:
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')

// 思路1：动态规划，思路看代码吧.....
var w1, w2 string
var cache72 map[[2]int]int

func minDistanceFunc1(word1 string, word2 string) int {
	w1, w2 = word1, word2
	cache72 = make(map[[2]int]int)
	return _minDistance(len(word1)-1, len(word2)-1)
}

func _minDistance(i, j int) int {
	if val, ok := cache72[[2]int{i, j}]; ok {
		return val
	}
	answer := 0

	if i == -1 {
		answer = j + 1
	} else if j == -1 {
		answer = i + 1
	} else if w1[i] == w2[j] {
		answer = _minDistance(i-1, j-1)
	} else {
		answer = min(min(_minDistance(i-1, j-1)+1, _minDistance(i, j-1)+1), _minDistance(i-1, j)+1)
	}

	cache72[[2]int{i, j}] = answer
	return answer
}

func minDistance(word1 string, word2 string) int {
	cache := make([][]int, len(word1)+1)

	for i := 0; i < len(word1)+1; i++ {
		cache[i] = make([]int, len(word2)+1)
		cache[i][0] = i
	}

	for i := 0; i < len(word2)+1; i++ {
		cache[0][i] = i
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				cache[i][j] = cache[i-1][j-1]
			} else {
				cache[i][j] = min(min(cache[i-1][j-1]+1, cache[i-1][j]+1), cache[i][j-1]+1)
			}
		}
	}

	return cache[len(word1)][len(word2)]
}

func min(left, right int) int {
	if left < right {
		return left
	}

	return right
}
