package dpquestions

import "math"

// editMinDistance 两个单词返回word1 转成word2 的最小操作数
func editminDistance(word1, word2 string) int {
	// 用于优化重叠子问题
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}

	}
	return dpforEditMinDistance(word1, word2, len(word1)-1, len(word2)-1, memo)
}

func dpforEditMinDistance(s1, s2 string, i, j int, memo [][]int) int {
	//base case ,如果i或者是j哪一个遍历完，则将另一个剩余的字符添加到另一方 ,i,j 都是下标所有要+1
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	res := 0
	// 如果两个字符相等，则代表不需要进行任何操作，i,j同时移动
	if s1[i] == s2[j] {
		res = dpforEditMinDistance(s1, s2, i-1, j-1, memo)
	} else {

		// 插入
		d1 := dpforEditMinDistance(s1, s2, i, j-1, memo) + 1
		// 删除
		d2 := dpforEditMinDistance(s1, s2, i-1, j, memo) + 1
		// 替换
		d3 := dpforEditMinDistance(s1, s2, i-1, j-1, memo) + 1

		res = int(math.Min(math.Min(float64(d1), float64(d2)), float64(d3)))
	}
	memo[i][j] = res
	return res
}

// editMinDistancebyDPTable DP table 的方式求解
// TODO: 还没有完全测试过
func editMinDistancebyDPTable(s1, s2 string) int {
	dpt := make([][]int, len(s1)+1)
	for d := range dpt {
		dpt[d] = make([]int, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			// 如果两字符相同
			if s1[i-1] == s2[j-1] {
				dpt[i][j] = dpt[i-1][j-1]
			} else {
				// 插入
				d1 := dpt[i][j-1] + 1
				d2 := dpt[i-1][j] + 1
				d3 := dpt[i-1][j-1] + 1
				dpt[i][j] = int(math.Min(math.Min(float64(d1), float64(d2)), float64(d3)))
			}
		}
	}
	return dpt[len(s1)][len(s2)]
}
