package backtracking

// solveSudoku leetcode 37 解数独  leetcode 测试未通过，改用c#按此思路重写
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
// 数独部分空格内已填入了数字，空白格用 '.' 表示
func solveSudoku(board [][]byte) {
	sudoBacktrack(board, 0, 0)
}

func sudoBacktrack(board [][]byte, i, j int) bool {
	m, n := 9, 9
	// 当列遍历完成，行+1
	if j == n {
		return sudoBacktrack(board, i+1, 0)
	}
	// 如果行遍历完毕，找到一个可行解
	if i == m {
		return true
	}
	if board[i][j] != '.' {

		return sudoBacktrack(board, i, j+1)
	}
	for ch := '1'; ch <= '9'; ch++ {
		if !isValid(board, i, j, byte(ch)) {
			continue
		}
		// select
		board[i][j] = byte(ch)
		if sudoBacktrack(board, i, j+1) {
			return true
		}
		// unselect
		board[i][j] = '.'
	}
	// 遍历完 1~9 依然找不到可行解
	return false
}

func isValid(board [][]byte, r, c int, ch byte) bool {
	for i := 0; i < len(board); i++ {
		if board[r][i] == ch {
			return false
		}
		if board[i][c] == ch {
			return false
		}
		if board[(r/3)*3+i/3][(c/3)*3+i%3] == ch {
			return false
		}
	}
	return true
}
