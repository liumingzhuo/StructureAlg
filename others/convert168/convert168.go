package main

import "fmt"

func convertToTitle(columnNumber int) string {
	ans := make([]byte, 0)
	for columnNumber > 0 {
		//区间[0,25]
		columnNumber--
		ans = append(ans, byte(columnNumber%26)+'A')
		columnNumber /= 26
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}

	return string(ans)
}

func main() {
	r := convertToTitle(28)
	fmt.Println(r)
}
