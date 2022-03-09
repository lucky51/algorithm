package dpquestions

import "math"

// leetcode 1143
// longestCommonSubsequence 计算最长公共子序列（Longest Common Subsequence，简称 LCS)
// 给你输入两个字符串 s1 和 s2，请你找出他们俩的最长公共子序列，返回这个子序列的长度。
// 自顶向下的解法
func longestCommonSubsequence(s1, s2 string) int {
	memo := make([][]int, len(s1))
	for i := range memo {
		memo[i] = make([]int, len(s2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp(s1, s2, 0, 0, memo)
}

func dp(s1, s2 string, i, j int, memo [][]int) int {
	if i == len(s1) || j == len(s2) {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s1[i] == s2[j] {
		memo[i][j] = 1 + dp(s1, s2, i+1, j+1, memo)
	} else {
		memo[i][j] = int(math.Max(float64(dp(s1, s2, i, j+1, memo)), float64(dp(s1, s2, i+1, j, memo))))
	}
	return memo[i][j]
}

// longestCommonSubsequenceSlim 自底向上的解法
func longestCommonSubsequenceSlim(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for d := range dp {
		dp[d] = make([]int, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(dp[i-1][j])))
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

// minDistance 两个字符串的删除操作 ,可转化为求取最长公共子序列
func minDistance(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	dp[0] = make([]int, len(s2)+1)
	for i := 1; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return len(s1) - dp[len(s1)][len(s2)] + len(s2) - dp[len(s1)][len(s2)]
}

func minimumDeleteSum(s1, s2 string) int {
	memo := make([][]int, len(s1))
	for m := range memo {
		memo[m] = make([]int, len(s2))
		for mi := range memo[m] {
			memo[m][mi] = -1
		}
	}
	return dpscii(s1, s2, 0, 0, memo)
}

func dpscii(s1, s2 string, i, j int, memo [][]int) int {
	res := 0
	if len(s1) == i {
		for ; j < len(s2); j++ {
			res += int(s2[j])
		}
		return res
	}
	if len(s2) == j {
		for ; i < len(s1); i++ {
			res += int(s1[i])
		}
		return res
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s1[i] == s2[j] {
		memo[i][j] = dpscii(s1, s2, i+1, j+1, memo)
	} else {
		memo[i][j] = int(math.Min(float64(int(s2[j])+dpscii(s1, s2, i, j+1, memo)), float64(int(s1[i])+dpscii(s1, s2, i+1, j, memo))))
	}
	return memo[i][j]

}
