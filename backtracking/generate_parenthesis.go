package backtracking

var gpRes []string

// generateParenthesis Leetcode 22 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
func generateParenthesis(n int) []string {
	gpRes = make([]string, 0)
	if n == 0 {
		return gpRes
	}
	track := ""
	gpBacktrack(n, n, track)
	return gpRes
}

func gpBacktrack(left, right int, track string) {
	if right < left {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		gpRes = append(gpRes, track)
		return
	}
	//select left
	track = track + "("
	gpBacktrack(left-1, right, track)
	track = track[:len(track)-1]

	// select right
	track = track + ")"
	gpBacktrack(left, right-1, track)
	track = track[:len(track)-1]
}
