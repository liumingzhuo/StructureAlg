package sudoku37

const (
	finalRow int = 9
	finalCol int = 9
)

/*
input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]

*/
func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}
func backtrack(board [][]byte, row int, col int) bool {
	//base case
	if col == finalCol { //列结束列  下一行开始
		return backtrack(board, row+1, 0)
	}
	if row == finalCol { // 行到最后结束
		return true
	}
	for i := row; i < finalRow; i++ {
		for j := col; j < finalCol; j++ {
			if board[i][j] != '.' {
				return backtrack(board, i, j+1)
			}
			for ch := 0; ch < 9; ch++ {
				//排除不合理的选择
				if !isValid(board, i, j, byte(ch+'1')) {
					continue
				}
				board[i][j] = byte(ch + '1')
				if backtrack(board, i, j+1) {
					return true
				}
				board[i][j] = '.'
			}
			return false
		}
	}
	return false
}
func isValid(board [][]byte, r int, c int, ch byte) bool {
	for i := 0; i < 9; i++ {
		//判断行
		if board[i][c] == ch {
			return false
		}
		//判断列
		if board[r][i] == ch {
			return false
		}
		//判断3*3
		if board[(r/3)*3+i/3][(c/3)*3+i%3] == ch {
			return false
		}
	}
	return true

}
