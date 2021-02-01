package n_queens51

/*
51. N 皇后
*/

var res [][]string

func solveNQueens(n int) [][]string {
	res = [][]string{}
	//初始化棋盘 n*n
	board := [][]string{}
	for i := 0; i < n; i++ { //行
		tmp := []string{}
		for j := 0; j < n; j++ { //列
			tmp = append(tmp, ".")
		}
		board = append(board, tmp)
	}
	backtrack(board, 0)
	return res
}

func backtrack(board [][]string, row int) {
	if row == len(board) {
		tmp := []string{}
		for _, v := range board { //[".",".","Q"] -> ["..Q"]
			str := ""
			for _, e := range v {
				str += e
			}
			tmp = append(tmp, str)
		}
		res = append(res, tmp)
		return
	}
	for col := 0; col < len(board[row]); col++ {
		//判断是否符合棋盘
		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = "Q"
		backtrack(board, row+1)
		board[row][col] = "."
	}
}
func isValid(board [][]string, row int, col int) bool {
	n := len(board)
	//判断列
	for r := 0; r < n; r++ {
		if board[r][col] == "Q" {
			return false
		}
	}
	//判断左上角
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == "Q" {
			return false
		}
	}
	//判断右上角
	for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
		if board[r][c] == "Q" {
			return false
		}
	}
	return true
}
