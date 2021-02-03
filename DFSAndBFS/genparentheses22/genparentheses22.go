package genparentheses22

/*
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Input: n = 1
Output: ["()"]
*/
// var res []string

func generateParenthesis(n int) []string {
	res := []string{}
	path := []byte{}
	if n <= 0 {
		return res
	}
	backtrack(&res, n, n, path)
	return res
}
func backtrack(res *[]string, left int, right int, path []byte) {
	if left > right {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		tmp := make([]byte, len(path))
		copy(tmp, path)
		*res = append(*res, string(tmp))
		return
	}
	//加入左括号
	path = append(path, '(')
	backtrack(res, left-1, right, path)
	path = path[:len(path)-1]

	//加入右括号
	path = append(path, ')')
	backtrack(res, left, right-1, path)
	path = path[:len(path)-1]
}
