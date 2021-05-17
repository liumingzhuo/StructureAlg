/*
125. 验证回文串
*/
package palindrome125

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for !isalnum(s[l]) && l < r {
			l++
		}
		for !isalnum(s[r]) && l < r {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
		}
		l++
		r--
	}
	return true

}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
