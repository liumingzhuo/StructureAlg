package n_queens51

/*
51. N 皇后
*/

var res [][]string

func solveNQueens(n int) [][]string {
	res = [][]string{}
	board := make([][]string, 0)
	for i := 0; i < n; i++ {
		temp := make([]string, 0)
		for j := 0; j < n; j++ {
			temp = append(temp, ".")
		}
		board = append(board, temp)
	}
	backtrack(board, 0)
	return res
}
func backtrack(board [][]string, row int) bool {
	if row == len(board) {
		temp := make([]string, 0)
		// 将每行的选择结果改为字符串  [[".",".","Q","."],..] => ["..Q.", ...]
		for _, v := range board {
			str := ""
			for _, e := range v {
				str += e
			}
			temp = append(temp, str)
		}
		res = append(res, temp)
		return true
	}
	n := len(board[row])
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}
		//排除不合法的选择
		board[row][col] = "Q"
		if backtrack(board, row+1) {
			return true
		}
		board[row][col] = "."

	}
	return false
}
func isValid(board [][]string, row int, col int) bool {
	n := len(board)
	//col
	for r := 0; r < n; r++ {
		if board[r][col] == "Q" {
			return false
		}
	}
	//左上角
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == "Q" {
			return false
		}
	}
	//右上角
	for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
		if board[r][c] == "Q" {
			return false
		}
	}
	return true
}
