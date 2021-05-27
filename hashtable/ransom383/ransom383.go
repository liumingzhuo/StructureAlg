/*
 *
 * [383] 赎金信
 */
package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := make(map[byte]int)
	for i, _ := range magazine {
		m[magazine[i]]++
	}
	for i, _ := range ransomNote {
		if v, ok := m[ransomNote[i]]; ok && v >= 1 {
			m[ransomNote[i]]--
		} else {
			return false
		}
	}
	return true
}
func main() {
	got := canConstruct("a", "b")
	fmt.Println(got)
}
