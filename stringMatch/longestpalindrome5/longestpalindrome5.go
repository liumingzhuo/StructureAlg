package longestpalindrome5

//5. 最长回文子串
func longestPalindrome(s string) string {
	//"cbbd"
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := palindrome(s, i, i)
		l2, r2 := palindrome(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]

}
func palindrome(s string, l, r int) (int, int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}
