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
	return dp(s1, s2, m, n)
}
func dp(s1, s2 string, i, j int) int {
	var res int = 0
	//base
	if i == len(s1) {
		for j := range s2 {
			res += s2[j]
		}
	}

}
