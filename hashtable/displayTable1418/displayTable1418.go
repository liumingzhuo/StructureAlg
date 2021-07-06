package main

import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	// 菜名
	dishM := make(map[string]struct{})
	// 桌号和菜名 以及数量
	tableDishC := make(map[int]map[string]int)
	for _, item := range orders {
		dishM[item[2]] = struct{}{}
		tableId, err := strconv.Atoi(item[1])
		if err != nil {
			continue
		}
		if _, ok := tableDishC[tableId]; !ok {
			tableDishC[tableId] = make(map[string]int)
		}
		tableDishC[tableId][item[2]]++
	}
	//对菜名进行排序
	m := len(dishM)
	dishs := make([]string, 0, m)
	for k := range dishM {
		dishs = append(dishs, k)
	}
	sort.Strings(dishs)

	// 对桌号进行排序
	n := len(tableDishC)
	ids := make([]int, 0, n)
	for id := range tableDishC {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	//填表
	ans := make([][]string, n+1)
	ans[0] = make([]string, 1, m+1)
	ans[0][0] = "Table"
	ans[0] = append(ans[0], dishs...)
	for i, id := range ids {
		dishc := tableDishC[id]
		ans[i+1] = make([]string, m+1)
		ans[i+1][0] = strconv.Itoa(id)
		for j, name := range dishs {
			count := dishc[name]
			ans[i+1][j+1] = strconv.Itoa(count)
		}
	}
	return ans
}

func main() {
	orders := [][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}}
	// want := [][]string{{"Table", "Beef Burrito", "Ceviche", "Fried Chicken", "Water"}, {"3", "0", "2", "1", "0"}, {"5", "0", "1", "0", "1"}, {"10", "1", "0", "0", "0"}}
	got := displayTable(orders)
	fmt.Println(got)
}
