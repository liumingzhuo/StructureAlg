package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func countOfAtoms(formula string) string {
	i, n := 0, len(formula)

	parseNum := func() (num int) {
		if i == n || !unicode.IsDigit(rune(formula[i])) {
			return 1
		}
		for ; i < n && unicode.IsDigit(rune(formula[i])); i++ {
			num = 10*num + int(formula[i]-'0')
		}
		return
	}

	parseAtom := func() string {
		start := i
		i++
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++
		}
		return formula[start:i]
	}
	// 初始化map的列表
	st := []map[string]int{{}}
	for i < n {
		if ch := formula[i]; ch == '(' {
			i++
			st = append(st, map[string]int{})
		} else if ch == ')' {
			//获取括号右侧的数据
			i++
			num := parseNum()
			//取出前一位队列的值
			atomNum := st[len(st)-1]
			st = st[:len(st)-1]
			for k, v := range atomNum {
				st[len(st)-1][k] += v * num
			}
		} else {
			// 字符或者是数字
			atom := parseAtom()
			num := parseNum()
			st[len(st)-1][atom] += num
		}
	}
	//全部存入栈后，第一位则是计算好的map
	atomNum := st[0]
	type pair struct {
		atom string
		num  int
	}
	pairs := make([]pair, 0, len(atomNum))
	for k, v := range atomNum {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].atom < pairs[j].atom })
	ans := []byte{}
	for _, p := range pairs {
		ans = append(ans, p.atom...)
		if p.num > 1 {
			ans = append(ans, strconv.Itoa(p.num)...)
		}
	}
	return string(ans)
}
func main() {
	input := "H2O"
	// want := "H2O"
	got := countOfAtoms(input)
	fmt.Println(got)
}
