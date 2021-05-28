/**
 * [202] 快乐数
使用 map 去重
*/
package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]struct{})
	for {
		s := getSum(n)
		if s == 1 {
			return true
		}
		//如果m中出现存的sum 则重复出现，无限循环了
		if _, ok := m[s]; ok {
			return false
		}
		m[s] = struct{}{}
		n = s
	}
}
func getSum(num int) (sum int) {
	for num > 0 {
		//num%10  num/10 求取每一位的数字
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return
}
func main() {
	got := isHappy(2)
	fmt.Println(got)
}
