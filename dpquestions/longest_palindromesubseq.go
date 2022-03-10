package dpquestions

import (
	"math"
)

// longestPalindRomeSubseq 最长回文子序列
func longestPalindRomeSubseq(s string) int {
	seqdp := make([][]int, len(s))
	for d := range seqdp {
		seqdp[d] = make([]int, len(s))
		for si := range seqdp[d] {
			// base case ,i和j相等代表i,j在同一位置上，所以指的是同一字符，或者说是仅一个字符的情况下都为1
			/*
				[1 0 0 0 0]
				[0 1 0 0 0]
				[0 0 1 0 0]
				[0 0 0 1 0]
			*/
			if si == d {
				seqdp[d][si] = 1
			}
		}
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				seqdp[i][j] = seqdp[i+1][j-1] + 2
			} else {
				if seqdp[i+1][j] > seqdp[i][j-1] {
					seqdp[i][j] = seqdp[i+1][j]
				} else {
					seqdp[i][j] = seqdp[i][j-1]
				}
			}
		}

	}
	return seqdp[0][len(s)-1]
}

// longestPalindRomeSubseqSlime 最长回文子序列降纬求解 ,不好理解！！！
func longestPalindRomeSubseqSlime(s string) int {

	//直接将二维替换为一维，就是数组中的i位置
	//dp := make([][]int, len(s))
	dp := make([]int, len(s))
	for d := range dp {
		dp[d] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < len(s); j++ {
			temp := dp[j]
			// dp[j]的位置
			if s[i] == s[j] {
				//dp[i][j] = dp[i-1][j-1] + 2
				//dp[j] = dp[i-1][j-1] + 2
				dp[j] = pre + 2
			} else {
				//dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
				dp[j] = int(math.Max(float64(dp[j]), float64(dp[j-1])))
			}
			pre = temp
		}

	}
	return dp[len(s)-1]
}

// minInsertions 让字符串成为回文串的最小插入次数
func minInsertions(s string) int {
	dp := make([][]int, len(s))
	for d := range dp {
		dp[d] = make([]int, len(s))
		// base case
		for di := range dp[d] {
			if d == di {
				dp[d][di] = 0
			}
		}
	}

	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			// 如果 si==si ，则判断子问题是否为回文串
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				// 如果不相同，取最小的
				dp[i][j] = int(math.Min(float64(dp[i][j-1]), float64(dp[i+1][j])))
			}
		}
	}
	return dp[0][len(s)-1]
}

// minInsertionsSlime 让字符串成为回文串的最小插入次数，降维解法
func minInsertionsSlime(s string) int {
	dp := make([]int, len(s))
	for i := len(s) - 2; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < len(s); j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = pre
			} else {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-1]))) + 1
			}
			pre = temp
		}
	}
	return dp[len(s)-1]
}
