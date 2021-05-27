/*
 *
 * [242] 有效的字母异位词
 */

package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//定义存储字符的数组
	record := [26]int{}
	for _, b := range s {
		record[b-'a']++
	}
	for _, b := range t {
		record[b-'a']--
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	got := isAnagram("rat", "car")
	fmt.Println(got)
}
