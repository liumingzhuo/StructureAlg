/*
712. 两个字符串的最小ASCII删除和
*/
package minimumDeleteSum712

var memo [][]int

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	memo = make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			memo[i][j] = -1
		}
	}
	return dp(s1, s2, 0, 0)
}
func dp(s1, s2 string, i, j int) int {
	var res int = 0
	//base
	if i == len(s1) { // 如果 s1走完 那么s2剩余部分都需要删除
		for ; j < len(s2); j++ {
			res += int(s2[j])
		}
		return res
	}
	//如果s2走完
	if j == len(s2) {
		for ; i < len(s1); i++ {
			res += int(s1[i])
		}
		return res
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if s1[i] == s2[j] {
		memo[i][j] = dp(s1, s2, i+1, j+1)
	} else {
		memo[i][j] = minValue(
			dp(s1, s2, i+1, j)+int(s1[i]),
			dp(s1, s2, i, j+1)+int(s2[j]),
		)
	}
	return memo[i][j]
}
func minValue(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
