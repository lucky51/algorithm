package dpquestions

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
