package backtracking

var gpRes []string

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
