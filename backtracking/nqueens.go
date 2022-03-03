package backtracking

var nqueensRes [][][]byte = make([][][]byte, 0)

// solveQueens 八皇后问题,输入棋盘边长n，返回所有合法的位置
func solveQueens(n int) [][][]byte {
	board := make([][]byte, n)
	for i := 0; i < len(board); i++ {
		board[i] = make([]byte, n)
	}
	soveQueensbacktrack(board, 0)
	return nqueensRes
}

// soveQueensbacktrack 做决策
func soveQueensbacktrack(board [][]byte, row int) {
	// 结束条件
	if row == len(board) {
		cpboard := make([][]byte, len(board))
		for i := 0; i < len(board); i++ {
			cpboard[i] = make([]byte, len(board))
			copy(cpboard[i], board[i])
		}
		nqueensRes = append(nqueensRes, cpboard)
		return
	}
	for i := 0; i < len(board); i++ {
		// 排除不合法的选择
		if !isInvalid(board, row, i) {
			continue
		}
		// 做选择
		board[row][i] = 1
		// 进入下一次决策
		soveQueensbacktrack(board, row+1)
		// 撤销选择
		board[row][i] = 0
	}
}

// isInvalid 验证冲突
func isInvalid(board [][]byte, row, col int) bool {
	// 判断同列中是否有冲突的
	for i := 0; i < len(board); i++ {
		if board[i][col] == 1 {
			return false
		}
	}
	// 判断左上方是否有冲突
	for j, k := row-1, col-1; j >= 0 && k >= 0; {
		if board[j][k] == 1 {
			return false
		}
		j--
		k--
	}
	// 判断右上方是否有冲突
	for rtr, rtc := row-1, col+1; rtr >= 0 && rtc < len(board); {
		if board[rtr][rtc] == 1 {
			return false
		}
		rtr--
		rtc++
	}
	return true
}
